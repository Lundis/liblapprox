package lmath

import (
	"testing"
	"math/big"
)

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
		if p[i].Cmp(correct[i]) != 0 {
			t.Errorf("p[%d] == %v, should be %v\n", i, p[i], correct[i])
		}
	}
	
	for ; i <= max.Degree(); i++ {
		if max[i].Cmp(big.NewRat(0, 1)) != 0 {
			t.Errorf("degreediff, p[%v] == %v, not 0", i, max[i])
		}
	}
}

func TestBigPolyDegree(t *testing.T) {
	for i := 0; i < 10; i++ {
		p := NewBigPoly(i)
		AssertEqualsInt(t, p.Degree(), i, "")
	}
	
}

func TestBigPolyAdd(t *testing.T) {
	p1 := NewBigPoly1f(1, 2) // 2x + 1
	p2 := NewBigPoly1f(3, 4) // 4x + 3
	p := p1.Plus(p2) // 6x + 4
	
	AssertEqualsBigPoly(t, p, NewBigPoly1f(4, 6))
}

func TestBigPolyMult(t *testing.T) {
	p1 := NewBigPoly1f(1, 2) // 2x + 1
	p2 := NewBigPoly1f(3, 4) // 4x + 3
	p := p1.Mult(p2) // (2x + 1)(4x + 3) = 8x^2 + 10x + 3
	AssertEqualsBigPoly(t, p, NewBigPoly2f(3, 10, 8))

}

func TestBigPolyPow(t *testing.T) {
	p := NewBigPoly1f(0, 1)
	p = p.Pow(10)
	AssertEqualsInt(t, p.Degree(), 10, "degree(x^10) == 10")
	p2 := NewBigPoly1f(1, 1).Pow(2)
	AssertEqualsBigPoly(t, p2, NewBigPoly2f(1, 2, 1))
	AssertEqualsBigPoly(t, p2.Pow(0), NewBigPoly0f(1))
}

func TestBigPolyDerive(t *testing.T) {
	
}

func TestBigPolyFunction(t *testing.T) {

}