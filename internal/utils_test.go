package internal

import "math"

const epsilon = 0.000001

func inDelta(a float64, b float64, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
