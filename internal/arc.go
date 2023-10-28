package internal

import "math"

type Arc struct {
	Part
	x0            float64
	y0            float64
	rx            float64
	ry            float64
	xAxisRotation float64
	largeArcFlag  bool
	sweepFlag     bool
	x1            float64
	y1            float64
	length        float64
}

func NewArc(x0, y0, rx, ry, xAxisRotation float64, largeArcFlag bool, sweepFlag bool, x1, y1 float64) *Arc {
	length := approximateArcLengthOfCurve(300, func(t float64) Point {
		return pointOnEllipticalArc(
			Point{X: x0, Y: y0},
			Point{X: rx, Y: ry},
			xAxisRotation,
			largeArcFlag,
			sweepFlag,
			Point{X: x1, Y: y1},
			t,
		).ToPoint()
	})

	return &Arc{
		x0:            x0,
		y0:            y0,
		rx:            rx,
		ry:            ry,
		xAxisRotation: xAxisRotation,
		largeArcFlag:  largeArcFlag,
		sweepFlag:     sweepFlag,
		x1:            x1,
		y1:            y1,
		length:        length,
	}
}

func (a Arc) GetTotalLength() float64 {
	return a.length
}

func (a Arc) GetPointAtLength(pos float64) Point {
	point := pointOnEllipticalArc(
		Point{X: a.x0, Y: a.y0},
		Point{X: a.rx, Y: a.ry},
		a.xAxisRotation,
		a.largeArcFlag,
		a.sweepFlag,
		Point{X: a.x1, Y: a.y1},
		pos/a.length,
	)

	return point.ToPoint()
}

func (a Arc) GetTangentAtLength(fractionLength float64) Point {
	if fractionLength < 0 {
		fractionLength = 0
	} else if fractionLength > a.length {
		fractionLength = a.length
	}

	// there's a note on this in the original source that this needs testing
	pointDist := .05
	p1 := a.GetPointAtLength(fractionLength - pointDist)

	var p2 Point
	if fractionLength < a.length-pointDist {
		p2 = a.GetPointAtLength(fractionLength - pointDist)
	} else {
		p2 = a.GetPointAtLength(fractionLength + pointDist)
	}

	xDist := p2.X - p1.X
	yDist := p2.Y - p1.Y
	dist := math.Sqrt(xDist*xDist + yDist*yDist)

	if fractionLength < a.length-pointDist {
		return Point{X: -xDist / dist, Y: -yDist / dist}
	} else {
		return Point{X: xDist / dist, Y: yDist / dist}
	}
}

func (a Arc) GetPropertiesAtLength(fractionLength float64) PointProperties {
	tangent := a.GetTangentAtLength(fractionLength)
	point := a.GetPointAtLength(fractionLength)

	return PointProperties{
		X:        point.X,
		Y:        point.Y,
		TangentX: tangent.X,
		TangentY: tangent.Y,
	}
}

type PointOnEllipticalArc struct {
	x                       float64
	y                       float64
	ellipticalArcStartAngle float64
	ellipticalArcEndAngle   float64
	ellipticalArcAngle      float64
	ellipticalArcCenter     Point
	resultantRx             float64
	resultantRy             float64
}

func (p PointOnEllipticalArc) ToPoint() Point {
	return Point{X: p.x, Y: p.y}
}

func approximateArcLengthOfCurve(numSteps int, pointOnCurve func(float64) Point) float64 {
	var length float64
	return length
}

