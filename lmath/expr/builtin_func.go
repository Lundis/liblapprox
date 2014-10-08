package expr

import (
	. "code.google.com/p/liblundis/lmath"
	. "code.google.com/p/liblundis/lmath/util/cont"
	"fmt"
	"math"
	"math/big"
	"errors"
	"strings"
)

type builtin_func struct {
	inside []*Node
	id string
}

var single_arg_funcs []string = []string{"cos", "sin", "sqrt", "abs"}
// future additions: max, min
var double_arg_funcs []string = []string{}

func NewBuiltinFunc(id string, inside []*Node) (*builtin_func, error) {
	if len(inside) == 1 {
		for _, s := range single_arg_funcs {
			if id == s {
				f := new(builtin_func)
				f.id = id
				f.inside = inside
				return f, nil
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("Unknown function / faulty parameters: %v", id))
}

func (self *builtin_func) String() string {
	if len(self.inside) == 0 {
		return fmt.Sprintf("%v()", self.id)
	} else {
		inside_strs := make([]string, len(self.inside))
		for i, n := range self.inside {
			inside_strs[i] = n.String()
		}
		args := strings.Join(inside_strs, ", ")
		return fmt.Sprintf("%v(%v)", self.id, args)
	}
	
}

func (self *builtin_func) Evaluate(vars map[string] *big.Rat) (*big.Rat, error) {
	r := NewRati(0)
	inside_rats := make([]*big.Rat, len(self.inside))
	for i, n := range self.inside {
		inside_rat, err := n.Evaluate(vars)
		if err != nil {
			return nil, err
		} else {
			inside_rats[i] = inside_rat
		}
	}
	if len(self.inside) == 1 {
		arg, _ := inside_rats[0].Float64()
		switch self.id {
		case "cos":
			r.SetFloat64(math.Cos(arg))
		case "sin":
			r.SetFloat64(math.Sin(arg))
		case "abs":
			r.SetFloat64(math.Abs(arg))
		case "sqrt":
			r.SetFloat64(math.Sqrt(arg))
		}
	}
	return r, nil
}

func (self *builtin_func) Replace(vars map[string] *big.Rat) {
	for _, n := range self.inside {
		n.Replace(vars)
	}
}

func (self *builtin_func) Function() Function {
	inside_funcs := make([]Function, len(self.inside))
	for i, n := range self.inside {
		inside_funcs[i] = n.Function()
	}
	if len(self.inside) == 1 {
		switch self.id {
		case "cos":
			return func(x float64) float64 {
				return math.Cos(inside_funcs[0](x))
			}
		case "sin":
			return func(x float64) float64 {
				return math.Sin(inside_funcs[0](x))
			}
		case "abs":
			return func(x float64) float64 {
				return math.Abs(inside_funcs[0](x))
			}
		case "sqrt":
			return func(x float64) float64 {
				return math.Sqrt(inside_funcs[0](x))
			}
		default:
			panic(fmt.Sprintf("unknown 1-parameter function: %v", self.id))
		}
	} else {
		panic("not implemented")
	}
}

func (self *builtin_func) Optimize() {
	for _, child := range self.inside {
		child.Optimize()
	}
}