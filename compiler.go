package main

import "errors"

type OPType int

const (
	OPIncPtr = iota // >
	OPDecPtr        // <
	OPInc           // +
	OPDec           // -
	OPOut           // .
	OPIn            // ,
	OPJmpFwd        // [
	OPJmpBwd        // ]
)

type Instruction struct {
	Operation OPType
	Operand   int
}

var ErrSyntaxError = errors.New("Syntax Error")

func Compile(input string) ([]Instruction, error) {
	program := []Instruction{}
	loopStack := []int{}

	for _, c := range input {
		switch c {
		case '>':
			program = append(program, Instruction{Operation: OPIncPtr, Operand: 1})
		case '<':
			program = append(program, Instruction{Operation: OPDecPtr, Operand: 1})
		case '+':
			program = append(program, Instruction{Operation: OPInc, Operand: 1})
		case '-':
			program = append(program, Instruction{Operation: OPDec, Operand: 1})
		case '.':
			program = append(program, Instruction{Operation: OPOut})
		case ',':
			program = append(program, Instruction{Operation: OPIn})
		case '[':
			loopStack = append(loopStack, len(program))
			program = append(program, Instruction{Operation: OPJmpFwd, Operand: 0})
		case ']':
			if len(loopStack) == 0 {
				return nil, ErrSyntaxError
			}
			jmpPos := loopStack[len(loopStack)-1]
			program[jmpPos].Operand = len(program)
			program = append(program, Instruction{Operation: OPJmpBwd, Operand: jmpPos})
		default:
			return nil, ErrSyntaxError
		}
	}

	return program, nil
}
