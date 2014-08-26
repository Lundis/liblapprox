package ipol

import(
	"fmt"
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/base/poly"
)

type LagrangeInterpolation struct {
	x Vector
	y Vector
	bases []*Poly
}

func NewLagrangeInterpolationvv(x, y Vector) LagrangeInterpolation {
	if len(x) != len(y) {
		fmt.Printf("NewLagrangeInterpolation error: len(x) != len(y)")
	} else if len(x) < 1 {
		fmt.Printf("NewLagrangeInterpolation error: can't be of degree zero")
	}
	lagrange := LagrangeInterpolation{make([]float64, len(x)), make([]float64, len(x)), make([]*Poly, len(x))}
	copy(lagrange.x, x)
	copy(lagrange.y, y)
	lagrange.generateBases()
	return lagrange
}

func NewLagrangeInterpolationfv(f Function, x Vector) LagrangeInterpolation {
	y := Values(f, x)
	return NewLagrangeInterpolationvv(x, y)
}

func (self *LagrangeInterpolation) generateBases() {
	for i := 0; i < len(self.bases); i++ {
		poly := NewPoly0(1)
		for j := 0; j < len(self.x); j++ {
			if i != j {
				poly = poly.Mult(NewPoly1(-self.x[j], 1))
			}
		}
		self.bases[i] = poly.MultConstant(1.0 / poly.ValueAt(self.x[i]))
	}
}

func (self LagrangeInterpolation) Bases() []*Poly {
	return self.bases
}

func (self LagrangeInterpolation) Function() Function {
	return func(x float64) float64 {
		sum := float64(0)
		for i := 0; i < len(self.bases); i++ {
			li := self.bases[i].Function()
			sum += self.y[i] * li(x)
		}
		return sum
	}
}

