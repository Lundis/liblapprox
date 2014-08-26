package ipol

import (
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/base/poly"
)

func L_k(k int, x Vector) BigPoly {
	divisor := NewBigPoly0f(1)
	poly := NewBigPoly0f(1)
	
	for i, x_i := range x {
		if i != k {
			poly = poly.Mult(NewBigPoly1f(-x_i, 1)) // (x - x_i)
			divisor = divisor.MultRat(NewRatf(x[k] - x_i)) // (x_k - x_i)
		}
	}
	tmp := NewRati(1)
	together := poly.MultRat(tmp.Inv(divisor[0]))
	return together
}
