package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const OP_HALT = 99
const OP_ADD = 1
const OP_MULT = 2

type VM struct {
	ip      int
	program []int
}

func NewVM(program []int) *VM {
	return &VM{
		0,
		program,
	}
}

func (vm *VM) Run() {
	for vm.program[vm.ip] != OP_HALT {
		switch vm.program[vm.ip] {
		case OP_MULT:
			i1, i2, o1 := vm.program[vm.ip+1], vm.program[vm.ip+2], vm.program[vm.ip+3]
			vm.program[o1] = vm.program[i1] * vm.program[i2]
			vm.ip += 4
			break
		case OP_ADD:
			i1, i2, o1 := vm.program[vm.ip+1], vm.program[vm.ip+2], vm.program[vm.ip+3]
			vm.program[o1] = vm.program[i1] + vm.program[i2]
			vm.ip += 4
			break
		case OP_HALT:
			return
		}
	}

	fmt.Println(vm.program)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()
		var numbers []int

		for _, i := range strings.Split(s, ",") {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, j)
		}

		vm := NewVM(numbers)
		vm.Run()
	}

}
