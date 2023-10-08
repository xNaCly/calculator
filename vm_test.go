package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestVm(t *testing.T) {
	tests := []struct {
		name string
		ops  []Operation
		exp  float64
	}{
		{
			name: "nop",
			ops: []Operation{
				{OP_NOP, 0},
				{OP_END, 0},
			},
			exp: 0,
		},
		{
			name: "addition",
			ops: []Operation{
				{OP_LOAD, 51},
				{OP_STORE, 1},
				{OP_LOAD, 12},
				{OP_ADD, 1},
				{OP_END, 0},
			},
			exp: 63,
		},
		{
			name: "subtraction",
			ops: []Operation{
				{OP_LOAD, 9},
				{OP_STORE, 1},
				{OP_LOAD, 5},
				{OP_SUBTRACT, 1},
				{OP_END, 0},
			},
			exp: -4,
		},
		{
			name: "multiplication",
			ops: []Operation{
				{OP_LOAD, 12},
				{OP_STORE, 1},
				{OP_LOAD, 12},
				{OP_MULTIPY, 1},
				{OP_END, 0},
			},
			exp: 144,
		},
		{
			name: "division",
			ops: []Operation{
				{OP_LOAD, 2},
				{OP_STORE, 1},
				{OP_LOAD, 25},
				{OP_DIVIDE, 1},
				{OP_END, 0},
			},
			exp: 12.5,
		},
	}
	v := Vm{trace: true}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v.NewVmIn(test.ops)
			v.Execute()
			if v.reg[0] != test.exp {
				in := make([]string, len(test.ops))
				for i, o := range test.ops {
					in[i] = fmt.Sprintf("%s:%f", OP_LOOKUP[o.Code], o.Arg)
				}
				t.Errorf("execution did not yield the correct result, wanted %f, got %f, for: \n%s\n", test.exp, v.reg[0], strings.Join(in, "\n"))
			}
		})
	}
}
