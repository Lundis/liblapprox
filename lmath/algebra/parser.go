package algebra

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
	return parsePlus(expr)	
}

func parsePlus(expr []byte) (*Node, error) {
	
	plus, minus := tokenize(expr, '+', '-')
	if len(plus) == 1 && len(minus) == 0 {
		// Nothing was found, move on to multiplication
		return parseMult(expr)
	}
	nodes := make([]*Node, len(plus) + len(minus))
	for i := 0; i < len(plus); i++ {
		n, err := parseMult(plus[i])
		if err != nil {
			return nil, err
		}
		nodes[i] = n
	}
	for i := 0; i < len(minus); i++ {
		m, err := parseMult(minus[i])
		if err != nil {
			return nil, err
		}
		nodes[len(plus) + i] = NewMinusNode(m)
	}
	return NewPlusNode(nodes), nil

}

func parseMult(expr []byte) (*Node, error) {
	mult, div := tokenize(expr, '*', '/')
	if len(mult) == 1 && len(div) == 0 {
		// Nothing was found, move on to parentheses
		return parseParentheses(expr)
	}
	nodes_mult := make([]*Node, len(mult))

	var err error
	for i := 0; i < len(mult); i++ {
		nodes_mult[i], err = parse(mult[i])
		if err != nil {
			return nil, err
		}
	}
	n1 := nodes_mult[0]
	if len(nodes_mult) != 1 {
		n1 = NewMultNode(nodes_mult)
	}

	if len(div) == 0 {
		return n1, nil
	} else {
		nodes_div := make([]*Node, len(div))
		for i := 0; i < len(div); i++ {
			nodes_div[i], err = parse(div[i])
			if err != nil {
				return nil, err
			}
		}
		n2 := NewMultNode(nodes_div)
		return NewDivNode(n1, n2), nil
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
		return nil, errors.New(fmt.Sprintf("Error: failed to convert %v to a number", string(expr)))
	} else {
		return NewRatNode(rat), nil
	}
}

func parseVariable(expr []byte) (*Node, error) {
	c := expr[0]
	if len(expr) != 1 {
		return nil, errors.New(fmt.Sprintf("Error: long (len>1) variable name: %v", string(expr)))
	} else if !(('a' <= c && c <= 'z')||('A' <= c && c <= 'Z')) {
		return nil, errors.New(fmt.Sprintf("Error: variable is not a letter: %v", c))
	} else {
		return NewVarNode(string(expr)), nil
	}
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
	}

	for i, v := range expr {
		if v == '(' {
			depth++
		} else if v == ')' {
			depth--
		} else if (depth == 0 && (v == op1 || v == op2)) || (i == len(expr) - 1) {
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

	return indices1, indices2
}