package internal

import "math"

func toRadians(angle float64) float64 {
	return angle * (math.Pi / 180)
}

func angleBetween(v0 Point, v1 Point) float64 {
	p := v0.X*v1.X + v0.Y*v1.Y
	n := math.Sqrt((v0.X*v0.X + v0.Y*v0.Y) * (v1.X*v1.X + v1.Y*v1.Y))
	var sign float64
	if v0.X*v1.Y-v0.Y*v1.X < 0 {
		sign = -1
	} else {
		sign = 1
	}

	return sign * math.Acos(p/n)
}
