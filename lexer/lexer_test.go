package lexer

import (
	"github.com/jvbenetti/zenith-lang.git/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		let x int64 = 10000;
		= + - ( ) { }
		10 == 10;
	`

	// Wait that Lexer can read the tokens
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.TYPE_INT64, "int64"},
		{token.ASSIGN, "="},
		{token.INT, "10000"},
		{token.SEMICOLON, ";"},

		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		{token.INT, "10"},
		{token.EQUAL, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	// Init Lexer with test string
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Testes[%d] - Tipo de token incorreto. Esperado: %q, Recebido: %q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Testes[%d] - Literal incorreto. Esperado: %q, Recebido: %q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
