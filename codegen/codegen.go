/* Notes
DO NOT manually keep track of memory pointer. Let the Code Generator perform those operations.
*/
package codegen

import (
	//"b2f/baselex"
	"b2f/baseparse"
	"fmt"
	"io"
	//"bufio"
)

type varType struct {
	name		string
	location	int
	value int		//SHOULD ONLY BE USED IN THE INITVAR FUNCTION
}

type bookShelf struct {
	AST		*baseparse.Program
	Pointer	int
	DynamicStart	int
	varTable	[]varType
	program string
}

func MakeProgram(r io.Reader) (string) {
	b := &bookShelf{}
	b.AST = baseparse.Parse(r)
	b.makeStaticRegion()
	return b.program
}

func (b *bookShelf) makeStaticRegion() {
	c, n := b.AST.GetVariables()
	fmt.Println(c)
	for _, v := range n {
		b.initVar(varType{name: v.GetChild(0).GetTokVal(), value: v.GetChild(1).GetTokVal()})
	}
}

func (b *bookShelf) initVar(n varType) {
	n.location = b.Pointer
	b.varTable = append(b.varTable, n)
	b.program += b.assignGen(n.value)
}

func (b *bookShelf) assignGen(val int) (string) {
	b.Pointer = b.DynamicStart	// May be redundant because of recursive call
	if val > 0 {
		assign := "+"
		assign += b.assignGen(val - 1)
		return assign
	}
	b.DynamicStart += 2
//	b.Pointer += 2
	return "\n>>\n"
}

func addGen(aVar, bVar varType) (string) {
	//	Generate correct number of '<' or '>' based on current pointer position - aLoc
	//	Loop generator
	//out := countGen(aVar
	return ""
}

func (b *bookShelf) countGen(varToCount varType, body string) (string) {
	//	TODO (Brad): replacement body parameter with something else...
	//	... Not quite sure what yet... Not good.
	//	Also, optimize by utilizing secondary loop
	loop := b.moveTo(varToCount)
	loop += "[->+<\n"	//	Start of a loop
	loop += body
	loop += "]\n[>-<+\n]\n"
	return loop
}

func (b *bookShelf) moveTo(v varType) (string) {
	b.Pointer = v.location
	return moveGen(v.location - b.Pointer)
}

func moveGen(r int) (string) {
	if r > 0 {
		inc := ">"
		inc += moveGen(r - 1)
		return inc
	} else if r < 0 {
		dec := "<"
		dec += moveGen(r + 1)
		return dec
	}
	return "\n"
}

/*
Generate members for different types
Type variable has generator function.
--Generator takes no arguments and only returns string code
--Generator function knows variable value
*/
