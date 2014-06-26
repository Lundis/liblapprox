package ipol

import (
	. "code.google.com/p/liblundis/lmath"
	"testing"
	"math"
	"fmt"
)

func TestCreateFejerHermitePolynomial(t *testing.T) {
	f := func(x float64) float64 {
		return 1 - math.Abs(x)
	}
	roots := GenerateChebyshevRoots(5, -1, 1)
	p := CreateFejerHermitePolynomialvf(roots, f, -1, 1)
	for i, root := range roots {
		AssertEqualsFloat64(t, p.ValueAt(root), 0, fmt.Sprintf("faulty root #%v", i))
	}
}