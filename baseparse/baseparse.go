package baseparse
// BASIC to BF parser

import (
		"b2f/baselex"
		"fmt"
		"io"
		"bufio"
)

func nextToken() baselex.Token {
	t := baselex.ReadToken()
	fmt.Println(t)
	return t
}

type table []string
var symbolTable table

func (t *table) insert(name string) {
	for _, val := range *t {
		if name == val {
			return
		}
	}
	*t = append(*t, name)
}

type node struct {
	token	baselex.Token	//	Could an interface go into here somehow?
	child	[]node
}

func (n node) GetTokName() baselex.TokName {
	return n.token.GetName()
}

func (n node) GetTokVal() string {
	return n.token.GetLexeme()
}

type root []node		// Root node of the program

func (n node) Is(name baselex.TokName) bool {
	return name == n.GetTokName()
}

//	Keep creating child node until we reach a newline character.
func Parse(r io.Reader) (*root, table) {
	program := &root{}
	baselex.SetReader(bufio.NewReader(r))
	for n := statementBuilder(&node{}); !n.Is(baselex.EOF); n = statementBuilder(&node{}) {
		*program = append(*program, *n)
	}
	fmt.Println("Program: ", program)
	fmt.Println("Table: ", symbolTable)
	return program, symbolTable
}

func statementBuilder(n *node) *node {
	//	return n if none of the cases match, 
	//	Commented thing in recursive function argument used to denote that function can be called with arbitrary input.
	parent := &node{token: nextToken()}
	switch parent.GetTokName() {
	case baselex.LET:
		//	call statementBuilder and pass return value into temporary storage.
		//	Check temporary store against conditions that are valid. Throw error if it does not work.
		parent.child = append(parent.child, *statementBuilder(&node{}))
		return parent
	case baselex.IDENTIFIER:
		symbolTable.insert(parent.GetTokVal())
		return statementBuilder(parent)
	case baselex.EQU:
		parent.child = append(parent.child, *n, *assignmentBuilder(&node{}))
	case baselex.FOR:
		parent.child = append(parent.child, *forBuilder(&node{}), *statementBuilder(&node{}))
		return parent
	case baselex.IF:
		parent.child = append(parent.child, *ifBuilder(&node{}), *statementBuilder(&node{}))
		return parent
	}
	return parent
}

func forBuilder(n *node) *node {
	parent := &node{token: nextToken()}
	switch parent.GetTokName() {
	case baselex.IDENTIFIER:
		symbolTable.insert(parent.GetTokVal())
		return forBuilder(n)
	case baselex.EQU:
		parent.child = append(parent.child, *n, *assignmentBuilder(&node{}))
		return parent
	}
	return n
}

func ifBuilder(n *node) *node {
	parent := &node{token: nextToken()}
	switch parent.GetTokName() {
	case baselex.IDENTIFIER:
		symbolTable.insert(parent.GetTokVal())
		return ifBuilder(n)
	case baselex.GREATERTHAN: //	Add more
	}
	return parent
}

//	IsKeyword
func assignmentBuilder(n *node) *node {
	parent:= &node{token: nextToken()}
	switch parent.GetTokName() {
	case baselex.IDENTIFIER:
		symbolTable.insert(parent.GetTokVal())
		return assignmentBuilder(parent)
	case baselex.ADD, baselex.SUB, baselex.MUL, baselex.DIV, baselex.MOD, baselex.TO:
		//	Root node is assignment operator.
		parent.child = append(parent.child, *n, *assignmentBuilder(n))
		//	The code below should only work if the right child is a constant number.
		symbolTable.insert(parent.child[0].GetTokVal())
		return parent
	case baselex.NEWLINE:
		return parent
	default:
		str := "Expected Identifier or Expression, found " + parent.GetTokVal() + " instead."
		panic(str)
	}
	return parent
}
