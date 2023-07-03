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
		l.currentCh = 0
	} else {
		l.currentCh = l.input[l.readPosition]
	}
	l.currentPosition = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	l.skipWhiteSpace();

	switch l.currentCh {
	    case '=':
			if(l.peekChar() == '='){
				tok.Type = token.EQ 
				tok.Literal = "=="
				l.readChar() 
				} else {
					tok = newToken(token.ASSIGN, l.currentCh)
				}
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
	    case '-':
		tok = newToken(token.MINUS, l.currentCh)
		case '!':
			if(l.peekChar() == '='){
				tok.Type = token.NOT_EQ
				tok.Literal = "!="
				l.readChar()
			} else {
				tok = newToken(token.BANG, l.currentCh)
			}
		case '/':
		tok = newToken(token.SLASH, l.currentCh)
		case '*':
		tok = newToken(token.ASTERISK, l.currentCh)
		case '<':
		tok = newToken(token.LT, l.currentCh)
		case '>':
		tok = newToken(token.GT, l.currentCh)
		case 0:
			tok.Type = token.EOF
			tok.Literal = ""
		default:
			if (isLetter(l.currentCh)){
				ident := l.readIdentifier()
				tok.Type = token.FindIdentTokenType(ident)
				tok.Literal = ident
				return tok
			} 
			if (isDigit(l.currentCh)) {
				number := l.readNumber()
				tok.Type = token.INT
				tok.Literal = number
				return tok

			} else {
				tok = newToken(token.ILLEGAL, l.currentCh)
			}
}

l.readChar();

return tok

}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_')
}
func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9')
}

func (l *Lexer) readNumber() string {
	startPos := l.currentPosition
    for(isDigit(l.currentCh)){
		l.readChar()
	}
	return l.input[startPos:l.currentPosition]
}

func (l *Lexer) readIdentifier() string {
	str := []byte{}
	for(isLetter(l.currentCh)){
	str = append(str, l.currentCh) 
	l.readChar();
	}
	return string(str);
}

func (l *Lexer) peekChar() byte {
	if(l.readPosition >= len(l.input)){
		return '0';
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhiteSpace (){
	for(l.currentCh == ' ' || l.currentCh == '\n' || l.currentCh == '\t' || l.currentCh == '\r'){
		l.readChar()
	}
} 



func newToken(tokenType token.TokenType, ch byte) token.Token{
return token.Token{Type: tokenType, Literal: string(ch)};
}