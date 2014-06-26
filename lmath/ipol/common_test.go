package ipol

import (
	. "code.google.com/p/liblundis/lmath"
	"testing"
	"fmt"
)

func TestL_k(t *testing.T) {
	degree := 2
	x := GenerateEquiDistanceRoots(degree, -1, 1)
	for k := range x {
		lk := L_k(k, x)
		for i, xi := range x {
			if i == k {
				AssertEqualsRat(t, lk.ValueAt(NewRatf(xi)), NewRati(1), fmt.Sprintf("l_%v(%v)", k, xi))
			} else {
				AssertEqualsRat(t, lk.ValueAt(NewRatf(xi)), NewRati(0), fmt.Sprintf("l_%v(%v)", k, xi))
			}
		}
		
	}
}