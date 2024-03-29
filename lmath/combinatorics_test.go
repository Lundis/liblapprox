package lmath

import (
	"testing"
)

func TestBinCoeff(t *testing.T) {
	AssertEqualsInt(t, int(BinCoeff(20, 10)), 184756, "BinCoeff(20, 10)")
	AssertEqualsInt(t, int(BinCoeff(15, 3)), 455, "BinCoeff(15, 3)")
	AssertEqualsInt(t, int(BinCoeff(20, 1)), 20, "BinCoeff(20, 1)")
	AssertEqualsInt(t, int(BinCoeff(20, 0)), 1, "BinCoeff(20, 0)")
}