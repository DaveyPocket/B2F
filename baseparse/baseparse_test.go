package baseparse

import (
	"testing"
	"os"
)
/*
func TestBuildTree(t *testing.T) {
	p := &parser
}
*/

func TestParse(t *testing.T) {
	f, err := os.Open("in.bas")
	if err != nil {
		panic(err)
	}
	Parse(f)
}
