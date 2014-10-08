package poly

import (
	"bytes"
	"fmt"
	"math/big"
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/util/cont"
)

type BigPoly []*big.Rat

func NewBigPoly(degree int) BigPoly {
	poly := make([]*big.Rat, degree + 1)
	for i := range poly {
		poly[i] = NewRati(0)
	}
	return poly
}

func NewBigPoly0(x0 *big.Rat) BigPoly {
	poly := NewBigPoly(0)
	poly[0].Set(x0)
	return poly
}

func NewBigPoly0f(x0 float64) BigPoly {
	return NewBigPoly0(NewRatf(x0))
}

func NewBigPoly1(x0, x1 *big.Rat) BigPoly {
	poly := NewBigPoly(1)
	poly[0].Set(x0)
	poly[1].Set(x1)
	return poly
}

func NewBigPoly1f(x0, x1 float64) BigPoly {
	return NewBigPoly1(NewRatf(x0), NewRatf(x1))
}

func NewBigPoly2(x0, x1, x2 *big.Rat) BigPoly {
	poly := NewBigPoly(2)
	poly[0].Set(x0)
	poly[1].Set(x1)
	poly[2].Set(x2)
	return poly
}

func NewBigPoly2f(x0, x1, x2 float64) BigPoly {
	return NewBigPoly2(NewRatf(x0), NewRatf(x1), NewRatf(x2))
}

func NewBigPoly3(x0, x1, x2, x3 *big.Rat) BigPoly {
	poly := NewBigPoly(3)
	poly[0].Set(x0)
	poly[1].Set(x1)
	poly[2].Set(x2)
	poly[3].Set(x3)
	return poly
}

func NewBigPoly3f(x0, x1, x2, x3 float64) BigPoly {
	return NewBigPoly3(NewRatf(x0), NewRatf(x1), NewRatf(x2), NewRatf(x3))
}

func (self BigPoly) Copy() BigPoly {
	poly := NewBigPoly(self.Degree())
	for i, v := range self {
		poly[i].Set(v)
	}
	return poly
}

func (self BigPoly) Degree() int {
	return len(self) - 1
}

func (self BigPoly) Plus(other BigPoly) BigPoly {
	
	var higher, lower BigPoly
	if self.Degree() >= other.Degree() {
		higher = self
		lower = other
	} else {
		higher = other
		lower = self
	}

	poly := higher.Copy()
	
	for i := 0; i <= lower.Degree(); i++ {
		poly[i].Add(poly[i], lower[i])
	}
	return poly
}

func (self BigPoly) MultRat(k *big.Rat) BigPoly {
	poly := self.Copy()
	for i := range poly {
		poly[i].Mul(poly[i], k)
	}
	return poly
}

func (self BigPoly) MultFloat64(k float64) BigPoly {
	return self.MultRat(NewRatf(k))
}

func (self BigPoly) MultInt64(k int64) BigPoly {
	return self.MultRat(NewRati64(k))
}

func (self BigPoly) Minus(other BigPoly) BigPoly {
	return self.Plus(other.MultFloat64(-1))
}

func (self BigPoly) Mult(other BigPoly) BigPoly {
	degree := self.Degree() + other.Degree()
	poly := NewBigPoly(degree)
	for grade1, val1 := range self {
		for grade2, val2 := range other {
			tmp := big.NewRat(1, 1).Mul(val1, val2)
			poly[grade1 + grade2].Add(poly[grade1 + grade2], tmp)
		}
	}
	return poly
}

func (self BigPoly) Pow(k int) BigPoly {
	if k == 0 {
		return NewBigPoly0(NewRati(1))
	}
	poly := self.Copy()
	for i := 1; i < k; i++ {
		poly = poly.Mult(self)
	}
	return poly
}

func (self BigPoly) ValueAt(x *big.Rat) *big.Rat {
	x_val := NewRati(1)
	tmp := NewRati(1)
	sum := NewRati(1).Set(self[0])
	for i := 1; i < len(self); i++ {
		x_val.Mul(x_val, x)
		tmp.Mul(x_val, self[i])
		sum.Add(sum, tmp)
	}
	return sum
}

func (self BigPoly) ValueAtf64(x float64) float64 {
	x_val := float64(1)
	sum, _ := self[0].Float64()
	for i := 1; i < len(self); i++ {
		x_val *= x
		k, _ := self[i].Float64()
		sum += x_val * k
	}
	return sum
}

func (self BigPoly) Derive() BigPoly {
	if self.Degree() == 0 {
		return NewBigPoly0(big.NewRat(0, 1))
	}
	d := NewBigPoly(self.Degree() - 1)
	for i := 0; i <= d.Degree(); i++ {
		d[i].Mul(self[i+1], NewRati(i + 1))
	}
	return d
}

func (self BigPoly) Function() Function {
	return func(x float64) float64 {
		return self.ValueAtf64(x)
	}
}

func (self BigPoly) String() string {
	var buffer bytes.Buffer
	first := true
	for i := len(self) - 1; i >= 0; i-- {
		if self[i].Cmp(big.NewRat(0, 1)) != 0 {
			if !first {
				buffer.WriteString(" + ")
			}
			fmt.Fprintf(&buffer, "(%v)", self[i])
			if i != 0 {
				fmt.Fprintf(&buffer, "*x^%d", i)
			}
			first = false
		}
	}
	return buffer.String()
}