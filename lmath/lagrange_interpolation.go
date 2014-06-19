package lmath

import(
	"fmt"
	"math"
)

type LagrangeInterpolation struct {
	x Vector
	y Vector
	bases []Polynomial
}

func NewLagrangeInterpolationvv(x, y Vector) LagrangeInterpolation {
	if len(x) != len(y) {
		fmt.Printf("NewLagrangeInterpolation error: len(x) != len(y)")
	} else if len(x) < 1 {
		fmt.Printf("NewLagrangeInterpolation error: can't be of degree zero")
	}
	lagrange := LagrangeInterpolation{make([]float64, len(x)), make([]float64, len(x)), make([]Polynomial, len(x))}
	copy(lagrange.x, x)
	copy(lagrange.y, y)
	lagrange.generateBases()
	return lagrange
}

func NewLagrangeInterpolationfv(f Func1to1, x Vector) LagrangeInterpolation {
	y := Values(f, x)
	return NewLagrangeInterpolationvv(x, y)
}

func (self *LagrangeInterpolation) generateBases() {
	for i := 0; i < len(self.bases); i++ {
		poly := Polynomial{1}
		for j := 0; j < len(self.x); j++ {
			if i != j {
				poly = poly.Mult(Polynomial{-self.x[j], 1})
			}
		}
		self.bases[i] = poly.MultConstant(1.0 / poly.ValueAt(self.x[i]))
	}
}

func (self LagrangeInterpolation) Bases() []Polynomial {
	return self.bases
}

func (self LagrangeInterpolation) Function() Func1to1 {
	return func(x float64) float64 {
		sum := float64(0)
		for i := 0; i < len(self.bases); i++ {
			li := self.bases[i].Function()
			sum += self.y[i] * li(x)
		}
		return sum
	}
}

func GenerateChebyshevRoots(degree int, start, end float64) []float64 {
	roots := make([]float64, degree)
	for k := 0; k < degree; k++ {
		roots[k] = math.Cos(math.Pi*float64(2 * k + 1)/ float64(2*degree))
		// adjust to interval
		roots[k] = (roots[k] * (end - start) + start + end)/2
	}
	return roots
}

func GenerateEquiDistanceRoots(degree int, start, end float64) []float64 {
	roots := make(Vector, degree)
	if (degree == 1) {
		roots[0] = (start + end)/2
		return roots
	}
	for k := 0; k < degree; k++ {
		roots[k] = start + (end - start)*float64(k)/float64(degree-1)
	}
	return roots
}