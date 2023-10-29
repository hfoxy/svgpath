package internal

import (
	"fmt"
	"math"
)

type BezierType string

const (
	BezierTypeCubic     BezierType = "cubic"
	BezierTypeQuadratic BezierType = "quadratic"
	BezierTypeEmpty     BezierType = "empty"
)

type Bezier struct {
	a          Point
	b          Point
	c          Point
	d          Point
	bezierType BezierType
	length     float64
}

var EmptyBezier = Bezier{
	a:          EmptyPoint,
	b:          EmptyPoint,
	c:          EmptyPoint,
	d:          EmptyPoint,
	bezierType: BezierTypeEmpty,
	length:     0,
}

func NewBezier(a Point, b Point, c Point, d Point) Bezier {
	r := Bezier{
		a: a,
		b: b,
		c: c,
	}

	if !d.blank {
		r.bezierType = BezierTypeCubic
		r.d = d
	} else {
		r.bezierType = BezierTypeQuadratic
		r.d = Point{X: 0, Y: 0}
	}

	r.length, _ = r.GetArcLength(
		[]float64{a.X, b.X, c.X, d.X},
		[]float64{a.Y, b.Y, c.Y, d.Y},
		1,
	)

	return r
}

func (b Bezier) GetArcLength(xs []float64, ys []float64, t float64) (float64, error) {
	if b.bezierType == BezierTypeCubic {
		return getCubicArcLength(xs, ys, t), nil
	} else if b.bezierType == BezierTypeQuadratic {
		return getQuadraticArcLength(xs, ys, t), nil
	}

	return 0, fmt.Errorf("invalid bezier type: %s", b.bezierType)
}

func (b Bezier) GetPoint(xs []float64, ys []float64, t float64) (Point, error) {
	if b.bezierType == BezierTypeCubic {
		return cubicPoint(xs, ys, t), nil
	} else if b.bezierType == BezierTypeQuadratic {
		return quadraticPoint(xs, ys, t), nil
	}

	return EmptyPoint, fmt.Errorf("invalid bezier type: %s", b.bezierType)
}

func (b Bezier) GetDerivative(xs []float64, ys []float64, t float64) (Point, error) {
	if b.bezierType == BezierTypeCubic {
		return cubicDerivative(xs, ys, t), nil
	} else if b.bezierType == BezierTypeQuadratic {
		return quadraticDerivative(xs, ys, t), nil
	}

	return EmptyPoint, fmt.Errorf("invalid bezier type: %s", b.bezierType)
}

func (b Bezier) GetTotalLength() float64 {
	return b.length
}

func (b Bezier) GetPointAtLength(pos float64) (Point, error) {
	xs := []float64{b.a.X, b.b.X, b.c.X, b.d.X}
	ys := []float64{b.a.Y, b.b.Y, b.c.Y, b.d.Y}
	t, err := t2length(pos, b.length, func(t float64) (float64, error) {
		return b.GetArcLength(xs, ys, t)
	})

	if err != nil {
		return EmptyPoint, err
	}

	return b.GetPoint(xs, ys, t)
}

func (b Bezier) GetTangentAtLength(pos float64) (Point, error) {
	xs := []float64{b.a.X, b.b.X, b.c.X, b.d.X}
	ys := []float64{b.a.Y, b.b.Y, b.c.Y, b.d.Y}
	t, err := t2length(pos, b.length, func(t float64) (float64, error) {
		return b.GetArcLength(xs, ys, t)
	})

	if err != nil {
		return EmptyPoint, err
	}

	derivative, err := b.GetDerivative(xs, ys, t)
	if err != nil {
		return EmptyPoint, err
	}

	mdl := math.Sqrt(derivative.X*derivative.X + derivative.Y*derivative.Y)

	var tangent Point
	if mdl > 0 {
		tangent = Point{X: derivative.X / mdl, Y: derivative.Y / mdl}
	} else {
		tangent = Point{X: 0, Y: 0}
	}

	return tangent, nil
}

func (b Bezier) GetPropertiesAtLength(pos float64) (PointProperties, error) {
	xs := []float64{b.a.X, b.b.X, b.c.X, b.d.X}
	ys := []float64{b.a.Y, b.b.Y, b.c.Y, b.d.Y}
	t, err := t2length(pos, b.length, func(t float64) (float64, error) {
		return b.GetArcLength(xs, ys, t)
	})

	if err != nil {
		return PointProperties{}, err
	}

	derivative, err := b.GetDerivative(xs, ys, t)
	if err != nil {
		return PointProperties{}, err
	}

	mdl := math.Sqrt(derivative.X*derivative.X + derivative.Y*derivative.Y)

	var tangent Point
	if mdl > 0 {
		tangent = Point{X: derivative.X / mdl, Y: derivative.Y / mdl}
	} else {
		tangent = Point{X: 0, Y: 0}
	}

	point, err := b.GetPoint(xs, ys, t)
	if err != nil {
		return PointProperties{}, err
	}

	return PointProperties{
		X:        point.X,
		Y:        point.Y,
		TangentX: tangent.X,
		TangentY: tangent.Y,
	}, nil
}

func (b Bezier) GetC() Point {
	return b.c
}

func (b Bezier) GetD() Point {
	return b.d
}
