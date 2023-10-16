package main

import (
	"strings"
	"testing"
)

func compareSlices[T comparable](t *testing.T, a []T, b []T) {
	if len(a) != len(b) {
		goto err
	}
	for i, e := range a {
		if e != b[i] {
			goto err
		}
	}
	return
err: // i know we are all goto haters, but i cant be assed to write this 3 times for tests :^)
	t.Errorf("%v != %v", a, b)
}

func TestLexer(t *testing.T) {
	tests := []struct {
		In  string
		Out []Token
	}{
		{
			In: "+-/*()",
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
			In: "",
			Out: []Token{
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			In: "123",
			Out: []Token{
				{TOKEN_NUMBER, "123"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			In: "10_000",
			Out: []Token{
				{TOKEN_NUMBER, "10_000"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			In: "10e5",
			Out: []Token{
				{TOKEN_NUMBER, "10e5"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			In: "0.005",
			Out: []Token{
				{TOKEN_NUMBER, "0.005"},
				{TOKEN_EOF, "TOKEN_EOF"},
			},
		},
		{
			In: "1_000_000.0_000_5",
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
		t.Run(test.In, func(t *testing.T) {
			in := strings.NewReader(test.In)
			out := NewLexer(in).Lex()
			compareSlices(t, out, test.Out)
		})
	}
}
