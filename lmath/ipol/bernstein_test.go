package ipol

import(
	. "code.google.com/p/liblundis/lmath"
	"testing"
)

func TestBernsteinPolynomial(t *testing.T) {
	f := func(x float64) float64 {
		return x*x
	}
	// for this function the bernstein polynomial is
	// Bnf = x/n + (n-1)/n * x^2
	for n := 1; n < 10; n++ {
		bnf := BernsteinPolynomial(n, f, 0, 1)
		expected := Polynomial{0, 1.0/float64(n), float64(n-1)/float64(n) }
		AssertEqualsPolynomial(t, bnf, expected)
	}
}