package main

import (
	"strings"
	"testing"
)

func compareNodeArrays(t *testing.T, a []Node, b []Node) {
	if len(a) != len(b) {
		builder := strings.Builder{}
		builder.WriteString("\nast a:\n")
		for _, c := range a {
			builder.WriteString(c.String(0))
		}
		builder.WriteString("--------------")
		builder.WriteString("\nast b:\n")
		for _, c := range b {
			builder.WriteString(c.String(0))
		}
		t.Errorf("length of slices did not match: len(a)=%d != len(b)=%d for asts: %s", len(a), len(b), builder.String())

	}
	for i, e := range a {
		if e != b[i] {
			t.Errorf("%+v != %+v", e.String(0), b[i].String(0))
		}
	}
}

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
			Out: []Node{&Addition{
				token: Token{TOKEN_PLUS, "+"},
				left:  &Number{token: Token{TOKEN_NUMBER, "5"}},
				right: &Number{token: Token{TOKEN_NUMBER, "1"}},
			}},
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := NewParser(test.In).Parse()
			compareNodeArrays(t, out, test.Out)
		})
	}
}
