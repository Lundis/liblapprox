package expr

import (
	"math/big"
	"errors"
	"fmt"
	"strings"
)

func ParseExpression(expr string) (*Node, error) {
	// remove spaces
	chars := strings.Replace(expr, " ", "", -1)
	top_node, err := parse([]byte(chars))
	return top_node, err
}

func parse(expr []byte) (*Node, error) {
	if len(expr) == 0 {
		return nil, errors.New("Syntax Error!")
	}
	return parsePlus(expr)	
}

func parsePlus(expr []byte) (*Node, error) {
	
	plus, minus := tokenize(expr, '+', '-')
	if len(plus) == 1 && len(minus) == 0 {
		// Nothing was found, move on to multiplication
		return parseMult(expr)
	}
	nodes_plus, err := parseExpressions(plus)
	if err != nil {
		return nil, err
	}
	nodes_minus, err := parseExpressions(minus)
	if err != nil {
		return nil, err
	}

	for _, node := range nodes_minus {
		nodes_plus = append(nodes_plus, NewMinusNode(node))
	}
	return NewPlusNode(nodes_plus), nil

}

// Parses multiplication and division
func parseMult(expr []byte) (*Node, error) {
	if len(expr) == 0 {
		return nil, errors.New("Syntax Error!")
	}
	mult, div := tokenize(expr, '*', '/')
	if len(mult) == 1 && len(div) == 0 {
		// Nothing was found, move on to parentheses
		return parseParentheses(expr)
	}
	nodes_mult, err := parseExpressions(mult)
	if err != nil {
		return nil, err
	}
	n1 := nodes_mult[0]
	if len(nodes_mult) != 1 {
		n1 = NewMultNode(nodes_mult)
	}

	if len(div) == 0 {
		return n1, nil
	} else {
		nodes_div, err := parseExpressions(div)
		if err != nil {
			return nil, err
		}
		n2 := nodes_div[0]
		if len(div) != 1 {
			n2 = NewMultNode(nodes_div)
		}
		div_node := NewDivNode(n1, n2)
		return div_node, nil
	}
}

func parseParentheses(expr []byte) (*Node, error) {
	if expr[0] == '(' && expr[len(expr) - 1] == ')' {
		return parse(expr[1:len(expr) - 1])
	} else {
		return parseAtom(expr)
	}
}

func parseAtom(expr []byte) (*Node, error) {
	if '0' <= expr[0] && expr[0] <= '9' || expr[0] == '.' {
		return parseNumber(expr)
	} else {
		return parseVariable(expr)
	}
}

func parseNumber(expr []byte) (*Node, error) {
	rat := big.NewRat(1,1)
	_, success := rat.SetString(string(expr))
	if !success {
		return nil, errors.New(fmt.Sprintf("Failed to convert %v to a number", string(expr)))
	} else {
		return NewRatNode(rat), nil
	}
}

func parseVariable(expr []byte) (*Node, error) {
	for _, c := range expr {
		if !(('a' <= c && c <= 'z')||('A' <= c && c <= 'Z')) {
			message := fmt.Sprintf("Variable %v is not an alphabetic sentence.", string(expr))
			return nil, errors.New(message)
		}
	}
	return NewVarNode(string(expr)), nil
}

// split returns the indices where expr[index] == op1 and where expr letter2, taking into account depth (parentheses).
func tokenize(expr []byte, op1, op2 byte) ([][]byte, [][]byte) {
	depth := 0
	indices1 := make([][]byte, 0, len(expr)/2)
	indices2 := make([][]byte, 0, len(expr)/2)
	start := 0
	which := 1
	if expr[0] == op2 {
		which = 2
		expr = expr[1:]
	}

	for i, v := range expr {
		if v == '(' {
			depth++
		} else if v == ')' {
			depth--
		} else if depth == 0 && (v == op1 || v == op2) {
			if which == 1 {
				indices1 = append(indices1, expr[start:i])
			} else {
				indices2 = append(indices2, expr[start:i])
			}

			if v == op1 {
				which = 1
			} else {
				which = 2
			}
			start = i+1

		}
	}
	// last one
	if which == 1 {
		indices1 = append(indices1, expr[start:])
	} else {
		indices2 = append(indices2, expr[start:])
	}

	return indices1, indices2
}

func parseExpressions(exprs [][]byte) ([]*Node, error) {
	nodes := make([]*Node, len(exprs))
	var err error
	for i := 0; i < len(exprs); i++ {
		nodes[i], err = parse(exprs[i])
		if err != nil {
			return nil, err
		}
	}
	return nodes, nil
}