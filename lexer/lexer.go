package lexer

import (
	"language_x/token"
)

type Lexer struct {
	input           string
	currentPosition int
	readPosition    int
	currentCh       byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) {
		l.readPosition = 0
	} else {
		l.currentCh = l.input[l.readPosition]
	}
	l.currentPosition = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	switch l.currentCh {
	    case '=':
		tok = newToken(token.ASSIGN, l.currentCh)
		case ';':
		tok = newToken(token.SEMICOLON, l.currentCh)
		case '(':
		tok = newToken(token.LPAREN, l.currentCh)
		case ')':
		tok = newToken(token.RPAREN, l.currentCh)
		case ',':
		tok = newToken(token.COMMA, l.currentCh)
		case '+':
		tok = newToken(token.PLUS, l.currentCh)
		case '{':
		tok = newToken(token.LBRACE, l.currentCh)
		case '}':
		tok = newToken(token.RBRACE, l.currentCh)
		case 0:
		tok.Literal = ""
		tok.Type = token.EOF
}


l.readChar();

return tok

}

func newToken(tokenType token.TokenType, ch byte) token.Token{
return token.Token{Type: tokenType, Literal: string(ch)};
}