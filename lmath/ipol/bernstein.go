package ipol

import(
	. "code.google.com/p/liblundis/lmath"
	"math/big"
)

type BernsteinInterpolation struct {
	start, end float64
	poly BigPoly
}


func NewBernsteinInterpolation(degree int, f Func1to1, start, end float64) BernsteinInterpolation {
	sum := NewBigPoly0f(0)
	for k := 0; k <= degree; k++ {
		p1 := NewBigPoly1f(0, 1).Pow(k)
		p2 := NewBigPoly1f(1, -1).Pow(degree-k)
		x_k := start + (end - start) * float64(k)/float64(degree)
		bin := big.NewRat(1, 1).SetInt(big.NewInt(1).Binomial(int64(degree), int64(k)))
		c := bin.Mul(bin, big.NewRat(1,1).SetFloat64(f(x_k)))
		sum = sum.Plus(p1.Mult(p2).MultRat(c))
	}
	return BernsteinInterpolation{start, end, sum}
}

func (self BernsteinInterpolation) Function() Func1to1 {
	f := self.poly.Function()
	return func(x float64) float64 {
		return f((x - self.start) / (self.end - self.start))
	}
}

func (self BernsteinInterpolation) BigPoly() BigPoly {
	return self.poly
}
