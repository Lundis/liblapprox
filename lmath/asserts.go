package lmath

import (
	"testing"
	"math/big"

)

func AssertEqualsFloat64(t *testing.T, x, expected float64, message string) {
	if !EqualsFloat(x, expected, 1e-6) {
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