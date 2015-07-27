package b2f/baseparse
// BASIC to BF parser

import (
		"b2f/baselex"
		"fmt"
		"io"
		"bufio"
)

func getToken() (token) {
	return baselex.ScanToken()
}

func getTokName() {

}

func Parse(r io.Reader) {
	lexer := NewLex(bufio.NewReader(r))
	buildAST()

}

func buildAST() {
	//	Loop until a token with the END keyword is found
	for /*!end*/ {
		n := buildNode()

	}
}

func buildNode() (node) {
	
}

func nodeBuilder(t *lexer.Token) (node) {
	switch t.GetName() {
	case baselex.StringToName("IDENT"):
		n := nodeBuilder(getToken())
		switch n.What() {
		case 1:
			//	Token cannot be a node
			n.lVal = t
		case -1:
			return 
		}
	case baselex.StringToName("EQU"):
		return equ{rVal: nodeBuilder(getToken())}
	case baselex.StringToName("NEWLINE"):
		return newLine
	}

}
/*
func buildNode() (node) {
	n := &node{}
	for tok := getToken(); tok.GetName() != baselex.StringToName("NEWLINE"); tok = getToken() {
		switch tok.GetName(){
		case baselex.StringToName("EQU"):
			//	Now that we have entered the EQU case, need to follow a set of rules to build node
			n.tok = tok
			n
		}
	}

}
*/

type ident lexer.Token

type equ struct {
	lVal,	rVal	node
}

type newLine struct {}

func (e equ) What() (int) {
	return 1
}

func (n newLine) What() (int() {
	return -1
}

type parser struct {
	//	some root node for AST
	//	Current token to do things to?
	//	First - What does the structure of a node look like?
	*program
}

type node interface{
	what() (int)
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
