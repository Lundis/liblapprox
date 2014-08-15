package algebra

import (
	"math"
	"math/big"
	"strings"
)

var common_constants map[string] *big.Rat = map[string] *big.Rat {
	"pi": newRatf64(math.Pi),
	"e": newRatf64(math.E),
}

func newRatf64(x float64) *big.Rat {
	rat := big.NewRat(1,1)
	rat.SetFloat64(x)
	return rat
}

// Evaluates an expression with common constants predefined.
// They include pi and e.
func EvalCommonf64(expr string) (float64, error) {
	rat, err := EvalCommon(expr)
	if err != nil {
		return 0, err
	} else {
		f64, _ := rat.Float64()
		return f64, nil
	}
}

func EvalCommon(expr string) (*big.Rat, error) {
	expr = strings.ToLower(expr)
	return Eval(expr, common_constants)
}

func Eval(expr string, vars map[string] *big.Rat) (*big.Rat, error) {
	node, err := ParseExpression(expr)
	if err != nil {
		return nil, err
	} else {
		return node.Evaluate(vars)
	}
}