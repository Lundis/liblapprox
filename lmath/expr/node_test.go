package expr

import (
	"testing"
	"math/big"
)
func TestEvaluateSimple(t *testing.T) {
	assertEquals(t, "1+3", big.NewRat(4, 1), nil)
	assertEquals(t, "1-3", big.NewRat(-2, 1), nil)
	assertEquals(t, "2*3", big.NewRat(6, 1), nil)
	assertEquals(t, "2/4", big.NewRat(2, 4), nil)
}

func TestEvaluatePrecedence(t *testing.T) {
	assertEquals(t, "5 * (3 + 2)", big.NewRat(25, 1), nil)
	assertEquals(t, "5 * 3 + 5 * 6", big.NewRat(45, 1), nil)
	assertEquals(t, "5 * 3 * (3 + 2)", big.NewRat(75, 1), nil)
	assertEquals(t, "2 + 3 * 4", big.NewRat(14, 1), nil)
}

func TestEvaluateVariables(t *testing.T) {
	vars := map[string]*big.Rat{
		"a": big.NewRat(5, 1),
		"b": big.NewRat(23, 1),
	}
	assertEquals(t, "a+b", big.NewRat(28, 1), vars)
	assertEquals(t, "a-b", big.NewRat(-18, 1), vars)
	assertEquals(t, "a*b", big.NewRat(5*23, 1), vars)
	assertEquals(t, "a/b", big.NewRat(5, 23), vars)
}

func assertEquals(t *testing.T, expr string, correct *big.Rat, vars map[string] *big.Rat) {
	node, _ := ParseExpression(expr)
	result, _ := node.Evaluate(vars)
	if correct.Cmp(result) != 0 {
		t.Errorf("Evaluated %v. Got %v. Expected %v", expr, result, correct)
	}
}

func TestFunction(t *testing.T) {
	// test cases at x=0, x=0.5, x=1
	fun_str := []string{"1 + x/2", "5*x", "pi - pi"}
	expected := [][]float64{{1, 1.25, 1.5}, {0, 2.5, 5}, {0, 0, 0}}
	for i := range fun_str {
		node, err := ParseExpression(fun_str[i])
		if err != nil {
			t.Errorf("Failed to parse %v: ", fun_str[i], err)
		}
		exp := expected[i]
		f := node.Function()
		for x, v := range exp {
			f_x := f(float64(x)/2)
			if f_x != v {
				t.Errorf("%v, x=%v, expected %v, got %v", fun_str[i], x/2, v, f_x)
			}
		}
	}
}