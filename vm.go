package main

import (
	"fmt"
	"log"
)

// Contains the virtual machine used for executing the calculators byte code.
//
// Register overview:
//
//   - 0: reserved for results of operations, loading data
//   - 1-3: free for use
//
// Example:
//
//   - LOAD    51      ; loads the numeric value 51 into register 0
//   - STORE   1       ; moves the value stored in register 0 in register 1, sets register 0 to 0
//   - LOAD    12      ; loads the numeric value 12 into register 0
//   - ADD     1       ; adds register 0 and register 1 together, assigns result to register 0
//   - INSPECT 0       ; prints the register 0s value

// represents an operation the vm performs
type OpCode uint8

const (
	OP_NOP      OpCode = iota
	OP_LOAD            // loads the argument into register0
	OP_STORE           // stores the value of register0 in the specified register, set register0 to 0
	OP_ADD             // adds the value of register0 and the value of the specified register together, stores the result in register0
	OP_SUBTRACT        // subtracts the value of register0 from the value of the specified register, stores the result in register0
	OP_MULTIPY         // multiplies the value of register0 with the value of the specified register, stores the result in register0
	OP_DIVIDE          // divides the value of register0 by the value of the specified register, stores the result in register0
	OP_INSPECT         // prints the value of the given register
	OP_END             // end of input
)

var OP_LOOKUP = map[OpCode]string{
	OP_NOP:      "OP_NOP",
	OP_LOAD:     "OP_LOAD",
	OP_STORE:    "OP_STORE",
	OP_ADD:      "OP_ADD",
	OP_SUBTRACT: "OP_SUBTRACT",
	OP_MULTIPY:  "OP_MULTIPY",
	OP_DIVIDE:   "OP_DIVIDE",
	OP_INSPECT:  "OP_INSPECT",
	OP_END:      "OP_END",
}

// represents an operation and its argument
type Operation struct {
	Code OpCode  // type of operation
	Arg  float64 // operation argument
}

type Vm struct {
	reg   [4]float64  // registers
	in    []Operation // operations to execute
	pos   int         // current position in input
	trace bool        // prints every operation if enabled
}

func (vm *Vm) NewVmIn(in []Operation) *Vm {
	vm.pos = 0
	vm.in = in
	vm.reg = [4]float64{0, 0, 0, 0}
	return vm
}

// if next position in vm.in boundary, increment position
func (vm *Vm) advance() {
	if vm.pos+1 < len(vm.in) {
		vm.pos++
	}
}

// return current operation
func (vm *Vm) cur() Operation {
	return vm.in[vm.pos]
}

// checks if index is in register boundary, converts to int, returns
func (vm *Vm) regBoundCheck(index float64) int {
	i := int(index)
	if i < 0 || i > len(vm.reg) {
		log.Panicf("vm: Out of bounds register access for %d", i)
	}
	return i
}

func (vm *Vm) Execute() {
	for {
		cur := vm.cur()

		if vm.trace {
			fmt.Printf("vm: %10s :: %f\n", OP_LOOKUP[vm.cur().Code], vm.cur().Arg)
		}

		if cur.Code == OP_END {
			break
		}

		switch cur.Code {
		case OP_NOP:
		case OP_LOAD:
			vm.reg[0] = cur.Arg
		case OP_STORE:
			i := vm.regBoundCheck(cur.Arg)
			vm.reg[i] = vm.reg[0]
			vm.reg[0] = 0
		case OP_INSPECT:
			i := vm.regBoundCheck(cur.Arg)
			fmt.Printf("vm: %10s :: reg[%d] => %f\n", "INSPECT", i, vm.reg[i])
		case OP_ADD:
			i := vm.regBoundCheck(cur.Arg)
			vm.reg[0] = vm.reg[0] + vm.reg[i]
		case OP_SUBTRACT:
			i := vm.regBoundCheck(cur.Arg)
			vm.reg[0] = vm.reg[0] - vm.reg[i]
		case OP_MULTIPY:
			i := vm.regBoundCheck(cur.Arg)
			vm.reg[0] = vm.reg[0] * vm.reg[i]
		case OP_DIVIDE:
			i := vm.regBoundCheck(cur.Arg)
			vm.reg[0] = vm.reg[0] / vm.reg[i]
		default:
			log.Panicf("Unkown operator %v", OP_LOOKUP[cur.Code])
		}
		vm.advance()
	}
}
