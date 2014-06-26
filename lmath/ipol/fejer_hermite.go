package ipol

import (
	. "code.google.com/p/liblundis/lmath"
)

func l_k(k int, x Vector) Polynomial {
	divisor := 1.0
	poly := Polynomial{1}
	for i, x_i := range x {
		if i != k {
			poly = poly.Mult(Polynomial{-x_i, 1})
			divisor *= x[k] - x_i
		}
	}
	return poly.MultConstant(1.0/divisor)
}

func CreateFejerHermitePolynomialvv(x, y Vector, start, end float64) Polynomial {
	sum := Polynomial{0}
	for k := range x {
		l_k := l_k(k, x)
		Dl_k := l_k.Derive()
		l_k2 := l_k.Pow(2)
		part := Polynomial{-x[k], 1}.MultConstant(2.0 * Dl_k.ValueAt(x[k]))
		sum = sum.Plus(Polynomial{1}.Minus(part).Mult(l_k2).MultConstant(y[k]))
	}
	return sum
}

func CreateFejerHermitePolynomialvf(x Vector, f Func1to1, start, end float64) Polynomial {
	y := Values(f, x)
	return CreateFejerHermitePolynomialvv(x, y, start, end)
}