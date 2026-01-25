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
	OPClear         // [-] [+]
)

type Instruction struct {
	Operation OPType
	Operand   int
}

var ErrUnmatchedBracket = errors.New("Unmatched Brackets Found")

func Compile(input string) ([]Instruction, error) {
	program := []Instruction{}
	loopStack := []int{}
	var lastInstruction *Instruction

	for _, c := range input {

		if len(program) > 0 {
			lastInstruction = &program[len(program)-1]
		}

		switch c {
		case '>':
			if lastInstruction != nil && lastInstruction.Operation == OPIncPtr {
				lastInstruction.Operand += 1
				continue
			}
			program = append(program, Instruction{Operation: OPIncPtr, Operand: 1})

		case '<':
			if lastInstruction != nil && lastInstruction.Operation == OPDecPtr {
				lastInstruction.Operand += 1
				continue
			}
			program = append(program, Instruction{Operation: OPDecPtr, Operand: 1})

		case '+':
			if lastInstruction != nil && lastInstruction.Operation == OPInc {
				lastInstruction.Operand += 1
				continue
			}
			program = append(program, Instruction{Operation: OPInc, Operand: 1})

		case '-':
			if lastInstruction != nil && lastInstruction.Operation == OPDec {
				lastInstruction.Operand += 1
				continue
			}
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
				return nil, ErrUnmatchedBracket
			}
			startPos := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]

			// optimise [-] & [+] Byte overflows and loops back to zero in GO
			if startPos+2 == len(program) {
				midInstruction := program[startPos+1]
				if midInstruction.Operation == OPInc || midInstruction.Operation == OPDec {
					if midInstruction.Operand == 1 {
						program = program[:startPos]
						program = append(program, Instruction{Operation: OPClear})
					}
				}
			}

			program[startPos].Operand = len(program)
			program = append(program, Instruction{Operation: OPJmpBwd, Operand: startPos})
		default:
			continue // Ignore all other characters
		}
	}

	if len(loopStack) > 0 {
		return nil, ErrUnmatchedBracket
	}

	return program, nil
}
