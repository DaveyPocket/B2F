//	Basically F*cked
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

type Table []string
var symbolTable Table

func (t *Table) insert(name string) {
	for _, val := range *t {
		if name == val {
			return
		}
	}
	*t = append(*t, name)
}

type Node struct {
	token	baselex.Token	//	Could an interface go into here somehow?
	child	[]Node
}

func (n Node) GetChild(index int) Node {
	return n.child[index]
}

func (n Node) GetTokName() baselex.TokName {
	return n.token.GetName()
}

func (n Node) GetTokVal() string {
	return n.token.GetLexeme()
}

type root []Node		// Root node of the program

func (r root) PPrint() {
	fmt.Println("Program")
	for _, n := range r {
		printNode(n, 0)
	}
	fmt.Println("End Program")
}

func printNode(n Node, d int) {
	for i := d; i > 0; i-- {
		fmt.Print("-")
	}
	fmt.Print("-")
	fmt.Println(n.GetTokVal())
	if len(n.child) > 0 {
		d++
		for _, c := range n.child {
			printNode(c, d)
		}
	}
}

func (n Node) Is(name baselex.TokName) bool {
	return name == n.GetTokName()
}

//	Keep creating child node until we reach a newline character.
func Parse(r io.Reader) (*root, Table) {
	program := &root{}
	baselex.SetReader(bufio.NewReader(r))
	for n := statementBuilder(&Node{}); !n.Is(baselex.EOF); n = statementBuilder(&Node{}) {
		if n.GetTokName() != baselex.NEWLINE {
			*program = append(*program, *n)
		}
	}
	fmt.Println("Program: ", program)
	fmt.Println("Table: ", symbolTable)
	return program, symbolTable
}

func statementBuilder(n *Node) *Node {
	//	return n if none of the cases match, 
	//	Commented thing in recursive function argument used to denote that function can be called with arbitrary input.
	parent := &Node{token: nextToken()}
	switch parent.GetTokName() {
	case baselex.LET:
		//	call statementBuilder and pass return value into temporary storage.
		//	Check temporary store against conditions that are valid. Throw error if it does not work.
		parent.child = append(parent.child, *statementBuilder(&Node{}))
		return parent
	case baselex.IDENTIFIER:
		symbolTable.insert(parent.GetTokVal())
		return statementBuilder(parent)
	case baselex.EQU:
		parent.child = append(parent.child, *n, *assignmentBuilder(&Node{}))
	case baselex.FOR:
		parent.child = append(parent.child, *forBuilder(&Node{}), *statementBuilder(&Node{}))
		return parent
	case baselex.IF:
		parent.child = append(parent.child, *ifBuilder(&Node{}), *statementBuilder(&Node{}))
		return parent
	case baselex.PRINT:
		parent.child = append(parent.child, *printBuilder(&Node{}))
	case baselex.NEXT:
		parent.child = append(parent.child, Node{token: nextToken()})
	}
	return parent
}

func printBuilder(n *Node) *Node {
	parent := &Node{token: nextToken()}
	if !parent.Is(baselex.IDENTIFIER) {
		panic("Expected identifier in PRINT statement")
	}
	return parent
}

func forBuilder(n *Node) *Node {
	parent := &Node{token: nextToken()}
	switch parent.GetTokName() {
	case baselex.IDENTIFIER:
		symbolTable.insert(parent.GetTokVal())
		return forBuilder(parent)
	case baselex.EQU:
		parent.child = append(parent.child, *n, *assignmentBuilder(&Node{}))
		return parent
	}
	return n
}

func ifBuilder(n *Node) *Node {
	parent := &Node{token: nextToken()}
	switch parent.GetTokName() {
	case baselex.IDENTIFIER:
		symbolTable.insert(parent.GetTokVal())
		return ifBuilder(n)
	case baselex.GREATERTHAN: //	Add more
	}
	return parent
}

//	IsKeyword
func assignmentBuilder(n *Node) *Node {
	parent := &Node{token: nextToken()}
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
//	default:
//		str := "Expected Identifier or Expression, found " + parent.GetTokVal() + " instead."
//		panic(str)
	}
	return n
}
