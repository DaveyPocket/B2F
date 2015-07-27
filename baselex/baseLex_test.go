package baselex

import ("testing"
			"fmt"
			"strings"
			"bufio"
		)

func TestToValidToken(t *testing.T) {
	fmt.Println(toValidToken(token{IDENTIFIER, "FOR"}))
}

func TestReadUntilEnd(t *testing.T) {
	//	Create some sort of test string and pass it to the lex type using the strings.Reader thing.....
	// string to test "word    "
	l := NewLex(bufio.NewReader(strings.NewReader("WORD     ")))
	f := NewLex(bufio.NewReader(strings.NewReader("WORD")))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from fault test")
		}
	}()
	fmt.Println(l.readUntilEnd([]rune{}) + "/end")
	fmt.Println(f.readUntilEnd([]rune{}) + "/end")
}

func TestReadSpace(t *testing.T) {
	l := NewLex(bufio.NewReader(strings.NewReader("")))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from fault test")
		}
	}()
	fmt.Println(l.readSpace())
}

func TestReadLiteral(t *testing.T) {
	l := NewLex(bufio.NewReader(strings.NewReader("  THE COW JUMPED OVER THE MOON\n")))
	//ReadLiteral should also return a boolean if it is still able to read?
	for i := 0; i < 6; i++ {
		fmt.Println(l.readLiteral() + "/end")
	}
}

func TestScanIdent(t *testing.T) {
	l := NewLex(bufio.NewReader(strings.NewReader(" THE COW JUMPED OVER THE FOR LOOP")))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from fault test")
		}
	}()
	for i := 0; i < 7; i++ {
		fmt.Println(l.scanIdent())
	}
}

func TestScanToken(t *testing.T) {
	l := NewLex(bufio.NewReader(strings.NewReader("LET THE COW PRINT NEXT TO THE WHITESPACE")))
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from fault test")
		}
	}()
	for i := 0; i < 8; i++ {
		fmt.Println(l.ScanToken())
	}
}
