package main

import "strconv"

type Node interface {
	Eval() float64
}

type Number struct {
	token Token
}

func (n *Number) Eval() float64 {
	val, err := strconv.ParseFloat(n.token.Raw, 64)
	if err != nil {
		panic(err)
	}
	return val
}

type Addition struct {
	token Token
	left  Node
	right Node
}

func (a *Addition) Eval() float64 {
	left := a.left.Eval()
	right := a.right.Eval()
	return left + right
}

type Subtraction struct {
	token Token
	left  Node
	right Node
}

func (s *Subtraction) Eval() float64 {
	left := s.left.Eval()
	right := s.right.Eval()
	return left - right
}

type Multiplication struct {
	token Token
	left  Node
	right Node
}

func (m *Multiplication) Eval() float64 {
	left := m.left.Eval()
	right := m.right.Eval()
	return left * right
}

type Division struct {
	token Token
	left  Node
	right Node
}

func (d *Division) Eval() float64 {
	left := d.left.Eval()
	right := d.right.Eval()
	return left / right
}
