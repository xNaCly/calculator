package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

const (
	TOKEN_UNKNOWN = iota + 1

	TOKEN_NUMBER
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_ASTERISK
	TOKEN_SLASH

	TOKEN_BRACE_LEFT
	TOKEN_BRACE_RIGHT

	TOKEN_EOF
)

// for debugging
var TOKEN_LOOKUP = map[int]string{
	TOKEN_UNKNOWN:     "UNKNOWN",
	TOKEN_NUMBER:      "TOKEN_NUMBER",
	TOKEN_PLUS:        "TOKEN_PLUS",
	TOKEN_MINUS:       "TOKEN_MINUS",
	TOKEN_ASTERISK:    "TOKEN_ASTERISK",
	TOKEN_SLASH:       "TOKEN_SLASH",
	TOKEN_BRACE_LEFT:  "TOKEN_BRACE_LEFT",
	TOKEN_BRACE_RIGHT: "TOKEN_BRACE_RIGHT",
	TOKEN_EOF:         "EOF",
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
		ttype := TOKEN_UNKNOWN

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
			ttype = TOKEN_PLUS
		case '-':
			ttype = TOKEN_MINUS
		case '/':
			ttype = TOKEN_SLASH
		case '*':
			ttype = TOKEN_ASTERISK
		case '(':
			ttype = TOKEN_BRACE_LEFT
		case ')':
			ttype = TOKEN_BRACE_RIGHT
		default:
			if (l.cur > '0' && l.cur < '9') || l.cur == '.' {
				t = append(t, l.number())
				continue
			}
		}

		if ttype != TOKEN_UNKNOWN {
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
		Type: TOKEN_EOF,
		Raw:  "TOKEN_EOF",
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
		Type: TOKEN_NUMBER,
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
