package baseparse
// BASIC to BF parser

import (
		"b2f/baselex"
		"fmt"
		"io"
		"bufio"
		"os"
)

type symbol struct {
	name string
	init	int
}

type table []symbol	//	Symbol table type

func (t *table) insert(name string, init int) {
	t = append(t, symbol{name, init})
}

nextToken := baselex.ReadToken
symbolTable := table

type node struct {
	token	baselex.Token	//	Could an interface go into here somehow?
	child	[]node
}

func (n node) GetTokName() baselex.TokName {
	return n.token.GetName()
}

type root []node		// Root node of the program

type program struct {
	*root			//	Root node
}

func (n node) Is(name baselex.TokName) bool {
	return name == n.GetTokName()
}

func Parse(r *bufio.Reader) *program {
	baselex.SetReader(r)
	instr, symTable := makeAST()
	return &program{/*stuff in here*/}
}

//	Keep creating child node until we reach a newline character.
func makeAST() *root, table {
	p := &root{}
	/*
	for t := nextToken(), !t.IsEOF(); t = nextToken() {
		switch t.GetName() {
		case baselex.LET:
		}
	}*/
}

func treeBuilder(n *node) node {
	switch n {
	case n.Is(baselex.LET):
		return assignmentBuilder
	}
}

//	IsKeyword
func assignmentBuilder(n *node) *node {
	parent:= &node{token: nextToken()}
	switch parent.GetTokName() {
	/*case n.Is(baselex.LET):
		panic("Found keyword LET, expected named type literal.")
	case n.Is(baselex.GOTO):
		panic("Found keyword GOTO, expecred named type literal.")*/
	case baselex.IDENTIFIER:
		return assignmentBuilder(parent)
	case baselex.EQU:
		//	Root node is assignment operator.
		parent.child = append(parent.child, n, assignmentBuilder(n))
		//	The code below should only work if the right child is a constant number.
		symbolTable.insert(parent.child[0].GetTokVal(), parent.child[1].GetTokVal())
		return root
	}
	return root
}

func (p parser) buildTree() (*Program) {
	prg := &Program{}
	*prg = append(*prg, p.getLines()...)
	if !getDelim(*prg).isEOF() && !getDelim(*prg).isEnd() {
		fmt.Println("Expected <EOF>, found", getDelim(*prg).tok.GetVal(), "instead.")
		os.Exit(1)
	}
	return prg
}

