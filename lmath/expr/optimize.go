package expr

import (
	"fmt"
	"math"
)

func (self *Node) OptimizeCommon() {
	self.Replace(common_constants)
	self.Optimize()
}

func (self *Node) Optimize() {
	for _, child := range self.nodes {
		child.Optimize()
	}
	switch self.op {
	case PLUS:
		// combine atoms
	case MINUS:
		// move down the chain
	case MULT:
		// combine atoms and move the multiplication down the chain if possible
	case DIV:

	case POW:
		// well.. uuhh yeah let's skip this one
	case ATOM:
		// nothing to do
	case FUNC:
		self.data.(*builtin_func).Optimize()
	default:
		panic(fmt.Sprintf("unknown op: %v", self.op))
	}
}

func (self *Node) Mult(other *Node) {

}

type Degree struct {
	vars map[string] int
	has_abs, has_pow bool
}

func (self *Degree) Equals(other *Degree) bool {
	// cant compare powers and absolute values
	if self.has_abs || other.has_abs {
		return false
	}
	if self.has_pow || other.has_pow {
		return false
	}
	for key, val1 := range self.vars {
		if val2, exists := other.vars[key]; exists {
			if val1 != val2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func (self *Degree) Mult(other *Degree) *Degree {
	deg := new(Degree)
	deg.has_pow = self.has_pow || other.has_pow
	deg.has_abs = self.has_abs || other.has_abs
	deg.vars = make(map[string]int)
	for key, val := range self.vars {
		deg.vars[key] = val
	}

	for key, val := range other.vars {
		if _, exists := deg.vars[key]; exists {
			deg.vars[key] += val
		} else {
			deg.vars[key] = val
		}
	}
	return deg
}

func (self *Degree) Add(other *Degree) *Degree {
	deg := new(Degree)
	deg.has_pow = self.has_pow || other.has_pow
	deg.has_abs = self.has_abs || other.has_abs
	deg.vars = make(map[string]int)
	for key, val := range self.vars {
		deg.vars[key] = val
	}

	for key, val := range other.vars {
		if _, exists := deg.vars[key]; exists {
			deg.vars[key] = int(math.Max(float64(deg.vars[key]), float64(val)))
		} else {
			deg.vars[key] = val
		}
	}
	return deg
}

func (self *Node) Degree() Degree {
	switch self.op {
	case PLUS:
	case MINUS:
	case MULT:
	case DIV:

	case POW:
		
	case FUNC:
		// this one is tricky!
	case ATOM:
		// nothing to do
	default:
		panic(fmt.Sprintf("unknown op: %v", self.op))
	}
	return Degree{}
}