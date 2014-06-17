package ltest

import (
	"testing"
	"liblundis"
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