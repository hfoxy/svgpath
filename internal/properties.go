package internal

type Properties interface {
	GetTotalLength() float64
	GetPointAtLength(float64) (Point, error)
	GetTangentAtLength(float64) (Point, error)
	GetPropertiesAtLength(float64) (PointProperties, error)
}

type PartProperties struct {
	Start  Point
	End    Point
	Length float64
	part   Part
}

func NewPartProperties(
	start Point,
	end Point,
	length float64,
	part Part,
) PartProperties {
	return PartProperties{
		Start:  start,
		End:    end,
		Length: length,
		part:   part,
	}
}

func (p PartProperties) GetTotalLength() float64 {
	return p.Length
}

func (p PartProperties) GetPointAtLength(pos float64) (Point, error) {
	return p.part.GetPointAtLength(pos)
}

func (p PartProperties) GetTangentAtLength(pos float64) (Point, error) {
	return p.part.GetTangentAtLength(pos)
}

func (p PartProperties) GetPropertiesAtLength(pos float64) (PointProperties, error) {
	return p.part.GetPropertiesAtLength(pos)
}

type Part interface {
	GetPointAtLength(pos float64) (Point, error)
	GetTangentAtLength(pos float64) (Point, error)
	GetPropertiesAtLength(pos float64) (PointProperties, error)
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
