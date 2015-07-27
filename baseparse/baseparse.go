package baseparse
// BASIC to BF parser

import (
		"b2f/baselex"
		"fmt"
		"io"
		"bufio"
)

func (p parser) scan() (baselex.Token) {
	return p.lex.ScanToken()
}

func (p parser) buildTree() (*program) {
	// TODO: Change return type to pointer to a program
	prg := &program{}
	m := p.treeBuilder(&node{})
	//fmt.Printf("%+v\n", *m)
	*prg = append(*prg, *m)
	return prg
}

func (p parser) treeBuilder(n *node) (*node) {
	m := n
	t := p.scan()
	fmt.Println(t)
	switch t.GetName() {
	case baselex.StringToName("IDENTIFIER"):
		m = p.treeBuilder(&node{tok: t})
	case baselex.StringToName("="), baselex.StringToName("+"), baselex.StringToName("-"), baselex.StringToName("*"), baselex.StringToName("/"):
		// Desired root node
		root := &node{tok: t}
		root.children = append(root.children, *n, *p.treeBuilder(root))
		return root
	}
	return m
}

func Parse(r io.Reader) {
	p := &parser{lex: baselex.NewLex(bufio.NewReader(r))}
	// Loop through multiple branches to build program
	p.prog = p.buildTree()
	fmt.Println("Program:", *p.prog)
}

type parser struct {
	lex *baselex.Lex
	//	some root node for AST
	//	Current token to do things to?
	//	First - What does the structure of a node look like?
	prog *program
}


type node struct {
	//	Node contain
	tok	baselex.Token	//	Could an interface go into here somehow?
	children	[]node
}

type program []node	// Root node of the program

func (p *program) add(t baselex.Token) {
//	p = append(p, node)// Might not work....
	//	Can items be appended to a slice even if the member receiver is of pointer type?
}

func (n *node) add(t baselex.Token) {
//	n.children = append(n.children, node)// Might not work.....
}

type adder interface {
	add(t baselex.Token)
}

// Part of the semantics phase???
type forLoop struct {
	//	Condition should only contain some form of expression?
	condition	node
	statements	[]node
}

//	leaf
type let struct {
	lVal		node	//Variable type???
	rVal		node	//Expression type??
}
