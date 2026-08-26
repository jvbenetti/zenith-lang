package lexer

import "github.com/jvbenetti/zenith-lang.git/token"

// Lexer save actually status
type Lexer struct {
	input        string
	position     int  // Actually position of string
	readPosition int  // Next read position (actually position + 1)
	ch           byte // Actually ch in analyzes
}

// New init the Lexer and read the first char
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}
