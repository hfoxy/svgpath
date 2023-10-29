package svgpath

import (
	"fmt"
	"github.com/hfoxy/svgpath/internal"
	"math"
)

type SVGPath struct {
	Length         float64
	partialLengths []float64
	parts          []internal.Properties
	initialPoint   internal.Point
}

func NewFromPath(path string) (SVGPath, error) {
	segments, err := Parse(path)
	if err != nil {
		return SVGPath{}, err
	}

	return NewFromSegments(segments)
}

func NewFromSegments(segments []Segment) (SVGPath, error) {
	r := SVGPath{}

	r.parts = make([]internal.Properties, 0, len(segments))

	cur := internal.Point{X: 0, Y: 0}
	prevPoint := internal.Point{X: 0, Y: 0}

	curve := internal.EmptyBezier
	ringStart := internal.Point{X: 0, Y: 0}

	for i, segment := range segments {
		if segment.Command == 'M' {
			if len(segment.Args) != 2 {
				return r, fmt.Errorf("malformed path data: '%c' must have 2 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			cur = internal.Point{X: segment.Args[0], Y: segment.Args[1]}
			ringStart = cur

			r.parts = append(r.parts, nil)
			if i == 0 {
				r.initialPoint = cur
			}
		} else if segment.Command == 'm' {
			if len(segment.Args) != 2 {
				return r, fmt.Errorf("malformed path data: '%c' must have 2 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			cur = internal.Point{X: segment.Args[0] + cur.X, Y: segment.Args[1] + cur.Y}
			ringStart = cur
			r.parts = append(r.parts, nil)
		} else if segment.Command == 'L' {
			if len(segment.Args) != 2 {
				return r, fmt.Errorf("malformed path data: '%c' must have 2 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			r.Length += math.Sqrt(math.Pow(cur.X-segment.Args[0], 2) + math.Pow(cur.Y-segment.Args[1], 2))
			r.parts = append(r.parts, internal.NewLinear(cur.X, segment.Args[0], cur.Y, segment.Args[1]))
			cur = internal.Point{X: segment.Args[0], Y: segment.Args[1]}
		} else if segment.Command == 'l' {
			if len(segment.Args) != 2 {
				return r, fmt.Errorf("malformed path data: '%c' must have 2 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			r.Length += math.Sqrt(math.Pow(segment.Args[0], 2) + math.Pow(segment.Args[1], 2))
			r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X+segment.Args[0], cur.Y, cur.Y+segment.Args[1]))
			cur = internal.Point{X: cur.X + segment.Args[0], Y: cur.Y + segment.Args[1]}
		} else if segment.Command == 'H' {
			if len(segment.Args) != 1 {
				return r, fmt.Errorf("malformed path data: '%c' must have 1 element and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			r.Length += math.Abs(cur.X - segment.Args[0])
			r.parts = append(r.parts, internal.NewLinear(cur.X, segment.Args[0], cur.Y, cur.Y))
			cur = internal.Point{X: segment.Args[0], Y: cur.Y}
		} else if segment.Command == 'h' {
			if len(segment.Args) != 1 {
				return r, fmt.Errorf("malformed path data: '%c' must have 1 element and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			r.Length += math.Abs(segment.Args[0])
			r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X+segment.Args[0], cur.Y, cur.Y))
			cur = internal.Point{X: cur.X + segment.Args[0], Y: cur.Y}
		} else if segment.Command == 'V' {
			if len(segment.Args) != 1 {
				return r, fmt.Errorf("malformed path data: '%c' must have 1 element and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			r.Length += math.Abs(cur.Y - segment.Args[0])
			r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X, cur.Y, segment.Args[0]))
			cur = internal.Point{X: cur.X, Y: segment.Args[0]}
		} else if segment.Command == 'v' {
			if len(segment.Args) != 1 {
				return r, fmt.Errorf("malformed path data: '%c' must have 1 element and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			r.Length += math.Abs(segment.Args[0])
			r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X, cur.Y, cur.Y+segment.Args[0]))
			cur = internal.Point{X: cur.X, Y: cur.Y + segment.Args[0]}
		} else if segment.Command == 'z' || segment.Command == 'Z' {
			if len(segment.Args) != 1 {
				return r, fmt.Errorf("malformed path data: '%c' must have 0 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			r.Length += math.Sqrt(math.Pow(ringStart.X-cur.X, 2) + math.Pow(ringStart.Y-cur.Y, 2))
			r.parts = append(r.parts, internal.NewLinear(cur.X, ringStart.X, cur.Y, ringStart.Y))
			cur = ringStart
		} else if segment.Command == 'C' {
			if len(segment.Args) != 6 {
				return r, fmt.Errorf("malformed path data: '%c' must have 6 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			curve = internal.NewBezier(
				cur,
				internal.Point{X: segment.Args[0], Y: segment.Args[1]},
				internal.Point{X: segment.Args[2], Y: segment.Args[3]},
				internal.Point{X: segment.Args[4], Y: segment.Args[5]},
			)

			r.Length += curve.GetTotalLength()
			r.parts = append(r.parts, curve)
			cur = internal.Point{X: segment.Args[4], Y: segment.Args[5]}
		} else if segment.Command == 'c' {
			if len(segment.Args) != 6 {
				recheckArgs, err2 := parseValues(segment.Raw[1:])
				if err2 != nil {
					return r, fmt.Errorf("malformed path data: '%c' => '%s': %w", segment.Command, segment.Raw, err2)
				}

				if len(recheckArgs) != 6 {
					return r, fmt.Errorf("[recheck] malformed path data: '%c' must have 6 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
				}

				return r, fmt.Errorf("malformed path data: '%c' must have 6 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			curve = internal.NewBezier(
				cur,
				internal.Point{X: segment.Args[0], Y: segment.Args[1]},
				internal.Point{X: segment.Args[2], Y: segment.Args[3]},
				internal.Point{X: segment.Args[4], Y: segment.Args[5]},
			)

			l := curve.GetTotalLength()
			if l > 0 {
				r.Length += l
				r.parts = append(r.parts, curve)
				cur = internal.Point{X: cur.X + segment.Args[4], Y: cur.Y + segment.Args[5]}
			} else {
				r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X, cur.Y, cur.Y))
			}
		} else if segment.Command == 'S' {
			if len(segment.Args) != 4 {
				return r, fmt.Errorf("malformed path data: '%c' must have 4 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			prevSegment := segments[i-1]
			if i > 0 && (prevSegment.Command == 'C' || prevSegment.Command == 'c' || prevSegment.Command == 'S' || prevSegment.Command == 's') {
				if curve != internal.EmptyBezier {
					c := curve.GetC()
					curve = internal.NewBezier(
						cur,
						internal.Point{X: 2*cur.X - c.X, Y: 2*cur.Y - c.Y},
						internal.Point{X: segment.Args[0], Y: segment.Args[1]},
						internal.Point{X: segment.Args[2], Y: segment.Args[3]},
					)
				}
			} else {
				curve = internal.NewBezier(
					cur,
					cur,
					internal.Point{X: segment.Args[0], Y: segment.Args[1]},
					internal.Point{X: segment.Args[2], Y: segment.Args[3]},
				)
			}

			if curve != internal.EmptyBezier {
				r.Length += curve.GetTotalLength()
				r.parts = append(r.parts, curve)
				cur = internal.Point{X: segment.Args[2], Y: segment.Args[3]}
			}
		} else if segment.Command == 's' {
			if len(segment.Args) != 4 {
				return r, fmt.Errorf("malformed path data: '%c' must have 4 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			prevSegment := segments[i-1]
			if i > 0 && (prevSegment.Command == 'C' || prevSegment.Command == 'c' || prevSegment.Command == 'S' || prevSegment.Command == 's') {
				if curve != internal.EmptyBezier {
					c := curve.GetC()
					d := curve.GetD()
					curve = internal.NewBezier(
						cur,
						internal.Point{X: cur.X + d.X - c.X, Y: cur.Y + d.Y - c.Y},
						internal.Point{X: cur.X + segment.Args[0], Y: cur.Y + segment.Args[1]},
						internal.Point{X: cur.X + segment.Args[2], Y: cur.Y + segment.Args[3]},
					)
				} else {
					curve = internal.NewBezier(
						cur,
						cur,
						internal.Point{X: cur.X + segment.Args[0], Y: cur.Y + segment.Args[1]},
						internal.Point{X: cur.X + segment.Args[2], Y: cur.Y + segment.Args[3]},
					)
				}
			}

			if curve != internal.EmptyBezier {
				r.Length += curve.GetTotalLength()
				r.parts = append(r.parts, curve)
				cur = internal.Point{X: cur.X + segment.Args[2], Y: cur.Y + segment.Args[3]}
			}
		} else if segment.Command == 'Q' {
			if len(segment.Args) != 4 {
				return r, fmt.Errorf("malformed path data: '%c' must have 4 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			if cur.X == segment.Args[0] && cur.Y == segment.Args[1] {
				linearCurve := internal.NewLinear(segment.Args[0], segment.Args[2], segment.Args[1], segment.Args[3])
				r.Length += linearCurve.GetTotalLength()
				r.parts = append(r.parts, linearCurve)
			} else {
				curve = internal.NewBezier(
					cur,
					internal.Point{X: segment.Args[0], Y: segment.Args[1]},
					internal.Point{X: segment.Args[2], Y: segment.Args[3]},
					internal.EmptyPoint,
				)

				r.Length += curve.GetTotalLength()
				r.parts = append(r.parts, curve)
				cur = internal.Point{X: segment.Args[2], Y: segment.Args[3]}
			}

			cur = internal.Point{X: segment.Args[2], Y: segment.Args[3]}
			prevPoint = internal.Point{X: segment.Args[0], Y: segment.Args[1]}
		} else if segment.Command == 'q' {
			if len(segment.Args) != 4 {
				return r, fmt.Errorf("malformed path data: '%c' must have 4 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			if segment.Args[0] != 0 && segment.Args[1] != 0 {
				curve = internal.NewBezier(
					cur,
					internal.Point{X: cur.X + segment.Args[0], Y: cur.Y + segment.Args[1]},
					internal.Point{X: cur.X + segment.Args[2], Y: cur.Y + segment.Args[3]},
					internal.EmptyPoint,
				)

				r.Length += curve.GetTotalLength()
				r.parts = append(r.parts, curve)
			} else {
				linearCurve := internal.NewLinear(cur.X+segment.Args[0], cur.X+segment.Args[2], cur.Y+segment.Args[1], cur.Y+segment.Args[3])
				r.Length += linearCurve.GetTotalLength()
				r.parts = append(r.parts, linearCurve)
			}

			prevPoint = internal.Point{X: cur.X + segment.Args[0], Y: cur.Y + segment.Args[1]}
			cur = internal.Point{X: cur.X + segment.Args[2], Y: cur.Y + segment.Args[3]}
		} else if segment.Command == 'T' {
			if len(segment.Args) != 2 {
				return r, fmt.Errorf("malformed path data: '%c' must have 2 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			prevSegment := segments[i-1]
			if i > 0 && (prevSegment.Command == 'Q' || prevSegment.Command == 'q' || prevSegment.Command == 'T' || prevSegment.Command == 't') {
				if curve != internal.EmptyBezier {
					c := curve.GetC()
					curve = internal.NewBezier(
						cur,
						internal.Point{X: 2*cur.X - c.X, Y: 2*cur.Y - c.Y},
						internal.Point{X: segment.Args[0], Y: segment.Args[1]},
						internal.EmptyPoint,
					)
				}
			} else {
				curve = internal.NewBezier(
					cur,
					cur,
					internal.Point{X: segment.Args[0], Y: segment.Args[1]},
					internal.EmptyPoint,
				)
			}

			if curve != internal.EmptyBezier {
				r.Length += curve.GetTotalLength()
				r.parts = append(r.parts, curve)
				cur = internal.Point{X: segment.Args[0], Y: segment.Args[1]}
			}
		} else if segment.Command == 't' {
			if len(segment.Args) != 2 {
				return r, fmt.Errorf("malformed path data: '%c' must have 2 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			prevSegment := segments[i-1]
			if i > 0 && (prevSegment.Command == 'Q' || prevSegment.Command == 'q' || prevSegment.Command == 'T' || prevSegment.Command == 't') {
				curve = internal.NewBezier(
					cur,
					internal.Point{X: 2*cur.X - prevPoint.X, Y: 2*cur.Y - prevPoint.Y},
					internal.Point{X: segment.Args[0], Y: segment.Args[1]},
					internal.EmptyPoint,
				)

				r.Length += curve.GetTotalLength()
				r.parts = append(r.parts, curve)
			} else {
				linear := internal.NewLinear(
					cur.X,
					segment.Args[0],
					cur.Y,
					segment.Args[1],
				)

				r.Length += linear.GetTotalLength()
				r.parts = append(r.parts, linear)
			}

			prevPoint = internal.Point{X: 2*cur.X - prevPoint.X, Y: 2*cur.Y - prevPoint.Y}
			cur = internal.Point{X: segment.Args[0], Y: segment.Args[1]}
		} else if segment.Command == 'A' {
			if len(segment.Args) != 7 {
				return r, fmt.Errorf("malformed path data: '%c' must have 7 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			arc := internal.NewArc(
				cur.X,
				cur.Y,
				segment.Args[0],
				segment.Args[1],
				segment.Args[2],
				segment.Args[3] == 1,
				segment.Args[4] == 1,
				segment.Args[5],
				segment.Args[6],
			)

			r.Length += arc.GetTotalLength()
			r.parts = append(r.parts, arc)
			cur = internal.Point{X: segment.Args[5], Y: segment.Args[6]}
		} else if segment.Command == 'a' {
			if len(segment.Args) != 7 {
				return r, fmt.Errorf("malformed path data: '%c' must have 7 elements and has %d: '%s'", segment.Command, len(segment.Args), segment.Raw)
			}

			arc := internal.NewArc(
				cur.X,
				cur.Y,
				segment.Args[0],
				segment.Args[1],
				segment.Args[2],
				segment.Args[3] == 1,
				segment.Args[4] == 1,
				cur.X+segment.Args[5],
				cur.Y+segment.Args[6],
			)

			r.Length += arc.GetTotalLength()
			r.parts = append(r.parts, arc)
			cur = internal.Point{X: cur.X + segment.Args[5], Y: cur.Y + segment.Args[6]}
		}

		r.partialLengths = append(r.partialLengths, r.Length)
	}

	return r, nil
}

type sVGPathPart struct {
	fraction float64
	i        int
}

func (p SVGPath) getPartAtLength(pos float64) sVGPathPart {
	if pos < 0 {
		pos = 0
	} else if pos > p.Length {
		pos = p.Length
	}

	i := len(p.partialLengths) - 1
	for i > 0 && p.partialLengths[i] > pos {
		i--
	}
	i++

	return sVGPathPart{fraction: pos - p.partialLengths[i-1], i: i}
}

func (p SVGPath) GetTotalLength() float64 {
	return p.Length
}

func (p SVGPath) GetPointAtLength(pos float64) (point internal.Point, err error) {
	fractionPart := p.getPartAtLength(pos)
	functionAtPart := p.parts[fractionPart.i]

	if functionAtPart != nil {
		return functionAtPart.GetPointAtLength(fractionPart.fraction)
	} else if p.initialPoint != internal.EmptyPoint {
		return p.initialPoint, nil
	}

	return internal.EmptyPoint, fmt.Errorf("wrong function at this part")
}

func (p SVGPath) GetTangentAtLength(pos float64) (internal.Point, error) {
	fractionPart := p.getPartAtLength(pos)
	functionAtPart := p.parts[fractionPart.i]

	if functionAtPart != nil {
		return functionAtPart.GetTangentAtLength(fractionPart.fraction)
	} else if p.initialPoint != internal.EmptyPoint {
		return internal.EmptyPoint, nil
	}

	return internal.EmptyPoint, fmt.Errorf("wrong function at this part")
}

func (p SVGPath) GetPropertiesAtLength(pos float64) (internal.PointProperties, error) {
	fractionPart := p.getPartAtLength(pos)
	functionAtPart := p.parts[fractionPart.i]

	if functionAtPart != nil {
		return functionAtPart.GetPropertiesAtLength(fractionPart.fraction)
	} else if p.initialPoint != internal.EmptyPoint {
		return internal.PointProperties{
			X:        p.initialPoint.X,
			Y:        p.initialPoint.Y,
			TangentX: 0,
			TangentY: 0,
		}, nil
	}

	return internal.PointProperties{}, fmt.Errorf("wrong function at this part")
}

func (p SVGPath) GetParts() ([]internal.PartProperties, error) {
	parts := make([]internal.PartProperties, 0, len(p.parts))
	for i, part := range p.parts {
		if part != nil {
			p0, err := part.GetPointAtLength(0)
			if err != nil {
				return nil, err
			}

			p1, err := part.GetPointAtLength(p.partialLengths[i] - p.partialLengths[i-1])
			if err != nil {
				return nil, err
			}

			parts = append(parts, internal.NewPartProperties(
				p0,
				p1,
				p.partialLengths[i+1]-p.partialLengths[i],
				part.GetPointAtLength,
				part.GetTangentAtLength,
				part.GetPropertiesAtLength,
			))
		}
	}

	return parts, nil
}
