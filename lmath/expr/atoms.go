package expr

import (
	. "code.google.com/p/liblundis/lmath"
	"fmt"
	"math/big"
	"errors"
)

// An Atom represents c * x^k, where c is a real number and k is an integer
type Atom struct {
	val *big.Rat
	var_id string
	pow int
}

func NewAtomVal(v *big.Rat) *Atom {
	return NewAtom(v, "", 1)
}

func NewAtomVar(v string) *Atom {
	return NewAtom(NewRati(1), v, 1)
}

func NewAtom(value *big.Rat, v string, pow int) *Atom {
	a := new(Atom)
	a.val = value
	a.var_id= v
	a.pow = pow
	return a
}

func (self *Atom) String() string {
	str := self.val.String()
	if self.val.IsInt() {
		// remove redundant "/1"
		str = str[:len(str)-2]
	}

	if self.var_id != "" {
		str = fmt.Sprintf("%v * %v^%v", str, self.var_id, self.pow)
	}
	return str
}

func (self *Atom) Evaluate(vars map[string] *big.Rat) (*big.Rat, error) {
	if self.var_id == "" {
		return self.val, nil
	} else {
		if val, exists := vars[self.var_id]; exists {
			tmp := NewRat(val)
			BigRatPow(tmp, self.pow)
			return tmp.Mul(tmp, self.val), nil
		} else {
			return nil, errors.New(fmt.Sprintf("Unknown variable: %v", self.var_id))
		}
	}
	
}

func (self *Atom) Replace(vars map[string] *big.Rat) {
	val, err := self.Evaluate(vars)
	if err == nil { // if the variable replacement was successful
		self.val = val
		self.var_id = ""
		self.pow = 1
	}
}