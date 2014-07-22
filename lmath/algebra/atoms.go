package algebra

import (
	"fmt"
	"math/big"
	"errors"
)

// Atom is used as Scalar or Variable
type Atom interface {
	Evaluate(vars map[string] *big.Rat) (*big.Rat, error)
}


type Variable struct {
	Identifier string
}

func (self Variable) Evaluate(vars map[string] *big.Rat) (*big.Rat, error) {
	if val, exists := vars[self.Identifier]; exists {
		return val, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Unknown variable: %v", self.Identifier))
	}
}

type Scalar64 struct {
	Value float64
}

type ScalarRat struct {
	Value *big.Rat
}

func (self ScalarRat) Evaluate(vars map[string] *big.Rat) (*big.Rat, error) {
	return self.Value, nil
}