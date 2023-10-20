package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		Name string
		In   string
		Out  []Token
	}{
		{
			Name: "empty input",
			In:   "",
			Out: []Token{
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "whitespace",
			In:   "\r\n\t             ",
			Out: []Token{
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "comment",
			In:   "# this is a comment",
			Out: []Token{
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "symbols",
			In:   "+-/*()",
			Out: []Token{
				{TOKEN_PLUS, "+"},
				{TOKEN_MINUS, "-"},
				{TOKEN_SLASH, "/"},
				{TOKEN_ASTERISK, "*"},
				{TOKEN_BRACE_LEFT, "("},
				{TOKEN_BRACE_RIGHT, ")"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "number",
			In:   "123",
			Out: []Token{
				{TOKEN_NUMBER, "123"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "number with underscore",
			In:   "10_000",
			Out: []Token{
				{TOKEN_NUMBER, "10_000"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "number with e",
			In:   "10e5",
			Out: []Token{
				{TOKEN_NUMBER, "10e5"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "number with .",
			In:   "0.005",
			Out: []Token{
				{TOKEN_NUMBER, "0.005"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			Name: "number with . and underscore",
			In:   "1_000_000.0_000_5",
			Out: []Token{
				{TOKEN_NUMBER, "1_000_000.0_000_5"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			In: "1_000_000.0_000_5+5",
			Out: []Token{
				{TOKEN_NUMBER, "1_000_000.0_000_5"},
				{TOKEN_PLUS, "+"},
				{TOKEN_NUMBER, "5"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			in := strings.NewReader(test.In)
			out := NewLexer(in).Lex()
			assert.EqualValues(t, test.Out, out)
		})
	}
}
