package ipol

import(
	. "code.google.com/p/liblundis/lmath"
	"testing"
	"fmt"
)

func TestBernsteinPolynomial(t *testing.T) {
	f := func(x float64) float64 {
		return x*x
	}
	// for this function the bernstein polynomial is
	// Bnf = x/n + (n-1)/n * x^2
	for n := 1; n < 10; n++ {
		bnf := NewBernsteinInterpolation(n, f, 0, 1).BigPoly()
		expected := NewBigPoly2f(0, 1.0/float64(n), float64(n-1)/float64(n))
		AssertEqualsBigPoly(t, bnf, expected)
	}
}

func TestBernsteinPolynomialInterval(t *testing.T) {
	f := func(x float64) float64 {
		return x*x
	}
	// for this function the bernstein polynomial in the range -1 <= x <= 1 is
	// Bnf = 1/(4n^2) + (n-1)/n * x^2
	for n := 1; n < 10; n++ {
		bnf := NewBernsteinInterpolation(n, f, -1, 1).BigPoly()
		//fmt.Printf("%v\n", bnf)
		expected := NewBigPoly2f(1.0/float64(4*n*n), 0, float64(n-1)/float64(n))
		g := func(x float64) float64 {
			f := expected.Function()
			return f((x + 1) / 2)
		}
		error := FindMaxDiff(bnf.Function(), g, -1, 1)
		AssertEqualsFloat64(t, error, 0, fmt.Sprintf("degree %v error: ", n))
	}
}