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

// readChar next pointer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 (end of file)
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken analyzes actually char and return token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Jump whitespaces
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// Enter logic to read letters
		tok = newToken(token.ILLEGAL, l.ch)
	}
	l.readChar() // Go to next char before return token
	return tok
}

// newToken is a auxiliary func to create token faster
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// skipWhitespace ignore whitespaces
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
