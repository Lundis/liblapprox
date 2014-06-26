package ipol

import(
	. "code.google.com/p/liblundis/lmath"
)

type BernsteinInterpolation struct {
	start, end float64
	poly Polynomial
}


func NewBernsteinInterpolation(degree int, f Func1to1, start, end float64) BernsteinInterpolation {
	sum := Polynomial{0}
	for k := 0; k <= degree; k++ {
		p1 := Polynomial{0, 1}.Pow(k)
		p2 := Polynomial{1, -1}.Pow(degree-k)
		x_k := start + (end - start) * float64(k)/float64(degree)
		c := float64(BinCoeff(degree, k)) * f(x_k)
		sum = sum.Plus(p1.Mult(p2).MultConstant(c))
	}
	return BernsteinInterpolation{start, end, sum}
}

func (self BernsteinInterpolation) Function() Func1to1 {
	f := self.poly.Function()
	return func(x float64) float64 {
		return f((x - self.start) / (self.end - self.start))
	}
}

func (self BernsteinInterpolation) Polynomial() Polynomial {
	return self.poly
}

