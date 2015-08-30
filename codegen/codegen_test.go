package codegen

import (
	"testing"
	"fmt"
	"os"
)

func TestAdd(t *testing.T) {
	add(0, 2, 4)
	fmt.Println(output)
}

func TestCompile(t *testing.T) {
	f, err := os.Open("evenmore.bas")
	if err != nil {
		panic(err)
	}

	output, err := Compile(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
