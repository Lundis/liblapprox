package ipol

import (
	"testing"
	"math"
	. "code.google.com/p/liblundis/lmath"
)

func assertLagrangeInterpolation(t *testing.T, x, y Vector) {
	lagrange := NewLagrangeInterpolationvv(x, y)
	L := lagrange.Function()
	for i := range x {
		// the lagrange interpolation should be equal to the source function for all x0 in x
		AssertEqualsFloat64(t, L(x[i]), y[i], "")
	}
}

func TestNewLagrangeInterpolationfv(t *testing.T) {
	f := func(x float64) float64 {
		return math.Pow(math.E, 2)
	}
	x := Vector{-1, -0.2, 0.1, 0.4, 1}
	y := Values(f, x)
	assertLagrangeInterpolation(t, x, y)
}

