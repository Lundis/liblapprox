package ipol

import (
	"testing"
	"math"
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/base"
	. "code.google.com/p/liblundis/lmath/base/poly"
	"fmt"
)

func TestGenerateChebyshevRoots(t *testing.T) {
	f := func(x float64) float64 {
        return math.Abs(x)
    }
    degree := 11
	x := GenerateChebyshevRoots(degree, -1, 1)
	y := Values(f, x)
	assertLagrangeInterpolation(t, x, y)

	poly11 := Poly{BasisImpl{V:[]float64{0, -11, 0, 220, 0, -1232, 0, 2816, 0, -2816, 0, 1024}}}

	// check that the roots are actually roots
	for i, xi := range x {
		AssertEqualsFloat64(t, poly11.ValueAt(xi), 0, fmt.Sprintf("root %v, Tn(%v) == %v ", i, xi, poly11.ValueAt(xi)))
	}
}

// Tests that the roots are in order
func TestGenerateChebyshevRoots_InOrder(t *testing.T) {
	roots := GenerateChebyshevRoots(3, -1, 1)
	for i := 0; i < len(roots)-1; i++ {
		if roots[i] > roots[i+1] {
			t.Errorf("roots are not in order")
		}
	}
}

func TestGenerateEquiDistanceRoots(t *testing.T) {
	roots := GenerateEquiDistanceRoots(65, -1, 1)
	sum := float64(0)
	for _, v := range roots {
		sum += v
	}
	AssertEqualsFloat64(t, sum, 0, "sum not equal to zero")
}