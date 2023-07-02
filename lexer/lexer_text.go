package lexer

import (
	"fmt"
	"language_x/token"
	"testing"
)

func TextCaseSymbols(t *testing.T) {
	str := "+={}()"
	input := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.ASSIGN, "="},

		{token.LBRACE, "{"},

		{token.RBRACE, "}"},

		{token.LPAREN, "("},

		{token.RPAREN, ")"},
	}

	l := New(str)

	for i,tt := range input{
		tok := l.NextToken()
		if(tok.Type != tt.expectedType){
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if(tok.Literal != tt.expectedLiteral){
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",i, tt.expectedLiteral, tok.Literal)
		}

		fmt.Printf( "%+v\n", tok )

	}

}
