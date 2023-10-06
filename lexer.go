package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

const (
	UNKNOWN = iota + 1

	NUMBER
	PLUS
	MINUS
	ASTERISK
	SLASH

	BRACE_LEFT
	BRACE_RIGHT

	EOF
)

// for debugging
var TOKEN_LOOKUP = map[int]string{
	UNKNOWN: "UNKNOWN",

	NUMBER:   "NUMBER",
	PLUS:     "PLUS",
	MINUS:    "MINUS",
	ASTERISK: "ASTERISK",
	SLASH:    "SLASH",

	BRACE_LEFT:  "BRACE_LEFT",
	BRACE_RIGHT: "BRACE_RIGHT",
	EOF:         "EOF",
}

type Token struct {
	Type int
	Raw  string
}

type Lexer struct {
	scanner bufio.Reader
	cur     rune
}

func NewLexer(reader io.Reader) *Lexer {
	l := &Lexer{
		scanner: *bufio.NewReader(reader),
	}
	l.advance()
	return l
}

// transform list of characters to list of tokens
func (l *Lexer) Lex() []Token {
	t := make([]Token, 0)
	for l.cur != 0 {
		ttype := UNKNOWN

		switch l.cur {
		case '#':
			for l.cur != '\n' {
				l.advance()
			}
			continue
		case ' ', '\n', '\t', '\r':
			l.advance()
			continue
		case '+':
			ttype = PLUS
		case '-':
			ttype = MINUS
		case '/':
			ttype = SLASH
		case '*':
			ttype = ASTERISK
		case '(':
			ttype = BRACE_LEFT
		case ')':
			ttype = BRACE_RIGHT
		default:
			if (l.cur > '0' && l.cur < '9') || l.cur == '.' {
				t = append(t, l.number())
				continue
			}
		}

		if ttype != UNKNOWN {
			t = append(t, Token{
				Type: ttype,
				Raw:  string(l.cur),
			})
		} else {
			log.Fatalf("unknown %q in input", l.cur)
		}

		l.advance()
	}
	t = append(t, Token{
		Type: EOF,
		Raw:  "EOF",
	})
	return t
}

// advances until cur char is no longer [0-9\._e], returns token with list of matching chars
func (l *Lexer) number() Token {
	b := strings.Builder{}
	for (l.cur > '0' && l.cur < '9') || l.cur == '.' || l.cur == '_' || l.cur == 'e' {
		b.WriteRune(l.cur)
		l.advance()
	}
	return Token{
		Raw:  b.String(),
		Type: NUMBER,
	}
}

// advance to the next character
func (l *Lexer) advance() {
	r, _, err := l.scanner.ReadRune()
	if err != nil {
		l.cur = 0
	} else {
		l.cur = r
	}
}
