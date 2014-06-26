package lmath

import (
	"testing"
	"code.google.com/p/liblundis"
	"fmt"
	"math/big"
)

func AssertEqualsFloat64(t *testing.T, x, expected float64, message string) {
	if !liblundis.Equals(x, expected) {
		t.Errorf("%v: %v != %v", message, x, expected)
	}
}

func AssertEqualsInt(t *testing.T, x, expected int, message string) {
	if x != expected {
		t.Errorf("%v: %v != %v", message, x, expected)
	}
}

func AssertEqualsRat(t *testing.T, x, expected *big.Rat, message string) {
	diff := big.NewRat(1, 1).Sub(x, expected)
	diff = diff.Abs(diff)
	error_margin := big.NewRat(1, 1e10)
	if diff.Cmp(error_margin) > 0 {
		error, _ := diff.Float64()
		t.Errorf("%v: %v != %v, error: %v", message, x, expected, error)
	}
}

func AssertEqualsPolynomial(t *testing.T, p, correct Polynomial) {
	var min, max Polynomial
	if p.Degree() < correct.Degree() {
		max = correct
		min = p
	} else {
		max = p
		min = correct
	}

	i := 0
	for ; i <= min.Degree(); i++ {
		AssertEqualsFloat64(t, p[i], correct[i], fmt.Sprintf("p[%d] == %v", i, p[i]))
	}
	
	for ; i <= max.Degree(); i++ {
		AssertEqualsFloat64(t, max[i], 0, fmt.Sprintf("degreediff, p[%v] != 0", i))
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
		AssertEqualsRat(t, p[i], correct[i], fmt.Sprintf("p[%d] == %v", i, p[i]))
	}
	
	for ; i <= max.Degree(); i++ {
		AssertEqualsRat(t, max[i], big.NewRat(0, 1), fmt.Sprintf("degreediff, p[%v] != 0", i))
	}
}