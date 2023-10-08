package main

func Compile(n []Node) []Operation {
	r := make([]Operation, 0)
	for _, c := range n {
		r = append(r, c.Compile()...)
	}
	r = append(r, Operation{OP_END, 0})
	return r
}
