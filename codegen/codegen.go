//	Global variable declaration during runtime.
//	No need to keep initial value in the symbol table.
/* Notes
DO NOT manually keep track of memory pointer. Let the Code Generator perform those operations.
Reading from locations in memory!
When typing literal pointer increments and decrments, make sure bfPointer matches location.

*/
package codegen

import (
	"b2f/baselex"
	"b2f/baseparse"
	"fmt"
	"io"
)

type stackType []int
var symbolTable baseparse.Table

func (s *stackType) push(n int) {
	*s = append(*s, n)
}

func (s stackType) peek() int {
	n := s[len(s)-1]
	fmt.Println(n)
	return n
}

func (s *stackType) pop() int {
	q := *s
	n := q[len(q)-1]
	q = q[:len(q)-1]
	*s = q
	return n
}

type byteType struct {
	location int
	locationPrime int
}

//	codegen package acts as a VM for BF.
var bfPointer int      //	Pointer to current location in memory.
var bfStackPointer int //	Pointer to top of a stack
var vmStack stackType  //	Stack
var output string      //	Output BF program.
/*
func _g_MoveTo() {
}

func g_Add(lVal, rVal, store byteType) {
	
}

func g_Sub() {
}

func g_Div() {
}

func g_Mul() {
}

func g_Print() {
}
*/

//lVal, rVal, result are all locations
func add(lVal, rVal, result int) {
	fmt.Println(lVal, rVal, result)
	//	This function makes the assumption that lVal, rVal, and result all contain complementary locations immediately to the right.
	output += moveTo(bfPointer, lVal) + "[->+<"
	//	bfPointer is still at lVal.
	output += moveTo(bfPointer, result) + "+" + moveTo(bfPointer, lVal) + "]"
	output += ">[-<+>]<"	//	Restore lVal.
	output += moveTo(bfPointer, rVal) + "[->+<"
	//	bfPointer is still at lVal.
	output += moveTo(bfPointer, result) + "+" + moveTo(bfPointer, rVal) + "]"
	output += ">[-<+>]<"	//	Restore rVal.
	output += moveTo(bfPointer, result)
}

//	TODO (Brad) - Remove 'from' statement. Always from bfPointer.
func moveTo(from, to int) (out string) {
	bfPointer = to
	n := to - from
	if n > 0 {
		for i := 0; i < n; i++ {
			out += ">"
		}
	} else if n < 0 {
		for i := 0; i < -n; i++ {
			out += "<"
		}
	}
	return
}

//	Change this
//func thing(n node) {
//	switch n.GetTokName() {
//	case baselex.PRINT:
//		g_Print( /*Zeroth child*/ )
//	}
//}
//	Organize symbol table function
//		Constants have 'constant' type, no location

func organizeTable() {
	//for i, v := range symbolTable {

	//}
}

func getLocation(n baseparse.Node) int {
	for i, val := range symbolTable {
		fmt.Println(val, i, n.GetTokVal())
		if n.GetTokVal() == val {
		fmt.Println(val, i)
			return i * 2
		}
	}
	return -1
}

//	End-all function

func Compile(r io.Reader) (bf string, cErr error) {
	output = ""
	bfPointer = 0
	program, n := baseparse.Parse(r)
	symbolTable = n
	for _, v := range *program {
		switch v.GetTokName() {
		case baselex.EQU:
			switch v.GetChild(1).GetTokName() {
			case baselex.ADD:
				add(getLocation(v.GetChild(1).GetChild(0)), getLocation(v.GetChild(1).GetChild(1)), getLocation(v.GetChild(0)))
			}
		}
	}
	bf = output
	return
}
