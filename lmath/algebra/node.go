package algebra

type Node struct {
	op Operator
	nodes []Node
}

func (n *Node) Replacef64(identifier string, value Scalar) {

}

