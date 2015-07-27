/*
	This package performs Lexical Analysis for BASIC
*/
package baselex

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
	END

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
	NEWLINE

	// Things
	OPENPAREN
	CLOSEPAREN
	OPENSQUAREBRACKET
	CLOSESQUAREBRACKET
	OPENBRACKET
	CLOSEBRACKET
	EOF
)

var reservedTok = map[string]tokName{
	"LET":			LET,
	"FOR":        FOR,
	"PRINT":      PRINT,
	"NEXT":       NEXT,
	"TO":         TO,
	"=":        EQU,
	"+":        ADD,
	"-":        SUB,
	"*":        MUL,
	"/":        DIV,
	"\"":      QUOTE,
	" ":			WHITESPACE,
	"END":			END,
	"\n":		NEWLINE,
	"":		EOF,
}
/*
func NameToString(tokName) (string) {
}
*/

func StringToName(str string) (tokName) {
	return reservedTok[str]
}

type Token struct {
	name tokName
	val  string
}

type Lex struct {
	b *bufio.Reader
}

func NewLex(b *bufio.Reader) (*Lex) {
	return &Lex{b}
}

func toValidToken(tk Token) Token {
	return Token{name: reservedTok[tk.val], val: tk.val}
}

func (l *Lex) read() (rune) {
	r, _, err := l.b.ReadRune()
	if err != nil {
		//panic(err)
		r = rune(0)
	}
	return r
}

func (l *Lex) readSpace() (bool) {
	r := l.read()
	if r != ' ' {
		l.b.UnreadRune()
		return false
	}

	return true
}

func (l *Lex) readSpaces() {
	for l.readSpace() {}
}
func (l *Lex) readUntilEnd(m []rune) string {
	// Read space must go in here
	r := l.read()

	if r != ' ' && r != '\n' && r != rune(0) {
		m = append(m, r)
		return l.readUntilEnd(m)
	}

	//l.b.UnreadRune()
	return string(m)
}

func (l *Lex) readLiteral() string {
	l.readSpaces()
	return l.readUntilEnd([]rune{})
}

func (l *Lex) scanIdent() Token {
	//	Better way to write this?
//	for b := l.readSpace(); b; b = l.readSpace() {
//	}
	return Token{name: IDENTIFIER, val: l.readLiteral()}
}

func (t *Token) GetName() (tokName) {
	return t.name
}

func (t *Token) GetVal() (string) {
	return t.val
}

func (l *Lex) ScanToken() Token {
	bI := l.scanIdent() // Base Identifier. Identifier until proven otherwise
	return toValidToken(bI)
}
