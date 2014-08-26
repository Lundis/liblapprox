package exchange

import (
	"testing"
	"fmt"
	"code.google.com/p/liblundis/lmath/approx"
	"code.google.com/p/liblundis/lmath/base/poly"
)

func TestApproximate(t *testing.T) {
	f := func(x float64) float64 {
		return x*x*x
	}
	approx := approx.NewApprox(f, 0, 1)
	iters := Approximate(approx, []int{0, 1, 2}, 1e-10, poly.PolyFromBasisImpl)
	for i, iter := range iters[0] {
		fmt.Printf("%v: %v\n", i, iter)
	}
	for i, iter := range iters[1] {
		fmt.Printf("%v: %v\n", i, iter)
	}
	for i, iter := range iters[2] {
		fmt.Printf("%v: %v\n", i, iter)
	}
}

func TestUpdateRoots(t *testing.T) {
	
}