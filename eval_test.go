package main

import (
	"math"
	"testing"
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
			name: "1.025*3+1",
			in: []Node{
				&Addition{
					left: &Multiplication{
						left:  &Number{token: Token{Raw: "1.025"}},
						right: &Number{token: Token{Raw: "3"}},
					},
					right: &Number{token: Token{Raw: "1"}},
				},
			},
			out: 4.075,
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
		{
			name: "readme example",
			in: []Node{
				&Addition{
					left: &Multiplication{
						left:  &Number{token: Token{Raw: "2"}},
						right: &Number{token: Token{Raw: "1"}},
					},
					right: &Number{token: Token{Raw: "2"}},
				},
			},
			out: 4,
		},
	}

	vm := Vm{trace: false}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vm.NewVmIn(Compile(test.in)).Execute()
			if !compareFloats(vm.reg[0], test.out) {
				t.Errorf("execution did not yield the correct result, wanted %.10f, got %.10f\n", test.out, vm.reg[0])
			}
		})
	}
}
