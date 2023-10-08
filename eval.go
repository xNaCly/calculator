package main

func Compile(n []Node) []Operation {
	r := make([]Operation, 0)
	for _, c := range n {
		r = append(r, c.Compile()...)
	}
	return r
}
