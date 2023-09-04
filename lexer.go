package main

import (
	"bufio"
	"io"
	"strings"
)

const (
	UNKNOWN = iota + 1
	LABEL
	NUMBER
	PLUS
	MINUS
	ASTERISK
	SLASH
	EOF
)

var TOKEN_LOOKUP = map[int]string{
	UNKNOWN:  "UNKNOWN",
	LABEL:    "LABEL",
	NUMBER:   "NUMBER",
	PLUS:     "PLUS",
	MINUS:    "MINUS",
	ASTERISK: "ASTERISK",
	SLASH:    "SLASH",
	EOF:      "EOF",
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
		}

		l.advance()
	}
	t = append(t, Token{
		Type: EOF,
		Raw:  "EOF",
	})
	return t
}

func (l *Lexer) number() Token {
	b := strings.Builder{}
	for (l.cur > '0' && l.cur < '9') || l.cur == '.' || l.cur == '_' {
		b.WriteRune(l.cur)
		l.advance()
	}
	return Token{
		Raw:  b.String(),
		Type: NUMBER,
	}
}

func (l *Lexer) advance() {
	r, _, err := l.scanner.ReadRune()
	if err != nil {
		l.cur = 0
	} else {
		l.cur = r
	}
}
