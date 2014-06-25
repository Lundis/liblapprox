package ipol

import(
	. "code.google.com/p/liblundis/lmath"
)



func BernsteinPolynomial(degree int, f Func1to1, start, end float64) Polynomial {
	sum := Polynomial{0}
	for k := 0; k <= degree; k++ {
		p1 := Polynomial{0, 1}.Pow(k)
		p2 := Polynomial{1, -1}.Pow(degree-k)
		c := float64(BinCoeff(degree, k)) * f(float64(k)/float64(degree))
		sum = sum.Plus(p1.Mult(p2).MultConstant(c))
	}
	return sum
}

