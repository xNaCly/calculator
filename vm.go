package main

import (
	"fmt"
	"log"
)

// represents an operation the virtual machine performs
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
}

// represents an operation and its argument
type Operation struct {
	Code OpCode  // type of operation
	Arg  float64 // operation argument
}

// max amount of registers in virtual machine
const REGISTER_COUNT int = 128

var CUR_REG float64 = 1

// The virtual machine (VM) is a way to simulate the inner workings of a processor.
//
// Programming using the byte code accepted by this virtual machine is a very
// assembly like experience. The VM exposes the ability to load data into
// registers, store this data in different registers and perform arithmetic
// operations on these registers.
//
// Compiling an abstract syntax tree to bytecode and executing this bytecode is
// significantly faster than a tree walk interpreter, this does however come at
// the cost of greater complexity in both the code base (having to walk the
// tree, compile to bytecode, executing bytecode in a VM) and the mental model
// the developer has of executing a programming language.
//
// Reference for acceptable bytecode:
//
//   - OP_NOP                      ; no operation
//   - OP_LOAD     <value>         ; loads the given 'value' into register 0
//   - OP_STORE    <register>      ; moves the value of register 0 to 'register'
//   - OP_ADD      <register>      ; adds the value of register 0 to the value at 'register', stores result in register 0
//   - OP_SUBTRACT <register>      ; subtracts the value of register 0 from the value at 'register', stores result in register 0
//   - OP_MULTIPY  <register>      ; multiplies the value of register 0 with the value at 'register', stores result in register 0
//   - OP_DIVIDE   <register>      ; divides the value of register 0 with the value at 'register', stores result in register 0
//   - OP_INSPECT  <register>      ; prints the value of 'register'
//
// All results operations such as OP_ADD generate are stored in register 0. The
// amount of available registers is defined in REGISTER_COUNT and by default
// set to 4. The VM expects the last instruction to contain the Operation code
// (OP_CODE) OP_END, otherwise it will be stuck in an endless loop.
type Vm struct {
	reg   [REGISTER_COUNT]float64 // registers
	in    []Operation             // operations to execute
	pos   int                     // current position in input
	trace bool                    // prints every operation if enabled
	atEnd bool                    // indicates if the vm reached the end of the input
}

// assigns new input to the vm, resets its state
func (vm *Vm) NewVmIn(in []Operation) *Vm {
	vm.pos = 0
	vm.in = in
	vm.reg = [REGISTER_COUNT]float64{}
	vm.atEnd = false
	return vm
}

// if next position in vm.in boundary, increment position
func (vm *Vm) advance() {
	if vm.pos+1 < len(vm.in) {
		vm.pos++
	} else {
		vm.atEnd = true
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
	if len(vm.in) == 0 {
		return
	}
	for !vm.atEnd {
		cur := vm.cur()
		if vm.trace {
			fmt.Printf("%-10s :: %f\n", OP_LOOKUP[cur.Code], cur.Arg)
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
