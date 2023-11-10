package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type RegisterAllocator struct {
	registers [REGISTER_COUNT]bool
}

func (r *RegisterAllocator) alloc() float64 {
	for i, v := range r.registers {
		if v == false {
			r.registers[i] = true
			return float64(i + 1)
		}
	}
	panic("Out of bounds, no more free registers")
}

func (r *RegisterAllocator) dealloc(index float64) {
	i := int(index)
	if r.registers[i-1] {
		r.registers[i-1] = false
	} else {
		panic("Register not previously occupied")
	}
}

var Allocator RegisterAllocator

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
	codes := a.left.Compile()
	i := Allocator.alloc()
	defer Allocator.dealloc(i)
	codes = append(codes, Operation{OP_STORE, i})
	codes = append(codes, a.right.Compile()...)
	codes = append(codes, Operation{OP_ADD, i})
	return codes
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
	codes := s.left.Compile()
	i := Allocator.alloc()
	defer Allocator.dealloc(i)
	codes = append(codes, Operation{OP_STORE, i})
	codes = append(codes, s.right.Compile()...)
	codes = append(codes, Operation{OP_SUBTRACT, i})
	return codes
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
	codes := m.left.Compile()
	i := Allocator.alloc()
	defer Allocator.dealloc(i)
	codes = append(codes, Operation{OP_STORE, i})
	codes = append(codes, m.right.Compile()...)
	codes = append(codes, Operation{OP_MULTIPY, i})
	return codes
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
	codes := d.left.Compile()
	i := Allocator.alloc()
	defer Allocator.dealloc(i)
	codes = append(codes, Operation{OP_STORE, i})
	codes = append(codes, d.right.Compile()...)
	codes = append(codes, Operation{OP_DIVIDE, i})
	return codes
}

func (d *Division) String(ident int) string {
	identStr := strings.Repeat(" ", ident)
	return fmt.Sprint(identStr, "/\n ", identStr, d.left.String(ident+1), "\n ", identStr, d.right.String(ident+1))
}

type Unary struct {
	token Token
	right Node
}

func (u *Unary) Compile() []Operation {
	codes := u.right.Compile()
	codes = append(codes, Operation{Code: OP_NEG})
	return codes
}

func (u *Unary) String(ident int) string {
	identStr := strings.Repeat(" ", ident)
	return fmt.Sprint(identStr, "/\n ", identStr, "-", u.right.String(ident+1))
}
