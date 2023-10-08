package main

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		in  []Node
		out float64
	}{
		{in: []Node{
			&Addition{
				left:  &Number{token: Token{Raw: "2"}},
				right: &Number{token: Token{Raw: "2"}},
			},
		},
			out: 4,
		},
	}

	vm := Vm{trace: true}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.in), func(t *testing.T) {
			vm.NewVmIn(Compile(test.in)).Execute()
			if vm.reg[0] != test.out {
				t.Errorf("execution did not yield the correct result, wanted %f, got %f\n", test.out, vm.reg[0])
			}
		})
	}
}
