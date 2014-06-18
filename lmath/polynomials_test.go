package lmath

import (
	"testing"
	"code.google.com/p/liblundis/ltest"
	"code.google.com/p/liblundis"
)

func testDegree(t *testing.T) {
	for i := 0; i < 10; i++ {
		p := NewPolynomial(i)
		ltest.AssertEqualsInt(t, p.Degree(), i, "")
	}
	
}

func assertEquals(t *testing.T, p, correct Polynomial) {
	if p.Degree() != correct.Degree() {
		t.Errorf("Different degrees: p1.Degree() == %v, correct.Degree() == %v\n", p.Degree(), correct.Degree())
		t.FailNow()
	}
	for i := 0; i <= p.Degree(); i++ {
		ltest.AssertEqualsFloat64(t, p[i], correct[i], "degree " + string(i))
		if !liblundis.Equals(p[i], correct[i]) {
			t.Errorf("p[%d] == %v, should be %v\n", i, p[0], correct[i])
		}
	}
}

func TestAdd(t *testing.T) {
	p1 := Polynomial{1, 2} // 2x + 1
	p2 := Polynomial{3, 4} // 4x + 3
	p := p1.Plus(p2) // 6x + 4
	
	assertEquals(t, p, Polynomial{4, 6})
}

func TestMult(t *testing.T) {
	p1 := Polynomial{1, 2} // 2x + 1
	p2 := Polynomial{3, 4} // 4x + 3
	p := p1.Mult(p2) // (2x + 1)(4x + 3) = 8x^2 + 10x + 3
	assertEquals(t, p, Polynomial{3, 10, 8})

}

func TestFunction(t *testing.T) {

}