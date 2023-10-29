package svgpath

import (
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

	var curve internal.Bezier
	ringStart := internal.Point{X: 0, Y: 0}

	for i, segment := range segments {
		if segment.Command == 'M' {
			cur = internal.Point{X: segment.Args[0], Y: segment.Args[1]}
			ringStart = cur

			r.parts = append(r.parts, nil)
			if i == 0 {
				r.initialPoint = cur
			}
		} else if segment.Command == 'm' {
			cur = internal.Point{X: segment.Args[0] + cur.X, Y: segment.Args[1] + cur.Y}
			ringStart = cur
			r.parts = append(r.parts, nil)
		} else if segment.Command == 'L' {
			r.Length += math.Sqrt(math.Pow(cur.X-segment.Args[1], 2) + math.Pow(cur.Y-segment.Args[2], 2))
			r.parts = append(r.parts, internal.NewLinear(cur.X, segment.Args[1], cur.Y, segment.Args[2]))
			cur = internal.Point{X: segment.Args[1], Y: segment.Args[2]}
		} else if segment.Command == 'l' {
			r.Length += math.Sqrt(math.Pow(segment.Args[1], 2) + math.Pow(segment.Args[2], 2))
			r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X+segment.Args[1], cur.Y, cur.Y+segment.Args[2]))
			cur = internal.Point{X: cur.X + segment.Args[1], Y: cur.Y + segment.Args[2]}
		} else if segment.Command == 'H' {
			r.Length += math.Abs(cur.X - segment.Args[1])
			r.parts = append(r.parts, internal.NewLinear(cur.X, segment.Args[1], cur.Y, cur.Y))
			cur = internal.Point{X: segment.Args[1], Y: cur.Y}
		} else if segment.Command == 'h' {
			r.Length += math.Abs(segment.Args[1])
			r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X+segment.Args[1], cur.Y, cur.Y))
			cur = internal.Point{X: cur.X + segment.Args[1], Y: cur.Y}
		} else if segment.Command == 'V' {
			r.Length += math.Abs(cur.Y - segment.Args[1])
			r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X, cur.Y, segment.Args[1]))
			cur = internal.Point{X: cur.X, Y: segment.Args[1]}
		} else if segment.Command == 'v' {
			r.Length += math.Abs(segment.Args[1])
			r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X, cur.Y, cur.Y+segment.Args[1]))
			cur = internal.Point{X: cur.X, Y: cur.Y + segment.Args[1]}
		} else if segment.Command == 'z' || segment.Command == 'Z' {
			r.Length += math.Sqrt(math.Pow(ringStart.X-cur.X, 2) + math.Pow(ringStart.Y-cur.Y, 2))
			r.parts = append(r.parts, internal.NewLinear(cur.X, ringStart.X, cur.Y, ringStart.Y))
			cur = ringStart
		} else if segment.Command == 'C' {
			curve = internal.NewBezier(
				cur,
				internal.Point{X: segment.Args[1], Y: segment.Args[2]},
				internal.Point{X: segment.Args[3], Y: segment.Args[4]},
				internal.Point{X: segment.Args[5], Y: segment.Args[6]},
			)

			r.Length += curve.GetTotalLength()
			r.parts = append(r.parts, curve)
			cur = internal.Point{X: segment.Args[5], Y: segment.Args[6]}
		} else if segment.Command == 'c' {
			curve = internal.NewBezier(
				cur,
				internal.Point{X: segment.Args[1], Y: segment.Args[2]},
				internal.Point{X: segment.Args[3], Y: segment.Args[4]},
				internal.Point{X: segment.Args[5], Y: segment.Args[6]},
			)

			l := curve.GetTotalLength()
			if l > 0 {
				r.Length += l
				r.parts = append(r.parts, curve)
				cur = internal.Point{X: cur.X + segment.Args[5], Y: cur.Y + segment.Args[6]}
			} else {
				r.parts = append(r.parts, internal.NewLinear(cur.X, cur.X, cur.Y, cur.Y))
			}
		}
	}

	return r, nil
}
