package lexer

import (
	"fmt"

	"github.com/akash1047/crepel/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte

	errors []string
}

func (l *Lexer) LastError() string {
	return l.errors[len(l.errors)-1]
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() (token.Token, bool) {
	ok := true
	var tok token.Token

	switch l.ch {
	case 0:
		return token.Token{Type: token.EOF, Literal: ""}, true
	default:
		tok.Type = token.ILLEGAL
		tok.Literal = string(l.ch)

		msg := fmt.Sprintf("stray '%c' in program", l.ch)
		l.errors = append(l.errors, msg)
		ok = false
	}

	l.readChar()

	return tok, ok
}
