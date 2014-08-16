package poly

import(
	"testing"
	"fmt"
	"math/big"
	"code.google.com/p/liblundis/lmath"
)

func AssertEqualsPoly(t *testing.T, p, correct Poly) {
	var min, max Poly
	if p.Degree() < correct.Degree() {
		max = correct
		min = p
	} else {
		max = p
		min = correct
	}

	i := 0
	for ; i <= min.Degree(); i++ {
		lmath.AssertEqualsFloat64(t, p[i], correct[i], fmt.Sprintf("p[%d] == %v", i, p[i]))
	}
	
	for ; i <= max.Degree(); i++ {
		lmath.AssertEqualsFloat64(t, max[i], 0, fmt.Sprintf("degreediff, p[%v] != 0", i))
	}
}

func AssertEqualsBigPoly(t *testing.T, p, correct BigPoly) {
	var min, max BigPoly
	if p.Degree() < correct.Degree() {
		max = correct
		min = p
	} else {
		max = p
		min = correct
	}

	i := 0
	for ; i <= min.Degree(); i++ {
		lmath.AssertEqualsRat(t, p[i], correct[i], fmt.Sprintf("p[%d] == %v", i, p[i]))
	}
	
	for ; i <= max.Degree(); i++ {
		lmath.AssertEqualsRat(t, max[i], big.NewRat(0, 1), fmt.Sprintf("degreediff, p[%v] != 0", i))
	}
}