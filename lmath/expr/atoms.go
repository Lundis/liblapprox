package expr

import (
	. "code.google.com/p/liblundis/lmath/util/cont"
	"math/big"
)

type Atom interface {
	String() string
	Evaluate(vars map[string] *big.Rat) (*big.Rat, error)
	Replace(vars map[string] *big.Rat)
	Function() Function
}