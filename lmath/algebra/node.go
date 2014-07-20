package algebra

import (
	"math/big"
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
	m.nodes[1] = n1
	m.nodes[2] = n2
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
	m.data = id
	return m
}

func (n *Node) Replacef64(identifier string, value Scalar) {

}

