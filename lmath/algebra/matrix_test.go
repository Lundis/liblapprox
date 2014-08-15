package algebra

import (
	"testing"
)

func TestSolve(t *testing.T) {
	m := NewMatrix(2, 3)
	copy(m[0], []float64{1, 2, 3})
	copy(m[1], []float64{4, 5, 6})
	m.Solve()
	expected0 := []float64{1, 0, -1}
	expected1 := []float64{0, 1, 2}
	for i, v := range m[0] {
		if v != expected0[i] {
			t.Errorf("got %v. Expected %v", v, expected0[i])
		}
	}
	for i, v := range m[1] {
		if v != expected1[i] {
			t.Errorf("got %v. Expected %v", v, expected1[i])
		}
	}
}