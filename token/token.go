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

	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"

	PLUS   = "+"
	ASSIGN = "="

	COMMA     = ","
	SEMICOLON = ";"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func FindIdentTokenType(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}