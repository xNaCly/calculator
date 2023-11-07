package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getRegister() float64 {
	t := CUR_REG
	CUR_REG = CUR_REG + 1
	return t
}

type Node interface {
	Compile() []Operation
	String(ident int) string
}

type Number struct {
	token Token
}

func (n *Number) Compile() []Operation {
	val, err := strconv.ParseFloat(n.token.Raw, 64)
	if err != nil {
		log.Panicf("compile: failed to parse float: %q", err)
	}
	return []Operation{{OP_LOAD, val}}
}

func (n *Number) String(ident int) string {
	return fmt.Sprint(strings.Repeat(" ", ident), n.token.Raw)
}

type Addition struct {
	token Token
	left  Node
	right Node
}

func (a *Addition) Compile() []Operation {
	op := a.left.Compile()
	r := getRegister()
	op = append(op, Operation{OP_STORE, r})
	op = append(op, a.right.Compile()...)
	op = append(op, Operation{OP_ADD, r})
	return op
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

func (s *Subtraction) Compile() []Operation {
	op := s.left.Compile()
	r := getRegister()
	op = append(op, Operation{OP_STORE, r})
	op = append(op, s.right.Compile()...)
	op = append(op, Operation{OP_SUBTRACT, r})
	return op
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

func (m *Multiplication) Compile() []Operation {
	op := m.left.Compile()
	r := getRegister()
	op = append(op, Operation{OP_STORE, r})
	op = append(op, m.right.Compile()...)
	op = append(op, Operation{OP_MULTIPY, r})
	return op
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

func (d *Division) Compile() []Operation {
	op := d.left.Compile()
	r := getRegister()
	op = append(op, Operation{OP_STORE, r})
	op = append(op, d.right.Compile()...)
	op = append(op, Operation{OP_DIVIDE, r})
	return op
}

func (d *Division) String(ident int) string {
	identStr := strings.Repeat(" ", ident)
	return fmt.Sprint(identStr, "/\n ", identStr, d.left.String(ident+1), "\n ", identStr, d.right.String(ident+1))
}
