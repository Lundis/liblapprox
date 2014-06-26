package lmath

import (
	"testing"
)



func TestDegree(t *testing.T) {
	for i := 0; i < 10; i++ {
		p := NewPolynomial(i)
		AssertEqualsInt(t, p.Degree(), i, "")
	}
	
}


func TestAdd(t *testing.T) {
	p1 := Polynomial{1, 2} // 2x + 1
	p2 := Polynomial{3, 4} // 4x + 3
	p := p1.Plus(p2) // 6x + 4
	
	AssertEqualsPolynomial(t, p, Polynomial{4, 6})
}

func TestMult(t *testing.T) {
	p1 := Polynomial{1, 2} // 2x + 1
	p2 := Polynomial{3, 4} // 4x + 3
	p := p1.Mult(p2) // (2x + 1)(4x + 3) = 8x^2 + 10x + 3
	AssertEqualsPolynomial(t, p, Polynomial{3, 10, 8})

}

func TestPow(t *testing.T) {
	p := Polynomial{0, 1}
	p = p.Pow(10)
	AssertEqualsInt(t, p.Degree(), 10, "degree(x^10) == 10")
	p2 := Polynomial{1, 1}.Pow(2)
	AssertEqualsPolynomial(t, p2, Polynomial{1, 2, 1})
	AssertEqualsPolynomial(t, p2.Pow(0), Polynomial{1})
}

func TestDerive(t *testing.T) {
	
}

func TestFunction(t *testing.T) {

}