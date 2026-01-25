package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrOutOfMemoryBound = errors.New("Accessing Memory is out of Bounds")
var ErrReadingData = errors.New("Error Reading Data")

const MAX_MEMORY = 300000

func Execute(program []Instruction) (string, error) {
	output := ""
	memory := make([]byte, MAX_MEMORY)
	ptr := 0
	pc := 0
	reader := bufio.NewReader(os.Stdin)

	for pc < len(program) {
		instruction := program[pc]
		operand := instruction.Operand

		switch instruction.Operation {
		case OPIncPtr:
			if ptr+operand >= MAX_MEMORY {
				return "", ErrOutOfMemoryBound
			}
			ptr = ptr + operand
		case OPDecPtr:
			if ptr-operand < 0 {
				return "", ErrOutOfMemoryBound
			}
			ptr = ptr - operand
		case OPInc:
			memory[ptr] = memory[ptr] + byte(operand)
		case OPDec:
			memory[ptr] = memory[ptr] - byte(operand)
		case OPOut:
			output = output + string(memory[ptr])
		case OPIn:
			input, err := reader.ReadByte()
			if err != nil {
				fmt.Printf("Error Reading Data: %s", err.Error())
				return "", ErrReadingData
			}
			memory[ptr] = input
		case OPJmpFwd:
			if memory[ptr] == 0 {
				pc = operand
			}
		case OPJmpBwd:
			pc = operand
			continue // if we dont continue we go to instruction after than [
		case OPClear:
			memory[ptr] = 0
		default:
			return "", ErrUnmatchedBracket
		}

		pc++
	}

	return output, nil
}
