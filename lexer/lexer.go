package lexer

import (
	"language_x/token"
	"unicode"
)

type Lexer struct {
	input           []rune
	currentPosition int
	readPosition    int
	currentRune       rune
}

func New(input string) *Lexer {
	runeInput := []rune(input)
	l := &Lexer{input: runeInput}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	
	if l.readPosition >= len(l.input) {
		l.currentRune = 0
	} else {
		l.currentRune = l.input[l.readPosition]
	}
	l.currentPosition = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token

	l.skipWhiteSpace();

	switch l.currentRune {
	    case '=':
			if(l.peekChar() == '='){
				tok.Type = token.EQ 
				tok.Literal = []rune("==")
				l.readChar() 
				} else {
					tok = newToken(token.ASSIGN, l.currentRune)
				}
		case ';':
		tok = newToken(token.SEMICOLON, l.currentRune)
		case '(':
		tok = newToken(token.LPAREN, l.currentRune)
		case ')':
		tok = newToken(token.RPAREN, l.currentRune)
		case ',':
		tok = newToken(token.COMMA, l.currentRune)
		case '+':
		tok = newToken(token.PLUS, l.currentRune)
		case '{':
		tok = newToken(token.LBRACE, l.currentRune)
		case '}':
		tok = newToken(token.RBRACE, l.currentRune)
	    case '-':
		tok = newToken(token.MINUS, l.currentRune)
		case '!':
			if(l.peekChar() == '='){
				tok.Type = token.NOT_EQ
				tok.Literal = []rune("!=")
				l.readChar()
			} else {
				tok = newToken(token.BANG, l.currentRune)
			}
		case '/':
		tok = newToken(token.SLASH, l.currentRune)
		case '*':
		tok = newToken(token.ASTERISK, l.currentRune)
		case '<':
		tok = newToken(token.LT, l.currentRune)
		case '>':
		tok = newToken(token.GT, l.currentRune)
		case 0:
			tok.Type = token.EOF
			tok.Literal = []rune("")
		default:
			if (isLetter(l.currentRune)){
				ident := l.readIdentifier()
				tok.Type = token.FindIdentTokenType(ident)
				tok.Literal = ident
				return tok
			} 
			if (isDigit(l.currentRune)) {
				number := l.readNumber()
				tok.Type = token.INT
				tok.Literal = number
				return tok

			} else {
				tok = newToken(token.ILLEGAL, l.currentRune)
			}
}

l.readChar();

return tok

}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || (ch == '_')
}
func isDigit(ch rune) bool {
	return ('0' <= ch && ch <= '9')
}

func (l *Lexer) readNumber() []rune {
	startPos := l.currentPosition
    for(isDigit(l.currentRune)){
		l.readChar()
	}
	return l.input[startPos:l.currentPosition]
}

func (l *Lexer) readIdentifier() []rune {
	str := []rune{}
	for(isLetter(l.currentRune)){
	str = append(str, l.currentRune) 
	l.readChar();
	}
	return str;
}

func (l *Lexer) peekChar() rune {
	if(l.readPosition >= len(l.input)){
		return '0';
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhiteSpace (){
	for(l.currentRune == ' ' || l.currentRune == '\n' || l.currentRune == '\t' || l.currentRune == '\r'){
		l.readChar()
	}
} 



func newToken(tokenType token.TokenType, ch rune) token.Token{
return token.Token{Type: tokenType, Literal: []rune{ch}};
}