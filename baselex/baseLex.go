/*
	This package performs Lexical Analysis for BASIC
*/
package b2f/baselex

import (
	"bufio"
)
// String member for tokName...
type tokName int

const (
	IDENTIFIER tokName = iota

	//	Keywords
	LET
	FOR
	PRINT
	NEXT
	TO

	//	Operators
	EQU
	ADD
	SUB
	MUL
	DIV

	//	Expressives
	LNUM
	QUOTE
	WHITESPACE
)

var reservedTok = map[string]tokName{
	"LET":			LET,
	"FOR":        FOR,
	"PRINT":      PRINT,
	"NEXT":       NEXT,
	"TO":         TO,
	"EQU":        EQU,
	"ADD":        ADD,
	"SUB":        SUB,
	"MUL":        MUL,
	"DIV":        DIV,
	"QUOTE":      QUOTE,
	"WHITESPACE": WHITESPACE,
}

type Token struct {
	name tokName
	val  string
}

type lex struct {
	b *bufio.Reader
}

func NewLex(b *bufio.Reader) (*lex) {
	return &lex{b}
}

func toValidToken(tk Token) Token {
	return Token{name: reservedTok[tk.val], val: tk.val}
}

func (l *lex) readSpace() (bool) {
	r, _, err := l.b.ReadRune()
	if err != nil {
		panic(err)
	}

	if r != ' ' {
		l.b.UnreadRune()
		return false
	}

	return true
}

func (l *lex) readSpaces() {
	for l.readSpace() {}
}
func (l *lex) readUntilEnd(m []rune) string {
	// Read space must go in here
	r, _, err := l.b.ReadRune()
	if err != nil {
		panic(err)
	}

	if r != ' ' && r != '\n' && r != rune(0) {
		m = append(m, r)
		return l.readUntilEnd(m)
	}

	//l.b.UnreadRune()
	return string(m)
}

func (l *lex) readLiteral() string {
	l.readSpaces()
	return l.readUntilEnd([]rune{})
}

func (l *lex) scanIdent() Token {
	//	Better way to write this?
//	for b := l.readSpace(); b; b = l.readSpace() {
//	}
	return Token{name: IDENTIFIER, val: l.readLiteral()}
}

func (l *lex) ScanToken() Token {
	bI := l.scanIdent() // Base Identifier. Identifier until proven otherwise
	return toValidToken(bI)
}
