package main

import "fmt"

func Eval(n []Node) {
	for _, c := range n {
		fmt.Println(c.Eval())
	}
}
