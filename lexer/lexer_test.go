package lexer

import (
	"testing"

	"github.com/petrostrak/the-monkey-programming-language/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	testCases := []struct {
		expectedType    token.TokenType
		expextedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range testCases {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong; expected: %q, got: %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expextedLiteral {
			t.Fatalf("test[%d] - literan wrong; expected: %q, got: %q", i, tt.expextedLiteral, tok.Literal)
		}
	}
}
