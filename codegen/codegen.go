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

//	Named variable type
type varType struct {
	name     string
	location int
	value    interface{} //SHOULD ONLY BE USED IN THE INITVAR FUNCTION
}

//	Conditional Type
type condType struct {
}

type forType struct {
}

//	No need for a bookshelf type
//	There only needs to exist one code generator a moment in time.
type bookShelf struct {
	AST          *baseparse.Program
	Pointer      int
	DynamicStart int
	varTable     []varType
	program      string
}

func MakeProgram(r io.Reader) string {
	b := &bookShelf{}
	b.AST = baseparse.Parse(r)
	b.makeStaticRegion()
	return b.program
}

func (b *bookShelf) makeStaticRegion() {
	c, n := b.AST.GetSymTable()
	fmt.Println(c)
	for _, v := range n {
		b.initVar(varType{name: v.GetName(), value: v.GetVal()})
	}
}

func (b *bookShelf) initVar(n varType) {
	n.location = b.Pointer
	b.varTable = append(b.varTable, n)
	b.program += b.assignGen(n.value)
}

func (b *bookShelf) assignGen(val int) string {
	b.Pointer = b.DynamicStart // May be redundant because of recursive call
	if val > 0 {
		assign := "+"
		assign += b.assignGen(val - 1)
		return assign
	}
	b.DynamicStart += 2
	//	b.Pointer += 2
	return "\n>>\n"
}

func addGen(aVar, bVar varType) string {
	//	Generate correct number of '<' or '>' based on current pointer position - aLoc
	//	Loop generator
	//out := countGen(aVar
	return ""
}

func (b *bookShelf) countGen(varToCount varType, body string) string {
	//	TODO (Brad): replacement body parameter with something else...
	//	... Not quite sure what yet... Not good.
	//	Also, optimize by utilizing secondary loop
	loop := b.moveTo(varToCount)
	loop += "[->+<\n" //	Start of a loop
	loop += body
	loop += "]\n[>-<+\n]\n"
	return loop
}

func (b *bookShelf) moveTo(v varType) string {
	b.Pointer = v.location
	return moveGen(v.location - b.Pointer)
}

func moveGen(r int) string {
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

//	Should be duplicate generator
func (b *bookShelf) duplicate(source, destination varType) (out string) {
}

func (b *bookshelf) printGen(t interface{}) string {
	switch t.(type) {
	case varType:
		switch t.value.(type){
		case int:
		}
	}
	//	ASCII character code offset for integer type, (Modulo for digits)
	//	String literals - stack style
	//	Type switch
}

/*
Generate members for different types
Type variable has generator function.
--Generator takes no arguments and only returns string code
--Generator function knows variable value
*/

//	Generator interface????
