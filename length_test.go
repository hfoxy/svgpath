package svgpath

import "testing"

func TestLineToLength(t *testing.T) {
	testPathLength(t, "m0,0l10,0", 10)
	testPathLength(t, "M0,0L10,0", 10)
	testPathLength(t, "M0,0L10,0M0,0L10,0", 20)
	testPathLength(t, "M0,0L10,0m0,0L10,0", 10)
	testPathLength(t, "M0,0L10,0l10,0", 20)
}

func TestHandVLength(t *testing.T) {
	testPathLength(t, "M0,0h10", 10)
	testPathLength(t, "M50,0H40", 10)
	testPathLength(t, "M0,0v10", 10)
	testPathLength(t, "M0,50V40", 10)
}

func TestZLength(t *testing.T) {
	testPathLength(t, "m0,0h10z", 20)
	testPathLength(t, "m0,0h10Z", 20)
}

func TestCubicBezierLength(t *testing.T) {
	testPathLength(t, "M100,25C10,90,110,100,150,195", 213.8)
	testPathLength(t, "m100,25c-90,65,10,75,50,170", 213.8)
	testPathLength(t, "M100,200 C100,100 250,100 250,200 S400,300 400,200", 475.746)
	testPathLength(t, "M100,200 c0,-100 150,-100 150,0 s150,100 150,0", 475.746)
	testPathLength(t, "M100,200 S400,300 400,200", 327.9618)
	testPathLength(t, "M100,200 s300,100 300,0", 327.9618)
}

func TestQuadraticBezierLength(t *testing.T) {
	testPathLength(t, "M200,300 Q400,50 600,300", 487.77)
	testPathLength(t, "M200,300 q200,-250 400,0", 487.77)
	testPathLength(t, "M0,100 Q50,-50 100,100 T200,100", 376.84)
	testPathLength(t, "M0,100 q50,-150 100,0 t100,0", 376.84)
	testPathLength(t, "M0,100 Q50,-50 100,100 T200,100 T300,100", 565.26)
	testPathLength(t, "M0,100 T200,100", 200)
	testPathLength(t, "M0,100 t200,100", 223.606)
}

func testPathLength(t *testing.T, path string, expected float64) {
	properties, err := NewFromPath(path)
	if err != nil {
		t.Errorf("expected no error (%s), got %v", path, err)
	}

	if !inDelta(properties.GetTotalLength(), expected, epsilon) {
		t.Errorf("expected path (%s) length to be %f, got %f", path, expected, properties.GetTotalLength())
	}
}
