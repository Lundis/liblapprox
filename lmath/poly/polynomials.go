package poly

import (
	"bytes"
	"fmt"
	"code.google.com/p/liblundis/lmath"
)

type Poly []float64

func NewPoly(degree int) Poly {
	return make([]float64, degree + 1)
}

func (self Poly) Copy() Poly {
	poly := NewPoly(self.Degree())
	copy(poly, self)
	return poly
}

func (self Poly) Degree() int {
	return len(self) - 1
}

func (self Poly) Plus(other Poly) Poly {
	
	var higher, lower Poly
	if self.Degree() >= other.Degree() {
		higher = self
		lower = other
	} else {
		higher = other
		lower = self
	}

	poly := higher.Copy()
	
	for i := 0; i <= lower.Degree(); i++ {
		poly[i] += lower[i]
	}
	return poly
}

func (self Poly) MultConstant(k float64) Poly {
	poly := self.Copy()
	for i := range poly {
		poly[i] *= k
	}
	return poly
}

func (self Poly) Minus(other Poly) Poly {
	return self.Plus(other.MultConstant(-1))
}

func (self Poly) Mult(other Poly) Poly {
	degree := self.Degree() + other.Degree()
	poly := NewPoly(degree)
	for grade1, val1 := range self {
		for grade2, val2 := range other {
			poly[grade1 + grade2] += val1*val2
		}
	}
	return poly
}

// TODO: this could be made more numerically stable for large values of k
func (self Poly) Pow(k int) Poly {
	if k == 0 {
		return Poly{1}
	}
	poly := self.Copy()
	for i := 1; i < k; i++ {
		poly = poly.Mult(self)
	}
	return poly
}

func (self Poly) ValueAt(x float64) float64 {
	x_val := float64(1)
	sum := self[0]
	for i := 1; i < len(self); i++ {
		x_val *= x
		sum += x_val*self[i]
	}
	return sum
}

func (self Poly) Derive() Poly {
	if self.Degree() == 0 {
		return Poly{0}
	}
	d := NewPoly(self.Degree() - 1)
	for i := 0; i <= d.Degree(); i++ {
		d[i] = float64(i+1)*self[i+1]
	}
	return d
}

func (self Poly) Function() lmath.Function {
	return func(x float64) float64 {
		return self.ValueAt(x)
	}
}

func (self Poly) String() string {
	var buffer bytes.Buffer
	first := true
	for i := len(self) - 1; i >= 0; i-- {
		if !lmath.EqualsFloat(self[i], 0, 1e-6) {
			if !first {
				buffer.WriteString(" + ")
			}
			fmt.Fprintf(&buffer, "%.6f", self[i])
			if i != 0 {
				fmt.Fprintf(&buffer, "*x^%d", i)
			}
			first = false
		}
	}
	return buffer.String()
}