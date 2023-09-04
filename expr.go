package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node interface {
	Eval() float64
	String(ident int) string
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

func (n *Number) String(ident int) string {
	return fmt.Sprint(strings.Repeat(" ", ident), n.token.Raw)
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

func (a *Addition) String(ident int) string {
	identStr := strings.Repeat(" ", ident)
	return fmt.Sprint(identStr, "+\n ", identStr, a.left.String(ident+1), "\n ", identStr, a.right.String(ident+1))
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

func (s *Subtraction) String(ident int) string {
	identStr := strings.Repeat(" ", ident)
	return fmt.Sprint(identStr, "-\n ", identStr, s.left.String(ident+1), "\n ", identStr, s.right.String(ident+1))
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
func (m *Multiplication) String(ident int) string {
	identStr := strings.Repeat(" ", ident)
	return fmt.Sprint(identStr, "*\n ", identStr, m.left.String(ident+1), "\n ", identStr, m.right.String(ident+1))
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

func (d *Division) String(ident int) string {
	identStr := strings.Repeat(" ", ident)
	return fmt.Sprint(identStr, "/\n ", identStr, d.left.String(ident+1), "\n ", identStr, d.right.String(ident+1))
}
