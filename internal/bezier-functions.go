package internal

import "math"

func cubicPoint(xs []float64, ys []float64, t float64) Point {
	x := (1-t)*(1-t)*(1-t)*xs[0] +
		3*(1-t)*(1-t)*t*xs[1] +
		3*(1-t)*t*t*xs[2] +
		t*t*t*xs[3]

	y := (1-t)*(1-t)*(1-t)*ys[0] +
		3*(1-t)*(1-t)*t*ys[1] +
		3*(1-t)*t*t*ys[2] +
		t*t*t*ys[3]

	return Point{X: x, Y: y}
}

func cubicDerivative(xs []float64, ys []float64, t float64) Point {
	return quadraticPoint(
		[]float64{3 * (xs[1] - xs[0]), 3 * (xs[2] - xs[1]), 3 * (xs[3] - xs[2])},
		[]float64{3 * (ys[1] - ys[0]), 3 * (ys[2] - ys[1]), 3 * (ys[3] - ys[2])},
		t,
	)
}

func getCubicArcLength(xs []float64, ys []float64, t float64) float64 {
	n := float64(20)

	z := t / n
	sum := float64(0)
	var correctedT float64

	nInt := int(n)
	for i := 0; i < int(n); i++ {
		correctedT = tValues[nInt][i] + z
		sum += cValues[nInt][i] * bFunc(xs, ys, correctedT)
	}

	return z * sum
}

func quadraticPoint(xs []float64, ys []float64, t float64) Point {
	x := (1-t)*(1-t)*xs[0] + 2*(1-t)*t*xs[1] + t*t*xs[2]
	y := (1-t)*(1-t)*ys[0] + 2*(1-t)*t*ys[1] + t*t*ys[2]
	return Point{X: x, Y: y}
}

func getQuadraticArcLength(xs []float64, ys []float64, t float64) float64 {
	ax := xs[0] - 2*xs[1] + xs[2]
	ay := ys[0] - 2*ys[1] + ys[2]
	bx := 2*xs[1] - 2*xs[0]
	by := 2*ys[1] - 2*ys[0]

	A := 4 * (ax*ax + ay*ay)
	B := 4 * (ax*bx + ay*by)
	C := bx*bx + by*by

	if A == 0 {
		return t * math.Sqrt(math.Pow(xs[2]-xs[0], 2)+math.Pow(ys[2]-ys[0], 2))
	}

	b := B / (2 * A)
	c := C / A
	u := t + b
	k := c - b*b

	var uuk float64
	if u*u+k > 0 {
		uuk = math.Sqrt(u*u + k)
	} else {
		uuk = 0
	}

	var bbk float64
	if b*b+k > 0 {
		bbk = math.Sqrt(b*b + k)
	} else {
		bbk = 0
	}

	var term float64
	if b+math.Sqrt(b*b+k) != 0 && ((u+uuk)/b+bbk) != 0 {
		term = math.Log((u + uuk) / (b + bbk))
	} else {
		term = 0
	}

	return (math.Sqrt(A) / 2) * (u*uuk - b*bbk + k*term)
}

func quadraticDerivative(xs []float64, ys []float64, t float64) Point {
	return Point{
		X: (1-t)*2*(xs[1]-xs[0]) + t*2*(xs[2]-xs[1]),
		Y: (1-t)*2*(ys[1]-ys[0]) + t*2*(ys[2]-ys[1]),
	}
}

func bFunc(xs []float64, ys []float64, t float64) float64 {
	xbase := getDerivative(1, t, xs)
	ybase := getDerivative(1, t, ys)
	return math.Sqrt(xbase*xbase + ybase*ybase)
}

func getDerivative(derivative float64, t float64, vs []float64) float64 {
	n := len(vs) - 1

	if n == 0 {
		return 0
	}

	if derivative == 0 {
		value := float64(0)
		for k := 0; k <= n; k++ {
			value += binomialCoefficients[n][k] * math.Pow(1-t, float64(n-k)) * math.Pow(t, float64(k)) * vs[k]
		}

		return value
	} else {
		vs2 := make([]float64, n)
		for k := 0; k < n; k++ {
			vs2[k] = float64(n) * (vs[k+1] - vs[k])
		}

		return getDerivative(derivative-1, t, vs2)
	}
}

func t2length(length float64, totalLength float64, f func(t float64) float64) float64 {
	e := float64(1)
	t := length / totalLength
	step := (length - f(t)) / totalLength

	numIterations := 0
	for e > 0.001 {
		increasedTLength := f(t + step)
		increasedTError := math.Abs(length-increasedTLength) / totalLength
		if increasedTError < e {
			e = increasedTError
			t += step
		} else {
			decreasedTLength := f(t - step)
			decreasedTError := math.Abs(length-decreasedTLength) / totalLength
			if decreasedTError < e {
				e = decreasedTError
				t -= step
			} else {
				step /= 2
			}
		}

		numIterations++
		if numIterations > 500 {
			break
		}
	}

	return t
}
