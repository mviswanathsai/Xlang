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

	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"
	GT       = ">"
	LT       = "<"
	MINUS    = "-"
	PLUS     = "+"
	ASSIGN   = "="
	EQ       = "=="
	NOT_EQ   = "!="

	COMMA     = ","
	SEMICOLON = ";"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func FindIdentTokenType(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}