package token

type TokenType string

// Token é a estrutura que o Lexer vai gerar.
// Ele guarda o tipo (ex: "+") e o valor literal exato que foi digitado (ex: "+").
type Token struct {
	Type    TokenType
	Literal string
}

// Dict to zenith can understand
const (
	ILLEGAL = "ILLEGAL" // simbols that the language dont understood
	EOF     = "EOF"     // End Of File

	// IDENT and INT
	IDENT = "IDENT" // Vars name: x, y, some, result
	INT   = "INT"   // Numbers: 1, 2, 100
	FLOAT = "FLOAT" // Floating-point numbers: 1.0, 2.5, 100.0

	// ASSIGN and PLUS Math operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MULT   = "*"
	DIV    = "/"
	Equal  = "=="

	// COMMA and SEMICOLON Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LSQUARE   = "["
	RSQUARE   = "]"

	// FUNCTION and LET Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
