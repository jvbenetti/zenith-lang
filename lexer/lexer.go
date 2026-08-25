package lexer

import "github.com/jvbenetti/zenith-lang.git/token"

// Lexer save actually status
type Lexer struct {
	input        string
	position     int  // Actually position of string
	readPosition int  // Next read position (actually position + 1)
	ch           byte // Actually ch in analyzes
}
