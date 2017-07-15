package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"
	// Operators
	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"
	// Delimters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAERN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	FUNCTION = "FUNCTION"
	LET      = "FUN"
)
