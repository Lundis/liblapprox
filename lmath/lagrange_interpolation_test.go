package lmath

import (
	"testing"
	"math"
	"code.google.com/p/liblundis/ltest"
	"fmt"
)

func assertLagrangeInterpolation(t *testing.T, x, y Vector) {
	lagrange := NewLagrangeInterpolationvv(x, y)
	L := lagrange.Function()
	for i := range x {
		// the lagrange interpolation should be equal to the source function for all x0 in x
		ltest.AssertEqualsFloat64(t, L(x[i]), y[i], "")
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

func TestGenerateChebyshevRoots(t *testing.T) {
	f := func(x float64) float64 {
        return math.Abs(x)
    }
    degree := 6
	x := GenerateChebyshevRoots(degree, -1, 1)
	y := Values(f, x)
	assertLagrangeInterpolation(t, x, y)

	// check that the roots are actually roots
	for i, xi := range x {
		ltest.AssertEqualsFloat64(t, math.Cos(float64(degree)*xi), 0, fmt.Sprintf("root %v, cos(%v) ", i, float64(degree)*xi))
	}
}

func TestGenerateEquiDistanceRoots(t *testing.T) {
	//roots := GenerateEquiDistanceRoots(3, -1, 1)
	
}