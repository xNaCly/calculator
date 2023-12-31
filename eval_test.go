package main

import (
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// from https://github.com/xNaCly/statlib/blob/master/distributions/dist.go
func compareFloats(a float64, b float64) bool {
	if a == b {
		return true
	}
	p := math.Abs(math.Abs(a) - math.Abs(b))
	return p < 1e-6 && p > 0
}

func TestEval(t *testing.T) {
	tests := []struct {
		name string
		in   []Node
		out  float64
	}{
		{
			name: "-2",
			in: []Node{
				&Unary{
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: -2,
		},
		{
			name: "2+2",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_PLUS},
					left:  &Number{token: Token{Raw: "2"}},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 4,
		},
		{
			name: "2-2",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_MINUS},
					left:  &Number{token: Token{Raw: "2"}},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 0,
		},
		{
			name: "2*2",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_ASTERISK},
					left:  &Number{token: Token{Raw: "2"}},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 4,
		},
		{
			name: "2/2",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_SLASH},
					left:  &Number{token: Token{Raw: "2"}},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 1,
		},
		{
			name: "1.025*3+1",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_PLUS},
					left: &Binary{
						token: Token{Type: TOKEN_ASTERISK},
						left:  &Number{token: Token{Raw: "1.025"}},
						right: &Number{token: Token{Raw: "3"}},
					},
					right: &Number{token: Token{Raw: "1"}},
				},
			},
			out: 4.074999999999999,
		},
		{
			name: "2*2+2",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_PLUS},
					left: &Binary{
						token: Token{Type: TOKEN_ASTERISK},
						left:  &Number{token: Token{Raw: "2"}},
						right: &Number{token: Token{Raw: "2"}},
					},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 6,
		},
		{
			name: "2*2*2+2",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_PLUS},
					left: &Binary{
						token: Token{Type: TOKEN_ASTERISK},
						left: &Binary{
							token: Token{Type: TOKEN_ASTERISK},
							left:  &Number{token: Token{Raw: "2"}},
							right: &Number{token: Token{Raw: "2"}},
						},
						right: &Number{token: Token{Raw: "2"}},
					},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 10,
		},
		{
			name: "2*2*2*2*2*2+2",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_PLUS},
					left: &Binary{
						token: Token{Type: TOKEN_ASTERISK},
						left: &Binary{
							token: Token{Type: TOKEN_ASTERISK},
							left: &Binary{
								token: Token{Type: TOKEN_ASTERISK},
								left: &Binary{
									token: Token{Type: TOKEN_ASTERISK},
									left: &Binary{
										token: Token{Type: TOKEN_ASTERISK},
										left:  &Number{token: Token{Raw: "2"}},
										right: &Number{token: Token{Raw: "2"}},
									},
									right: &Number{token: Token{Raw: "2"}},
								},
								right: &Number{token: Token{Raw: "2"}},
							},
							right: &Number{token: Token{Raw: "2"}},
						},
						right: &Number{token: Token{Raw: "2"}},
					},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 66,
		},
		{
			name: "readme example",
			in: []Node{
				&Binary{
					token: Token{Type: TOKEN_PLUS},
					left: &Binary{
						token: Token{Type: TOKEN_ASTERISK},
						left:  &Number{token: Token{Raw: "2"}},
						right: &Number{token: Token{Raw: "1"}},
					},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 4,
		},
	}

	vm := Vm{trace: true}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vm.NewVmIn(Compile(test.in)).Execute()
			assert.Equal(t, test.out, vm.reg[0])
		})
	}
}

func TestCompile(t *testing.T) {
	tests := []struct {
		In  string
		Out []Operation
	}{
		{In: "2+1*1", Out: []Operation{
			{OP_LOAD, 2},
			{OP_STORE, 1},
			{OP_LOAD, 1},
			{OP_STORE, 2},
			{OP_LOAD, 1},
			{OP_MULTIPY, 2},
			{OP_ADD, 1},
		}},
	}
	for _, test := range tests {
		t.Run(test.In, func(t *testing.T) {
			token := NewLexer(strings.NewReader(test.In)).Lex()
			ast := NewParser(token).Parse()
			assert.EqualValues(t, test.Out, Compile(ast))
		})
	}
}
