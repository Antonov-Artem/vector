package vector

import "testing"

func TestVector(t *testing.T) {
	expect := vector64{
		X: 1,
		Y: 1,
		Z: 1,
	}
	result := Vector(1, 1, 1)

	if expect != result {
		t.Errorf("Expected: (%v); Result: (%v)", expect, result)
	}
}
