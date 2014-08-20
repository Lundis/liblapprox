package expr

import(
	"testing"
)

// This test just makes sure that the expressions are parsed without throwing any errors.
func TestParse(t *testing.T) {
	exprs := []string{"1 + 3", "1 * 3", "3 * (1 + 3)", "1 / 5", "1 + 3 * (1 + 3) / 5", "1 000 000", "-1", "abs(-5)", "x^2"}	
	for _, v := range exprs {
		expr, err := ParseExpression(v)
		if err != nil {
			t.Errorf("parsed %v; Got %v; %v", v, expr, err)
		}
	}
}

// Checks that the parser throws error for faulty expressions
func TestParseFaulty(t *testing.T) {
	exprs := []string{"1 +", "* 3", "() / 5", "(1 + 3 * (1 + 3) / 5", "() + ()", "sin(,x)", "abs(-5,)", "isudf(1)", "x^", "^x"}	
	for _, v := range exprs {
		expr, err := ParseExpression(v)
		if err == nil {
			t.Errorf("parsed %v; Got %v, but expected error.", v, expr)
		}
	}
}