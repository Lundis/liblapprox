package expr

import (
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/util/cont"
	"fmt"
	"math"
	"math/big"
	"errors"
)

// An Literal represents c * x^k, where c is a real number and k is an integer
type Literal struct {
	val *big.Rat
	var_id string
	pow int
}

func NewLiteralVal(v *big.Rat) *Literal {
	return NewLiteral(v, "", 1)
}

func NewLiteralVar(v string) *Literal {
	return NewLiteral(NewRati(1), v, 1)
}

func NewLiteral(value *big.Rat, v string, pow int) *Literal {
	a := new(Literal)
	a.val = value
	a.var_id= v
	a.pow = pow
	return a
}

func (self *Literal) String() string {
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

func (self *Literal) Evaluate(vars map[string] *big.Rat) (*big.Rat, error) {
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

func (self *Literal) Replace(vars map[string] *big.Rat) {
	val, err := self.Evaluate(vars)
	if err == nil { // if the variable replacement was successful
		self.val = val
		self.var_id = ""
		self.pow = 1
	}
}

func (self *Literal) Function() Function {
	val, _ := self.val.Float64()
	if self.var_id == "" {
		return func(x float64) float64 {
			return val
		}
	} else {
		return func(x float64) float64 {
			return val * math.Pow(x, float64(self.pow))
		}
	}
}