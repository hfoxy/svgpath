package internal

import "math"

type Linear struct {
	Part
	x0      float64
	x1      float64
	y0      float64
	y1      float64
	length  float64
	tangent Point
}

func NewLinear(x0 float64, y0 float64, x1 float64, y1 float64) Linear {
	length := math.Sqrt(math.Pow(x0-x1, 2) + math.Pow(y0-y1, 2))
	tangent := Point{
		X: (x1 - x0) / length,
		Y: (y1 - y0) / length,
	}

	return Linear{
		x0:      x0,
		x1:      x1,
		y0:      y0,
		y1:      y1,
		length:  math.Sqrt(math.Pow(x0-x1, 2) + math.Pow(y0-y1, 2)),
		tangent: tangent,
	}
}

func (l Linear) GetTotalLength() float64 {
	return l.length
}

func (l Linear) GetPointAtLength(pos float64) (Point, error) {
	fraction := pos / l.length

	newDeltaX := (l.x1 - l.x0) * fraction
	newDeltaY := (l.y1 - l.y0) * fraction
	return Point{
		X: l.x0 + newDeltaX,
		Y: l.y0 + newDeltaY,
	}, nil
}

func (l Linear) GetTangentAtLength(_ float64) (Point, error) {
	return l.tangent, nil
}

func (l Linear) GetPropertiesAtLength(pos float64) (PointProperties, error) {
	point, _ := l.GetPointAtLength(pos)
	return PointProperties{
		X:        point.X,
		Y:        point.Y,
		TangentX: l.tangent.X,
		TangentY: l.tangent.Y,
	}, nil
}
