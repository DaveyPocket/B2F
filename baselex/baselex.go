//	baselex is a lexical analyzer for the BASIC programming language
package baselex

import (
	"bufio"
	"strings"
)

//	buf defines the input buffer
var buf *bufio.Reader

func SetReader(b *bufio.Reader) {
	buf = b
}

//	TokName represents the name of a token
type TokName int

const (
	IDENTIFIER TokName = iota

	//	Keywords
	LET
	FOR
	PRINT
	TO
	IF
	THEN
	ELSE
	GOTO
	GOSUB
	//	Delimiting keywords
	NEXT
	ENDIF
	END
	RETURN

	//	Operators
	EQU
	ADD
	SUB
	MUL
	DIV
	MOD
	GREATERTHAN
	LESSTHAN
	GREATERTHANEQU
	LESSTHANEQU

	//	Expressives
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

	//	For tests
	THISGOESLAST
)

var reservedTok = map[string]TokName{
	"LET":    LET,
	"FOR":    FOR,
	"PRINT":  PRINT,
	"NEXT":   NEXT,
	"RETURN": RETURN,
	"TO":     TO,
	"=":      EQU,
	"+":      ADD,
	"-":      SUB,
	"*":      MUL,
	"/":      DIV,
	"\"":     QUOTE,
	" ":      WHITESPACE,
	"END":    END,
	"ENDIF":  ENDIF,
	"IF":     IF,
	"ELSE":   ELSE,
	"THEN":   THEN,
	"GOTO":   GOTO,
	"GOSUB":  GOSUB,
	"%":      MOD,
	"\n":     NEWLINE,
	":":      COLON,
	"":       EOF,
	"(":      OPENPAREN,
	")":      CLOSEPAREN,
	"[":      OPENSQUAREBRACKET,
	"]":      CLOSESQUAREBRACKET,
	"{":      OPENBRACKET,
	"}":      CLOSEBRACKET,
	">":		GREATERTHAN,
	"<":		LESSTHAN,
	">=":		GREATERTHANEQU,
	"<=":		LESSTHANEQU,
	"⾨":      THISGOESLAST,
}

//	StringToName converts a string to a token name if the string matches the name of a token.
func StringToName(str string) TokName {
	return reservedTok[str]
}

//	Token is a struct that defines a lexical token.
type Token struct {
	name   TokName
	lexeme string
}

//	GetName returns the name of the token.
func (t Token) GetName() TokName {
	return t.name
}

//	GetVal returns the value of the token.
func (t Token) GetLexeme() string {
	return t.lexeme
}

func (t Token) IsEOF() bool {
	return t.name == EOF && t.lexeme == ""
}

// readNextChar reads a rune from the buffer and returns said rune.
//	If the end of the buffer as been reached, return EOF.
func readNextRune() rune {
	r, _, err := buf.ReadRune()
	if err != nil {
		r = rune(0) //	EOF case
	}
	return r
}

func putRuneBack() {
	buf.UnreadRune()
}

//	isWhiteSpace returns true if the rune argument is a whitespace or
// horizontal tab.
func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isParen(r rune) bool {
	return r == '(' || r == ')'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

//	TODO - Remove isEOF or only check rune type
func isEOF(tk Token) bool {
	return tk.GetName() == EOF
}

func isSpecial(r rune) bool {
	return (r >= '!' && r <= '/') || (r >= ':' && r <= '@')
}

//	readSpace reads a single whitespace - freeing it from the buffer.
//	Frees rune from buffer until non-whitespace character found.
func readSpace() bool {
	r := readNextRune() //	Read a rune.
	if !isWhitespace(r) {
		putRuneBack() //	Put character back if it is not a whitespace.
		return false  //	Thing just read is not a space.
	}
	return true
}

// readSpaces calls readSpace repeatedly until a non-whitespace is found.
//	Useful for freeing tabs and spaces used for formatting program.
func clearSpaces() {
	for readSpace() {
	} //	Eat spaces and tabs until character.
}

//	readUntilEnd reads a sequence of characters recursively until a whitespace character is found.
//	TODO - Rename to getNextLexeme
func getNextLexeme(m []rune) string {
	// Read space must go in here
	r := readNextRune()

	// Code until * can be made compact and not so if-y
	if len(m) == 0 && r == '\n' {
		return string(r)
	}

	if (isLetter(r) || isNumber(r)) && r != '\n' && r != rune(0) {
		m = append(m, r)
		return getNextLexeme(m)
	}

	if isSpecial(r) && len(m) == 0 {
		return string(r)
	}
	//*
	putRuneBack()
	return string(m)
}

//	readLiteral returns a whitespace delimited string of text.
//	Intended for placing values into lexical tokens.
func readLexeme() string {
	clearSpaces() //	Clear preceding whitespaces
	return getNextLexeme([]rune{})
}

//	scanIdent returns an identifier token with value assigned to string literal read
// from buffer.
func readIdentifier() Token {
	return Token{name: IDENTIFIER, lexeme: readLexeme()}
}

func toValidToken(tk Token) Token {
	return Token{name: reservedTok[strings.ToUpper(tk.GetLexeme())], lexeme: tk.GetLexeme()}
}

//	ReadToken grabs the next lexical token from the buffer.
//	It first assumes that the token is an identifier - if the identifier
// just so happens to be another type, it changes it accordingly.
func ReadToken() Token {
	return toValidToken(readIdentifier()) // Base Identifier. Identifier until proven otherwise)
}
