package algebra

import (
	"math/big"
)

// Atom is used as Scalar or Variable
type Atom interface {}

// Scalar is either Scalar64 or ScalarRat
type Scalar interface {}

type Variable struct {
	Identifier string
}

type Scalar64 struct {
	Value float64
}

type ScalarRat struct {
	Value *big.Rat
}