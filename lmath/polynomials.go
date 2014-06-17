package lmath

import (
	"bytes"
	"fmt"
	"liblundis"
)

type Polynomial []float64

func NewPolynomial(degree int) Polynomial {
	return make([]float64, degree + 1)
}

func (self Polynomial) Copy() Polynomial {
	poly := NewPolynomial(self.Degree())
	copy(poly, self)
	return poly
}

func (self Polynomial) Degree() int {
	return len(self) - 1
}

func (self Polynomial) Plus(other Polynomial) Polynomial {
	
	var higher, lower Polynomial
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

func (self Polynomial) MultConstant(k float64) Polynomial {
	poly := self.Copy()
	for i := range poly {
		poly[i] *= k
	}
	return poly
}

func (self Polynomial) Minus(other Polynomial) Polynomial {
	return self.Plus(other.MultConstant(-1))
}

func (self Polynomial) Mult(other Polynomial) Polynomial {
	degree := self.Degree() + other.Degree()
	poly := NewPolynomial(degree)
	for grade1, val1 := range self {
		for grade2, val2 := range other {
			poly[grade1 + grade2] += val1*val2
		}
	}
	return poly
}

func (self Polynomial) ValueAt(x float64) float64 {
	x_val := float64(1)
	sum := self[0]
	for i := 1; i < len(self); i++ {
		x_val *= x
		sum += x_val*self[i]
	}
	return sum
}

func (self Polynomial) Function() Func1to1 {
	return func(x float64) float64 {
		return self.ValueAt(x)
	}
}

func (self Polynomial) String() string {
	var buffer bytes.Buffer
	first := true
	for i := len(self) - 1; i >= 0; i-- {
		if !liblundis.Equals(self[i], 0) {
			if !first {
				buffer.WriteString(" + ")
			}
			fmt.Fprintf(&buffer, "%.4f", self[i])
			if i != 0 {
				fmt.Fprintf(&buffer, "*x^%d", i)
			}
			first = false
		}
	}
	return buffer.String()
}