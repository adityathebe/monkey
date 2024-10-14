package token

type TokenType string

const (
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"

	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// Operators
	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

var tokenmap = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}

func LookupIdent(ident string) TokenType {
	if v, ok := tokenmap[ident]; ok {
		return v
	}

	return IDENT
}
