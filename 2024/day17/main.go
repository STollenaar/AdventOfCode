package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	registry, program []int
	instruction       int
	outputs           []string

	operation = map[int]func(c int){0: adv, 1: bxl, 2: bst, 3: jnz, 4: bxc, 5: out, 6: bdv, 7: cdv}
)

func main() {
	lines := internal.Reader()

	load(lines)
	inputProgram, inputRegistyA := strings.Split(lines[4], ": ")[1], registry[0]
	for ; instruction < len(program); instruction += 2 {
		operation[program[instruction]](program[instruction+1])
	}

	fmt.Printf("Part 1: %s\n", strings.Join(outputs, ","))

	i := 1
	for {
		if i == inputRegistyA {
			i++
			continue
		}
		load(lines)
		registry[0] = i
		for ; instruction < len(program); instruction += 2 {
			operation[program[instruction]](program[instruction+1])
		}
		if strings.Join(outputs, ",") == inputProgram {
			fmt.Printf("Part 2: %d\n", i)
			break
		}
		i++
	}
}

func adv(combo int) {
	denum := int(math.Pow(2, float64(getCombo(combo))))
	registry[0] = registry[0] / denum
}

func bxl(combo int) {
	registry[1] = registry[1] ^ combo
}

func bst(combo int) {
	registry[1] = getCombo(combo) % 8
}

func jnz(combo int) {
	if registry[0] == 0 {
		return
	}
	instruction = combo
	instruction -= 2
}

func bxc(combo int) {
	registry[1] = registry[1] ^ registry[2]
}

func out(combo int) {
	outputs = append(outputs, strconv.Itoa(getCombo(combo)%8))
}

func bdv(combo int) {
	denum := int(math.Pow(2, float64(getCombo(combo))))
	registry[1] = registry[0] / denum
}

func cdv(combo int) {
	denum := int(math.Pow(2, float64(getCombo(combo))))
	registry[2] = registry[0] / denum
}

func getCombo(combo int) int {
	switch combo {
	case 4:
		return registry[0]
	case 5:
		return registry[1]
	case 6:
		return registry[2]
	default:
		return combo
	}
}

func load(lines []string) {
	aS := strings.Split(lines[0], ": ")[1]
	a, _ := strconv.Atoi(aS)
	bS := strings.Split(lines[1], ": ")[1]
	b, _ := strconv.Atoi(bS)
	cS := strings.Split(lines[2], ": ")[1]
	c, _ := strconv.Atoi(cS)
	registry = []int{a, b, c}

	pS := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	program = []int{}
	outputs = []string{}
	instruction = 0
	for _, pp := range pS {
		p, _ := strconv.Atoi(pp)
		program = append(program, p)
	}
}
