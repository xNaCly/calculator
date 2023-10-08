package main

import (
	"fmt"
	"log"
)

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

// max amount of registers in virtual machine
const REGISTER_COUNT int = 4

type Vm struct {
	reg   [REGISTER_COUNT]float64 // registers
	in    []Operation             // operations to execute
	pos   int                     // current position in input
	trace bool                    // prints every operation if enabled
}

// assigns new input to the vm, resets its state
func (vm *Vm) NewVmIn(in []Operation) *Vm {
	vm.pos = 0
	vm.in = in
	vm.reg = [REGISTER_COUNT]float64{}
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
func regBoundCheck(index float64) int {
	i := int(index)
	if i < 0 || i > REGISTER_COUNT {
		log.Panicf("vm: Out of bounds register access for %d", i)
	}
	return i
}

func (vm *Vm) Execute() {
	for {
		cur := vm.cur()
		if vm.trace {
			fmt.Printf("vm: %10s :: %f\n", OP_LOOKUP[cur.Code], cur.Arg)
		}

		if cur.Code == OP_END {
			break
		}

		switch cur.Code {
		case OP_NOP:
		case OP_LOAD:
			vm.reg[0] = cur.Arg
		case OP_STORE:
			i := regBoundCheck(cur.Arg)
			vm.reg[i] = vm.reg[0]
			vm.reg[0] = 0
		case OP_INSPECT:
			i := regBoundCheck(cur.Arg)
			fmt.Printf("vm: %10s :: reg[%d] => %f\n", "INSPECT", i, vm.reg[i])
		case OP_ADD:
			i := regBoundCheck(cur.Arg)
			vm.reg[0] = vm.reg[i] + vm.reg[0]
		case OP_SUBTRACT:
			i := regBoundCheck(cur.Arg)
			vm.reg[0] = vm.reg[i] - vm.reg[0]
		case OP_MULTIPY:
			i := regBoundCheck(cur.Arg)
			vm.reg[0] = vm.reg[i] * vm.reg[0]
		case OP_DIVIDE:
			i := regBoundCheck(cur.Arg)
			vm.reg[0] = vm.reg[i] / vm.reg[0]
		default:
			log.Panicf("Unkown operator %v", OP_LOOKUP[cur.Code])
		}
		vm.advance()
	}
}
