package poly

import (
	"testing"
	"math/big"
	. "code.google.com/p/liblundis/lmath"
)

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

func TestBigPolyMinus(t *testing.T) {
	p1 := NewBigPoly1f(1, 2) // 2x + 1
	p2 := NewBigPoly1f(3, 4) // 4x + 3
	p := p1.Minus(p2) // 6x + 4
	
	AssertEqualsBigPoly(t, p, NewBigPoly1f(-2, -2))
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
	p := NewBigPoly0f(1)
	AssertEqualsBigPoly(t, p.Derive(), NewBigPoly0f(0))
	p = NewBigPoly2f(2, 3, 4)
	AssertEqualsBigPoly(t, p.Derive(), NewBigPoly1f(3, 8))
	p = NewBigPoly3f(2, 3, 4, 5)
	AssertEqualsBigPoly(t, p.Derive(), NewBigPoly2f(3, 8, 15))
}

func TestBigPolyValueAt(t *testing.T) {
	p := NewBigPoly1f(-1, 1).MultFloat64(-0.5)
	AssertEqualsRat(t, p.ValueAt(NewRati(-1)), NewRati(1), "p(-1)")
	AssertEqualsRat(t, p.ValueAt(NewRatf(-0.5)), big.NewRat(3, 4), "p(-1/2)")
	AssertEqualsRat(t, p.ValueAt(NewRati(0)), big.NewRat(1, 2), "p(0)")
	AssertEqualsRat(t, p.ValueAt(NewRatf(0.5)), big.NewRat(1, 4), "p(1/2)")
	AssertEqualsRat(t, p.ValueAt(NewRati(1)), NewRati(0), "p(1)")
}

func TestBigPolyValueAtf64(t *testing.T) {
	p := NewBigPoly1f(-1, 1).MultFloat64(-0.5)
	AssertEqualsFloat64(t, p.ValueAtf64(-1), 1, "p(-1)")
	AssertEqualsFloat64(t, p.ValueAtf64(0), 0.5, "p(0)")
	AssertEqualsFloat64(t, p.ValueAtf64(1), 0, "p(1)")
}

func TestBigPolyFunction(t *testing.T) {

}