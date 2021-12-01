package lexer

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
