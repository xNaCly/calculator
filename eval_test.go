package main

import (
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		name string
		in   []Node
		out  float64
	}{
		{
			name: "2+2",
			in: []Node{
				&Addition{
					left:  &Number{token: Token{Raw: "2"}},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 4,
		},
		{
			name: "2-2",
			in: []Node{
				&Subtraction{
					left:  &Number{token: Token{Raw: "2"}},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 0,
		},
		{
			name: "2*2",
			in: []Node{
				&Multiplication{
					left:  &Number{token: Token{Raw: "2"}},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 4,
		},
		{
			name: "2/2",
			in: []Node{
				&Division{
					left:  &Number{token: Token{Raw: "2"}},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 1,
		},
		{
			name: "2*2+2",
			in: []Node{
				&Addition{
					left: &Multiplication{
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
				&Addition{
					left: &Multiplication{
						left: &Multiplication{
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
				&Addition{
					left: &Multiplication{
						left: &Multiplication{
							left: &Multiplication{
								left: &Multiplication{
									left: &Multiplication{
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
	}

	vm := Vm{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vm.NewVmIn(Compile(test.in)).Execute()
			if vm.reg[0] != test.out {
				t.Errorf("execution did not yield the correct result, wanted %f, got %f\n", test.out, vm.reg[0])
			}
		})
	}
}
