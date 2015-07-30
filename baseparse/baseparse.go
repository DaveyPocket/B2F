package baseparse
// BASIC to BF parser

import (
		"b2f/baselex"
		"fmt"
		"io"
		"bufio"
		"os"
)

func (p parser) scan() (baselex.Token) {
	return p.lex.ScanToken()
}

func getDelim(n []node) (node) {
	return n[len(n) - 1]
}

func (p parser) getLines() (lines []node) {
	m := p.treeBuilder(&node{}) // Get a node from the treeBuilder
	for ; !m.isDelim(); m = p.treeBuilder(&node{}) {
		lines = append(lines, *m)
	}
	lines = append(lines, *m)
	return
}

func (p parser) buildTree() (*program) {
	prg := &program{}
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
		n = p.treeBuilder(&node{tok: t})
	case baselex.StringToName("="):
		// Desired root node
		root.children = append(root.children, *n, *p.assignmentBuilder(root))
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
		}
		return root
	case baselex.StringToName("IF"):
		root.children = append(root.children, *p.treeBuilder(root))
		root.children = append(root.children, p.getLines()...)
		if !getDelim(root.children).isElse() {
			return root
		}
		root.children = append(root.children, p.getLines()...)
		return root
	case baselex.StringToName("THEN"):
		return n
	case baselex.StringToName("ELSE"):
		return &node{tok: t}
		//	Then/else are used as a delimiter to denote the end of a if clause.
	case baselex.StringToName("ENDIF"):
		//	Signals the end of an if statement
		fallthrough
	case baselex.StringToName("END"):
		//return &node{tok: t}
		fallthrough
	case baselex.StringToName(""):
		return &node{tok: t}
	case baselex.StringToName("NEXT"):
		root.children = append(root.children, *p.treeBuilder(root))
		return root
	}
	fmt.Println("End of switch block....")
	return n
}

func (p parser) assignmentBuilder(n *node) (*node) {
	m := n
	t := p.scan()
	fmt.Println(t)
	switch t.GetName() {
	case baselex.StringToName("IDENTIFIER"):
		m = p.assignmentBuilder(&node{tok: t})
	case baselex.StringToName("+"), baselex.StringToName("-"), baselex.StringToName("*"), baselex.StringToName("/"), baselex.StringToName("TO"):
		// Desired root node
		root := &node{tok: t}
		root.children = append(root.children, *n, *p.assignmentBuilder(root))
		return root
	case baselex.StringToName("="):
		panic("Double assignment")
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

func (n node) isEOF() (bool) {
	return n.tok.GetName() == baselex.StringToName("")
}

func (n node) isNext() (bool) {
	return n.tok.GetName() == baselex.StringToName("NEXT")
}

func (n node) isEndIf() (bool) {
	return n.tok.GetName() == baselex.StringToName("ENDIF")
}

func (n node) isEnd() (bool) {
	return n.tok.GetName() == baselex.StringToName("END")
}

func (n node) isThen() (bool) {
	return n.tok.GetName() == baselex.StringToName("THEN")
}

func (n node) isElse() (bool) {
	return n.tok.GetName() == baselex.StringToName("ELSE")
}

func (n node) isNewline() (bool) {
	return n.tok.GetName() == baselex.StringToName("\n")
}

func (n node) isDelim() (bool) {
	return n.isEndIf() || n.isEOF() || n.isNext() || n.isEnd() || n.isThen() || n.isElse() || n.isNewline()
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
