package b2f/baseparse
// BASIC to BF parser

import ("b2f/baselex"
		"fmt")

type parse struct {
	//	some root node for AST
	//	Current token to do things to?
	//	First - What does the structure of a node look like?
	*program
}

type node struct {
	tok	baselex.Token	//	Could an interface go into here somehow?
	children	[]node
}

type program []node	// Root node of the program

func (p *program) add(t baselex.Token) {
	p = append(p, node)// Might not work....
	//	Can items be appended to a slice even if the member receiver is of pointer type?
}

func (n *node) add(t baselex.Token) {
	n.children = append(n.children, node)// Might not work.....
}

type adder interface {
	add(t baselex.token)
}
