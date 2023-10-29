package internal

import "testing"

func TestBezierLengthQuadratic(t *testing.T) {
	bezier := NewBezier(
		Point{X: 200, Y: 300},
		Point{X: 400, Y: 50},
		Point{X: 600, Y: 300},
		EmptyPoint,
	)

	if !inDelta(bezier.GetTotalLength(), 487.7710938955238, epsilon) {
		t.Errorf("Expected bezier.GetTotalLength() to be 487.7710938955238, got %v", bezier.GetTotalLength())
	}
}

func TestBezierLengthCubic(t *testing.T) {
	bezier := NewBezier(
		Point{X: 200, Y: 200},
		Point{X: 275, Y: 100},
		Point{X: 575, Y: 100},
		Point{X: 500, Y: 200},
	)

	if !inDelta(bezier.GetTotalLength(), 383.4438582551745, epsilon) {
		t.Errorf("Expected bezier.GetTotalLength() to be 383.4438582551745, got %v", bezier.GetTotalLength())
	}
}

func TestGetPointAtLengthQuadratic(t *testing.T) {
	bezier := NewBezier(
		Point{X: 200, Y: 300},
		Point{X: 400, Y: 50},
		Point{X: 600, Y: 300},
		EmptyPoint,
	)

	point, err := bezier.GetPointAtLength(487.77 / 6)
	if err != nil {
		t.Errorf("Expected bezier.GetPointAtLength(487.77/6) to not return an error, got %v", err)
	}

	if !inDelta(point.X, 255.03382490461, epsilon) {
		t.Errorf("Expected point.X to be 255.24, got %v", point.X)
	}

	if !inDelta(point.Y, 240.67247475558528, epsilon) {
		t.Errorf("Expected point.Y to be 240.67247475558528, got %v", point.Y)
	}
}

func TestGetPointAtLengthCubic(t *testing.T) {
	bezier := NewBezier(
		Point{X: 200, Y: 200},
		Point{X: 275, Y: 100},
		Point{X: 575, Y: 100},
		Point{X: 500, Y: 200},
	)

	point, err := bezier.GetPointAtLength(383.44 / 6)
	if err != nil {
		t.Errorf("Expected bezier.GetPointAtLength(383.44/6) to not return an error, got %v", err)
	}

	if !inDelta(point.X, 249.42120403254987, epsilon) {
		t.Errorf("Expected point.X to be 249.42120403254987, got %v", point.X)
	}

	if !inDelta(point.Y, 160.41015379568006, epsilon) {
		t.Errorf("Expected point.Y to be 160.41015379568006, got %v", point.Y)
	}
}
