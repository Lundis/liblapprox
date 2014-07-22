package algebra

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type Operator int

const (
	PLUS Operator = iota
	MINUS
	MULT
	DIV
	VAR
	RAT
)

type Node struct {
	op Operator
	nodes []*Node
	data Atom
}

func (self *Node) childrenStrings() []string {
	strs := make([]string, len(self.nodes))
	for i := range strs {
		strs[i] = self.nodes[i].String()
	}
	return strs
}

func (self *Node) String() string {
	switch self.op {
	case PLUS:
		return "(" + strings.Join(self.childrenStrings(), " + ") + ")"
	case MINUS:
		return " - (" + self.nodes[0].String() + ")"
	case MULT:
		return strings.Join(self.childrenStrings(), " * ")
	case DIV:
		return self.nodes[0].String() + " / " + self.nodes[1].String()
	case VAR:
		v := self.data.(Variable)
		return v.Identifier
	case RAT:
		rat := self.data.(ScalarRat).Value
		str := rat.String()
		if rat.IsInt() {
			// remove redundant "/1"
			return str[:len(str)-2]
		} else {
			return str
		}
	default:
		panic(fmt.Sprintf("Unknown Node type in String(): %v", self.op))
	}
}

func NewPlusNode(nodes []*Node) *Node {
	n := new(Node)
	n.op = PLUS
	n.nodes = nodes
	return n
}

// Wraps n in a minus node
func NewMinusNode(n *Node) *Node {
	m := new(Node)
	m.op = MINUS
	m.nodes = make([]*Node, 1)
	m.nodes[0] = n
	return m
}

func NewMultNode(nodes []*Node) *Node {
	n := new(Node)
	n.op = MULT
	n.nodes = nodes
	return n
}

func NewDivNode(n1, n2 *Node) *Node {
	m := new(Node)
	m.op = DIV
	m.nodes = make([]*Node, 2)
	m.nodes[0] = n1
	m.nodes[1] = n2
	return m
}

func NewRatNode(num *big.Rat) *Node {
	m := new(Node)
	m.op = RAT
	m.data = ScalarRat{num}
	return m
}

func NewVarNode(id string) *Node {
	m := new(Node)
	m.op = VAR
	m.data = Variable{id}
	return m
}

func (self *Node) Evaluate(vars map[string] *big.Rat) (*big.Rat, error) {
	switch self.op {
	case PLUS:
		sum := big.NewRat(0, 1)
		for _, v := range self.nodes {
			val, err := v.Evaluate(vars)
			if err != nil {
				return nil, err
			}
			sum.Add(sum, val)
		}
		return sum, nil
	case MINUS:
		inv := big.NewRat(-1, 1)
		val, err := self.nodes[0].Evaluate(vars)
		if err != nil {
			return nil, err
		}
		return inv.Mul(inv, val), nil
	case MULT:
		prod := big.NewRat(1, 1)
		for _, v := range self.nodes {
			val, err := v.Evaluate(vars)
			if err != nil {
				return nil, err
			}
			prod.Mul(prod, val)
		}
		return prod, nil
	case DIV:
		first, err1 := self.nodes[0].Evaluate(vars)
		if err1 != nil {
			return nil, err1
		}
		second, err2 := self.nodes[1].Evaluate(vars)
		if err2 != nil {
			return nil, err2
		}
		result := big.NewRat(1, 1).Set(second)
		result.Inv(result)
		
		return result.Mul(result, first), nil
	case VAR:
		fallthrough
	case RAT:
		return self.data.Evaluate(vars)
	default:
		return nil, errors.New(fmt.Sprintf("Evaluate(): Unknown op: %v", self.op))
	}
}