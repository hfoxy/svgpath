package internal

type Properties interface {
	GetTotalLength() float64
	GetPointAtLength(float64) Point
	GetTangentAtLength(float64) Point
	GetPropertiesAtLength(float64) PointProperties
}

type PartProperties struct {
	Start  Point
	End    Point
	Length float64
}

type Part interface {
	GetPointAtLength(pos float64) Point
	GetTangentAtLength(pos float64) Point
	GetPropertiesAtLength(pos float64) PointProperties
}

var EmptyPoint = Point{blank: true}

type Point struct {
	X     float64
	Y     float64
	blank bool
}

type PointProperties struct {
	X        float64
	Y        float64
	TangentX float64
	TangentY float64
}
