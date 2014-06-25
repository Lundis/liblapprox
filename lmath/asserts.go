package lmath

import (
	"testing"
	"code.google.com/p/liblundis"
	"fmt"
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
		if !liblundis.Equals(p[i], correct[i]) {
			t.Errorf("p[%d] == %v, should be %v\n", i, p[0], correct[i])
		}
	}
	
	for ; i <= max.Degree(); i++ {
		AssertEqualsFloat64(t, max[i], 0, fmt.Sprintf("degreediff, p[%v] != 0", i))
	}


}