func pointOnEllipticalArc(p0 Point, r Point, xAxisRotation float64, largeArcFlag bool, sweepFlag bool, p1 Point, t float64) PointOnEllipticalArc {
	rx := math.Abs(r.X)
	ry := math.Abs(r.Y)

	xAxisRotation = math.Mod(xAxisRotation, 360)
	xAxisRotationRadius := toRadians(xAxisRotation)

	if p0.X == p1.X && p0.Y == p1.Y {
		return PointOnEllipticalArc{x: p0.X, y: p0.Y, ellipticalArcAngle: 0}
	}

	if rx == 0 || ry == 0 {
		return PointOnEllipticalArc{x: 0, y: 0, ellipticalArcAngle: 0}
	}

	dx := (p0.X - p1.X) / 2
	dy := (p0.Y - p1.Y) / 2

	transformedPoint := Point{
		X: math.Cos(xAxisRotationRadius)*dx + math.Sin(xAxisRotationRadius)*dy,
		Y: -math.Sin(xAxisRotationRadius)*dx + math.Cos(xAxisRotationRadius)*dy,
	}

	radiiCheck := math.Pow(transformedPoint.X, 2)/math.Pow(rx, 2) + math.Pow(transformedPoint.Y, 2)/math.Pow(ry, 2)
	if radiiCheck > 1 {
		rx *= math.Sqrt(radiiCheck)
		ry *= math.Sqrt(radiiCheck)
	}

	cSquareNumerator := math.Pow(rx, 2)*math.Pow(ry, 2) - math.Pow(rx, 2)*math.Pow(transformedPoint.Y, 2) - math.Pow(ry, 2)*math.Pow(transformedPoint.X, 2)
	cSquareRootDenom := math.Pow(rx, 2)*math.Pow(transformedPoint.Y, 2) + math.Pow(ry, 2)*math.Pow(transformedPoint.X, 2)
	cRadicand := cSquareNumerator / cSquareRootDenom

	if cRadicand < 0 {
		cRadicand = 0
	}

	var cCoef float64
	if largeArcFlag == sweepFlag {
		cCoef = -1
	} else {
		cCoef = 1
	}

	cCoef *= math.Sqrt(cRadicand)

	transformedCenter := Point{
		X: cCoef * ((rx * transformedPoint.Y) / ry),
		Y: cCoef * (-(ry * transformedPoint.X) / rx),
	}

	center := Point{
		X: math.Cos(xAxisRotationRadius)*transformedCenter.X - math.Sin(xAxisRotationRadius)*transformedCenter.Y + (p0.X+p1.X)/2,
		Y: math.Sin(xAxisRotationRadius)*transformedCenter.X + math.Cos(xAxisRotationRadius)*transformedCenter.Y + (p0.Y+p1.Y)/2,
	}

	startVector := Point{
		X: (transformedPoint.X - transformedCenter.X) / rx,
		Y: (transformedPoint.Y - transformedCenter.Y) / ry,
	}

	startAngle := angleBetween(Point{X: 1, Y: 0}, startVector)

	endVector := Point{
		X: (-transformedPoint.X - transformedCenter.X) / rx,
		Y: (-transformedPoint.Y - transformedCenter.Y) / ry,
	}

	sweepAngle := angleBetween(startVector, endVector)

	if !sweepFlag && sweepAngle > 0 {
		sweepAngle -= 2 * math.Pi
	} else if sweepFlag && sweepAngle < 0 {
		sweepAngle += 2 * math.Pi
	}

	// TODO: validate
	sweepAngle = math.Mod(sweepAngle, 2*math.Pi)

	angle := startAngle + sweepAngle*t
	ellipseComponentX := rx * math.Cos(angle)
	ellipseComponentY := ry * math.Sin(angle)

	point := PointOnEllipticalArc{
		x:                       math.Cos(xAxisRotationRadius)*ellipseComponentX - math.Sin(xAxisRotationRadius)*ellipseComponentY + center.X,
		y:                       math.Sin(xAxisRotationRadius)*ellipseComponentX + math.Cos(xAxisRotationRadius)*ellipseComponentY + center.Y,
		ellipticalArcStartAngle: startAngle,
		ellipticalArcEndAngle:   startAngle + sweepAngle,
		ellipticalArcAngle:      angle,
		ellipticalArcCenter:     center,
		resultantRx:             rx,
		resultantRy:             ry,
	}

	return point
}
