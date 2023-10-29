package internal

import "testing"

func TestLinearDirectly(t *testing.T) {
	linear := NewLinear(0, 0, 10, 0)

	if !inDelta(linear.GetTotalLength(), 10, epsilon) {
		t.Errorf("Expected linear.GetTotalLength() to be 10, got %v", linear.GetTotalLength())
	}
}
