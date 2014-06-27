package ipol

import(
	. "code.google.com/p/liblundis/lmath"
	"code.google.com/p/liblundis"
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
	// ||(Bnx^2)(x) - x^2||inf = 1/(4n^2) + 1/n
	for n := 1; n < 20; n++ {
		bernstein := NewBernsteinInterpolation(n, f, -1, 1)
		error := FindMaxDiff(bernstein.Function(), f, -1, 1)
		experr := float64(1+4*n)/float64(4*n*n)
		if error > experr && !liblundis.Equals(error, experr) {
			fmt.Sprintf("degree %v error: %v > %v", n, error, experr)
		} 
	}
}