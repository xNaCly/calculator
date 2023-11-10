package main

func Compile(n []Node) []Operation {
	o := make([]Operation, 0)
	for _, node := range n {
		o = append(o, node.Compile()...)
	}
	return o
}
