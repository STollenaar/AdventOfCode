package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	program, registerA := load(lines)

	fmt.Printf("Part 1: %d\n", runProgram(program, registerA))
	out := part2(program)
	fmt.Printf("Part 2: %d, %d\n", out, runProgram(program, out))
}

func part2(program []uint64) (registerA uint64) {
	for itr := len(program)-1; itr >= 0; itr-- {
		registerA <<= 3
		for !slices.Equal(runProgram(program, registerA), program[itr:]) {
			registerA++
		}
	}
	return
}

func runProgram(program []uint64, registerA uint64) (out []uint64) {
	var instruction int
	registry := []uint64{registerA, 0, 0}

	getoperand := func(operand uint64) uint64 {
		switch operand {
		case 4:
			return registry[0]
		case 5:
			return registry[1]
		case 6:
			return registry[2]
		default:
			return operand
		}
	}

	for instruction < len(program)-1 {
		operator, operand := program[instruction], program[instruction+1]

		switch operator {
		case 0:
			registry[0] >>= getoperand(operand)
		case 1:
			registry[1] ^= operand
		case 2:
			registry[1] = getoperand(operand) & 7
		case 3:
			if registry[0] != 0 {
				instruction = int(operand)
				continue
			}
		case 4:
			registry[1] ^= registry[2]
		case 5:
			out = append(out, getoperand(operand)&7)
		case 6:
			registry[1] = registry[0] >> getoperand(operand)
		case 7:
			registry[2] = registry[0] >> getoperand(operand)

		}
		instruction += 2
	}
	return
}

func load(lines []string) (program []uint64, registerA uint64) {
	aS := strings.Split(lines[0], ": ")[1]
	registerA, _ = strconv.ParseUint(aS, 10, 64)

	pS := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	for _, pp := range pS {
		p, _ := strconv.ParseUint(pp, 10, 64)
		program = append(program, p)
	}
	return
}
