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

type Binary struct {
	token Token
	left  Node
	right Node
}

func (b *Binary) Compile() []Operation {
	codes := b.left.Compile()
	i := Allocator.alloc()
	defer Allocator.dealloc(i)
	codes = append(codes, Operation{OP_STORE, i})
	codes = append(codes, b.right.Compile()...)

	operation := OP_NOP
	switch b.token.Type {
	case TOKEN_PLUS:
		operation = OP_ADD
	case TOKEN_MINUS:
		operation = OP_SUBTRACT
	case TOKEN_SLASH:
		operation = OP_DIVIDE
	case TOKEN_ASTERISK:
		operation = OP_MULTIPY
	default:
		panic("Unknown type")
	}

	codes = append(codes, Operation{operation, i})
	return codes
}

func (b *Binary) String(ident int) string {
	identStr := strings.Repeat(" ", ident)
	return fmt.Sprint(identStr, b.token.Raw, "\n ", identStr, b.left.String(ident+1), "\n ", identStr, b.right.String(ident+1))
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
