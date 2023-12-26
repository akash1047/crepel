package lexer

import (
	"fmt"
	"strings"

	"github.com/akash1047/crepel/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte

	errors []Error
}

func (l *Lexer) LastError() Error {
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

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) getLine(index int) string {
	if l.input[index] == '\n' {
		return ""
	}

	start, end := 0, len(l.input)

	for i := index; i > 0; i-- {
		if l.input[i] == '\n' {
			start = i + 1
			break
		}
	}

	// find end

	i := strings.IndexByte(l.input[index:], '\n')

	if i != -1 {
		end = 1
	}

	return l.input[start:end]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() (token.Token, bool) {
	l.skipWhitespace()

	ok := true
	var tok token.Token

	switch l.ch {
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)

	case '/':
		switch l.peekChar() {
		case '/':
			l.skipSinglellineComment()
			return l.NextToken()

		case '*':
			if l.skipMultilineComment() {
				return l.NextToken()
			} else {
				return token.Token{Type: token.ILLEGAL, Literal: "/*"}, false
			}

		case '=':
			ch := l.ch
			l.readChar()
			tok.Type = token.DIV_ASSIGN
			tok.Literal = string(ch) + string(l.ch)

		default:
			tok = newToken(token.SLASH, l.ch)
		}

	case 0:
		return token.Token{Type: token.EOF, Literal: ""}, true
	default:

		tok = newToken(token.ILLEGAL, l.ch)

		// ‘@’
		err := Error{
			Message: fmt.Sprintf("stray ‘%c’ in program", l.ch),
			Line:    l.getLine(l.position),
			Span:    [2]int{l.position, l.position + 1},
		}

		l.errors = append(l.errors, err)
		ok = false
	}

	l.readChar()

	return tok, ok
}

func (l *Lexer) skipSinglellineComment() {
	// l.ch = /
	// l.peekChar() = /

	l.readChar()
	l.readChar()

	for l.ch != 0 {
		if l.ch == '\n' {
			l.readChar()
			break
		}

		l.readChar()
	}
}

func (l *Lexer) skipMultilineComment() bool {
	position := l.position

	// l.ch = /
	// l.peekChar() = *

	l.readChar()
	l.readChar()

	for {
		switch l.ch {
		case '*':
			l.readChar()
			if l.ch == '/' {
				l.readChar()
				return true
			}
		case 0: // error unterminated comment
			err := Error{
				Message: "unterminated comment",
				Line:    l.getLine(position),
				Span:    [2]int{position, position + 2},
			}

			l.errors = append(l.errors, err)
			return false
		}

		l.readChar()
	}
}
