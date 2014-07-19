package algebra

import (
	"code.google.com/p/liblundis/lmath"
)

type Expression Node

// Simplify the expression. The expression is flattened and all addable terms are added together.
func (self *Expression) Simplify() {

}

// Creates a Polynomial representing this Expression. If the expression cant be reduced to a polynomial it returns nil.
// The Expression can only be reduced to a polynomial if it contains just plus/minus/multiplication and has one or no variables.
func (self *Expression) AsPolynomial() *lmath.Polynomial {
	
}