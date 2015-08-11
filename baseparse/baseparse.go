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

type table []string
var symbolTable table

<<<<<<< HEAD
//	GetStatements
func (p parser) getLines() (lines []node) {
	m := p.treeBuilder(&node{}) // Get a node from the treeBuilder
	for ; !m.isDelim(); m = p.treeBuilder(&node{}) {
		lines = append(lines, *m)
	}
	lines = append(lines, *m)
	return
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

func (p parser) treeBuilder(n *node) (*node) {
	t := p.scan()
	fmt.Println(t)
	root := &node{tok: t}
	switch t.GetName() {
	case baselex.StringToName("IDENTIFIER"):
		n = p.treeBuilder(root)
	case baselex.StringToName("="):
		// Desired root node
		root.children = append(root.children, *n, *p.assignmentBuilder(root))
		p.symbol = append(p.symbol, symbol{root.children[0].tok.GetVal(), root.children[1].tok.GetVal()})
		return root
	case baselex.StringToName("LET"):
		root.children = append(root.children, *p.treeBuilder(root))
		return root
	case baselex.StringToName("PRINT"):
		root.children = append(root.children, *p.treeBuilder(root))
		if root.children[0].tok.GetName() == baselex.StringToName("=") {
			fmt.Println("Assignment not allowed in PRINT statement.")
			os.Exit(1)
		}
		return root
	case baselex.StringToName("FOR"):
		root.children = append(root.children, *p.treeBuilder(root))	//	Append conditional to leftmost branch
		root.children = append(root.children, p.getLines()...) // Append statements to not-leftmost branch
		if !getDelim(root.children).isNext() {
			fmt.Println("Expected NEXT, found", getDelim(root.children).tok.GetVal(), "instead.")
			os.Exit(1)
=======
func (t *table) insert(name string) {
	for _, val := range *t {
		if name == val {
			return
>>>>>>> newParse
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
	case baselex.PRINT:
		parent.child = append(parent.child, *printBuilder(&node{}))
	}
	return parent
}

func printBuilder(n *node) *node {
	parent := &node{token: nextToken()}
	if !parent.Is(baselex.IDENTIFIER) {
		panic("Expected identifier in PRINT statement")
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
