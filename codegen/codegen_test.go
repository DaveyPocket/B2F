package codegen

import (
	"testing"
	"os"
	"fmt"
)

func TestAssignGen(t *testing.T) {
	f, err := os.Open("in.bas")
	if err != nil {
		panic(err)
	}
	fmt.Println(MakeProgram(f))
}
