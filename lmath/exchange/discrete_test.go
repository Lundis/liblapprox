package exchange

import (
	"testing"
	"fmt"
	"code.google.com/p/liblundis/lmath/approx"
	"code.google.com/p/liblundis/lmath/util/discrete"
	"code.google.com/p/liblundis/lmath/base/poly"
)

func TestApproximateDiscrete(t *testing.T) {
	f := func(x float64) float64 {
		return x*x*x
	}
	x := []float64{0, 0.5, 1, 1.5, 2, 2.5, 3}
	y := discrete.Values(f, x)
	fmt.Println("TestApproximateDiscrete")
	approx := approx.NewDiscreteApprox(x, y)
	iters := ApproximateDiscrete(approx, []int{0, 1, 2}, poly.PolyFromBasisImpl)
	for i, iter := range iters[0] {
		fmt.Printf("%v: %v\n", i, iter)
	}
	for i, iter := range iters[1] {
		fmt.Printf("%v: %v\n", i, iter)
	}
	for i, iter := range iters[2] {
		fmt.Printf("%v: %v\n", i, iter)
	}
	fmt.Println("TestApproximateDiscrete done \n")
}