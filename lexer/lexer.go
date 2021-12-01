package lexer

import "github.com/petrostrak/the-monkey-programming-language/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char).
	readPosition int  // readPosition always points to the next character in the input.
	ch           byte // current char under examination.
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar gives us the next character and advance our position in the input string.
func (l *Lexer) readChar() {

	// if we have reached the end of the input, we set l.ch to 0 (ASCII code for null).
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {

		// set l.ch to the next character
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken looks at the current character under examination (l.ch) and
// returns a token depending on which character it is.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	// before returning the token, we advance our pointers into the input
	// so when we call NextToken() again the l.ch field is already updated.
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
