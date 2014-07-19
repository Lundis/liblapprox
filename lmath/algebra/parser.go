package algebra

func ParseExpression(expr string) *Expression {
	// remove spaces
	chars := strings.Replace(expr, " ", "", -1)
	top_node := parse(chars)
}

func parse(expr []byte) *Node {
	
}

//
func parsePlus(expr []byte) *Node {
	terms := make([]Node)

	node :=
	return node
}
func parseParenthesis(expr []byte) *Node {
	
}

func parseMult(expr []byte) *Node {

}

func parseDiv(expr []byte) *Node {

}