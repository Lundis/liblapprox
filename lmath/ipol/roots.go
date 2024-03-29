package ipol

import (
	"math"
)

func GenerateChebyshevRoots(degree int, start, end float64) []float64 {
	roots := make([]float64, degree)
	for k := 0; k < degree; k++ {
		roots[k] = math.Cos(-math.Pi + math.Pi/2 * float64(2 * k + 1)/ float64(degree))
		// adjust to interval
		roots[k] = (roots[k] * (end - start) + start + end)/2
	}
	return roots
}

func GenerateEquiDistanceRoots(degree int, start, end float64) []float64 {
	roots := make([]float64, degree)
	if (degree == 1) {
		roots[0] = (start + end)/2
		return roots
	}
	for k := 0; k < degree; k++ {
		roots[k] = start + (end - start)*float64(k)/float64(degree-1)
	}
	return roots
}