/*
	This package performs Lexical Analysis for BASIC
*/
package baselex

import (
	"bufio"
	"strings"
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
	IF
	THEN
	ELSE
	ENDIF
	END
	GOTO
	GOSUB

	//	Operators
	EQU
	ADD
	SUB
	MUL
	DIV
	MOD

	//	Expressives
	LABEL
	LNUM
	QUOTE
	WHITESPACE
	NEWLINE
	COLON
	EOF

	// Things
	OPENPAREN
	CLOSEPAREN
	OPENSQUAREBRACKET
	CLOSESQUAREBRACKET
	OPENBRACKET
	CLOSEBRACKET
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
	"ENDIF":			ENDIF,
	"IF":				IF,
	"ELSE":			ELSE,
	"THEN":			THEN,
	"GOTO":			GOTO,
	"GOSUB":			GOSUB,
	"%":			MOD,
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
	return Token{name: reservedTok[strings.ToUpper(tk.val)], val: tk.val}
}

// read reads a rune from the buffer and returns said rune.
//	If the end of the buffer as been reached, return EOF.
func (l *Lex) read() (rune) {
	r, _, err := l.b.ReadRune()
	if err != nil {
		r = rune(0)
	}
	return r
}

//	isWhiteSpace returns true if the rune argument is a whitespace or
// horizontal tab.
func isWhitespace(r rune) (bool) {
	return r == ' ' || r == '\t'
}

//	readSpace reads a single whitespace - freeing it from the buffer.
//	Frees rune from buffer until non-whitespace character found.
func (l *Lex) readSpace() (bool) {
	r := l.read()				//	Read a rune.
	if !isWhitespace(r) {
		l.b.UnreadRune()		//	Put character back if it is not a whitespace.
		return false			//	Thing just read is not a space.
	}

	return true
}

// readSpaces calls readSpace repeatedly until a non-whitespace is found.
//	Useful for freeing tabs and spaces used for formatting program.
func (l *Lex) readSpaces() {
	for l.readSpace() {}		//	Eat spaces and tabs until character.
}

//	readUntilEnd reads a sequence of characters recursively until a whitespace character is found.
func (l *Lex) readUntilEnd(m []rune) string {
	// Read space must go in here
	r := l.read()

	// Code until * can be made compact and not so if-y
	if len(m) == 0 && r == '\n' {
		return string(r)
	}

	// isWhitespace
	if !isWhitespace(r) && r != '\n' && r != rune(0) {
		m = append(m, r)
		return l.readUntilEnd(m)
	}
	//*

	l.b.UnreadRune()
	return string(m)
}

//	readLiteral returns a whitespace delimited string of text.
//	Intended for placing values into lexical tokens.
func (l *Lex) readLiteral() string {
	l.readSpaces()
	return l.readUntilEnd([]rune{})
}

//	scanIdent returns an identifier token with value assigned to string literal read
// from buffer.
func (l *Lex) scanIdent() Token {
	return Token{name: IDENTIFIER, val: l.readLiteral()}
}

//	GetName returns the name of the token.
func (t Token) GetName() (tokName) {
	return t.name
}

//	GetVal returns the value of the token.
func (t Token) GetVal() (string) {
	return t.val
}

//	ScanToken grabs the next lexical token from the buffer.
//	It first assumes that the token is an identifier - if the identifier
// just so happens to be another type, it changes it accordingly.
func (l *Lex) ScanToken() Token {
	bI := l.scanIdent() // Base Identifier. Identifier until proven otherwise
	return toValidToken(bI)
}
