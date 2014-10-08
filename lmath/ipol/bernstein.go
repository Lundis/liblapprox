package ipol

import(
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/util/cont"
	"code.google.com/p/liblundis/lmath/base/poly"
	"math/big"
)

type BernsteinInterpolation struct {
	start, end float64
	poly poly.BigPoly
}


func NewBernsteinInterpolation(degree int, f Function, start, end float64) BernsteinInterpolation {
	sum := poly.NewBigPoly0f(0)
	for k := 0; k <= degree; k++ {
		p1 := poly.NewBigPoly1f(0, 1).Pow(k)
		p2 := poly.NewBigPoly1f(1, -1).Pow(degree-k)
		x_k := start + (end - start) * float64(k)/float64(degree)
		bin := big.NewRat(1, 1).SetInt(big.NewInt(1).Binomial(int64(degree), int64(k)))
		c := bin.Mul(bin, NewRatf(f(x_k)))
		sum = sum.Plus(p1.Mult(p2).MultRat(c))
	}
	return BernsteinInterpolation{start, end, sum}
}

func (self BernsteinInterpolation) Function() Function {
	f := self.poly.Function()
	return func(x float64) float64 {
		return f((x - self.start) / (self.end - self.start))
	}
}

func (self BernsteinInterpolation) BigPoly() poly.BigPoly {
	return self.poly
}

