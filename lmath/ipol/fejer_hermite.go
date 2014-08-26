package ipol

import (
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/base/poly"
)

func CreateFejerHermitePolynomialvv(x, y Vector, start, end float64) BigPoly {
	sum := NewBigPoly0f(0)
	for k := range x {
		l_k := L_k(k, x)
		Dl_k := l_k.Derive()
		l_k2 := l_k.Pow(2)
		Dl_k_x_k := Dl_k.ValueAt(NewRatf(x[k]))
		part := NewBigPoly1f(-x[k], 1).MultRat(Dl_k_x_k).MultFloat64(2)
		sum = sum.Plus(NewBigPoly0f(1).Minus(part).Mult(l_k2).MultFloat64(y[k]))
	}
	return sum
}

func CreateFejerHermitePolynomialvf(x Vector, f Function, start, end float64) BigPoly {
	y := Values(f, x)
	return CreateFejerHermitePolynomialvv(x, y, start, end)
}