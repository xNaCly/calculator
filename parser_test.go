package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	tests := []struct {
		Name string
		In   []Token
		Out  []Node
	}{
		{
			Name: "Empty input",
			In:   []Token{{TOKEN_EOF, "TOKEN_EOF"}},
			Out:  []Node{},
		},
		{
			Name: "Addition",
			In: []Token{
				{TOKEN_NUMBER, "5"},
				{TOKEN_PLUS, "+"},
				{TOKEN_NUMBER, "1"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
			Out: []Node{&Binary{
				token: Token{TOKEN_PLUS, "+"},
				left:  &Number{token: Token{TOKEN_NUMBER, "5"}},
				right: &Number{token: Token{TOKEN_NUMBER, "1"}},
			}},
		},
		{
			Name: "Subtraction",
			In: []Token{
				{TOKEN_NUMBER, "5"},
				{TOKEN_MINUS, "-"},
				{TOKEN_NUMBER, "1"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
			Out: []Node{&Binary{
				token: Token{TOKEN_MINUS, "-"},
				left:  &Number{token: Token{TOKEN_NUMBER, "5"}},
				right: &Number{token: Token{TOKEN_NUMBER, "1"}},
			}},
		},
		{
			Name: "Multiplication",
			In: []Token{
				{TOKEN_NUMBER, "5"},
				{TOKEN_ASTERISK, "*"},
				{TOKEN_NUMBER, "1"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
			Out: []Node{&Binary{
				token: Token{TOKEN_ASTERISK, "*"},
				left:  &Number{token: Token{TOKEN_NUMBER, "5"}},
				right: &Number{token: Token{TOKEN_NUMBER, "1"}},
			}},
		},
		{
			Name: "Division",
			In: []Token{
				{TOKEN_NUMBER, "5"},
				{TOKEN_SLASH, "/"},
				{TOKEN_NUMBER, "1"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
			Out: []Node{&Binary{
				token: Token{TOKEN_SLASH, "/"},
				left:  &Number{token: Token{TOKEN_NUMBER, "5"}},
				right: &Number{token: Token{TOKEN_NUMBER, "1"}},
			}},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := NewParser(test.In).Parse()
			assert.EqualValues(t, test.Out, out)
		})
	}
}
