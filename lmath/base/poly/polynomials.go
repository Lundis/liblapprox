package poly

import (
	"bytes"
	"fmt"
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/util/cont"
	. "code.google.com/p/liblundis/lmath/base"
)

type Poly struct {
	BasisImpl
}

func PolyFromBasisImpl(b *BasisImpl) Basis {
	p := Poly{*b}
	return &p
}

func NewPoly(degree int) *Poly {
	p := new(Poly)
	p.V = make([]float64, degree + 1)
	return p
}

func NewPoly0(a float64) *Poly {
	p := NewPoly(0)
	p.V[0] = a
	return p
}

func NewPoly1(a, b float64) *Poly {
	p := NewPoly(1)
	p.V[0] = a
	p.V[1] = b
	return p
}
func NewPoly2(a, b, c float64) *Poly {
	p := NewPoly(2)
	p.V[0] = a
	p.V[1] = b
	p.V[2] = c
	return p
}

func (self *Poly) Copy() *Poly {
	p := NewPoly(self.Degree())
	copy(p.V, self.V)
	return p
}

func (self *Poly) Degree() int {
	return len(self.V) - 1
}

func (self *Poly) Plus(other *Poly) *Poly {
	
	var higher, lower *Poly
	if self.Degree() >= other.Degree() {
		higher = self
		lower = other
	} else {
		higher = other
		lower = self
	}

	poly := higher.Copy()
	
	for i := 0; i <= lower.Degree(); i++ {
		poly.V[i] += lower.V[i]
	}
	return poly
}

func (self *Poly) MultConstant(k float64) *Poly {
	poly := self.Copy()
	for i := range poly.V {
		poly.V[i] *= k
	}
	return poly
}

func (self *Poly) Minus(other Poly) *Poly {
	return self.Plus(other.MultConstant(-1))
}

func (self *Poly) Mult(other *Poly) *Poly {
	degree := self.Degree() + other.Degree()
	poly := NewPoly(degree)
	for grade1, val1 := range self.V {
		for grade2, val2 := range other.V {
			poly.V[grade1 + grade2] += val1 * val2
		}
	}
	return poly
}

// TODO: this could be made more numerically stable for large values of k
func (self *Poly) Pow(k int) *Poly {
	if k == 0 {
		return NewPoly0(1)
	}
	poly := self.Copy()
	for i := 1; i < k; i++ {
		poly = poly.Mult(self)
	}
	return poly
}

func (self *Poly) ValueAt(x float64) float64 {
	x_val := float64(1)
	sum := self.V[0]
	for i := 1; i < len(self.V); i++ {
		x_val *= x
		sum += x_val * self.V[i]
	}
	return sum
}

func (self *Poly) Derive() *Poly {
	if self.Degree() == 0 {
		return NewPoly0(0)
	}
	d := NewPoly(self.Degree() - 1)
	for i := 0; i <= d.Degree(); i++ {
		d.V[i] = float64(i+1) * self.V[i+1]
	}
	return d
}

func (self *Poly) Function() Function {
	return func(x float64) float64 {
		return self.ValueAt(x)
	}
}

func (self *Poly) String() string {
	var buffer bytes.Buffer
	first := true
	for i := len(self.V) - 1; i >= 0; i-- {
		if !EqualsFloat(self.V[i], 0, 1e-6) {
			if !first {
				buffer.WriteString(" + ")
			}
			fmt.Fprintf(&buffer, "%.6f", self.V[i])
			if i != 0 {
				fmt.Fprintf(&buffer, "*x^%d", i)
			}
			first = false
		}
	}
	return buffer.String()
}