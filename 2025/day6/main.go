package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	fmt.Printf("Part 1: %v\n", part1(lines))
	fmt.Printf("Part 2: %v\n", part2(lines))
}

func part1(lines []string) int {
	var part1 [][]int
	var ops []string
	for i, line := range lines {
		nmbrs := strings.Fields(line)
		if i == 0 {
			part1 = make([][]int, len(nmbrs))
		}
		if i != len(lines)-1 {
			for j, nmbr := range nmbrs {
				n, _ := strconv.Atoi(nmbr)
				part1[j] = append(part1[j], n)
			}
		} else {
			ops = nmbrs
		}
	}

	return sum(processProblems(part1, ops))
}

func processProblems(problems [][]int, ops []string) (total []int) {

	for j, op := range ops {
		var t int
		if op == "*" {
			t = 1
		}
		for _, n := range problems[j] {
			switch op {
			case "+":
				t += n
			case "*":
				t *= n
			}
		}
		total = append(total, t)
	}
	return
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func part2(lines []string) (total int) {
	if len(lines) == 0 {
		return 0
	}
	var ops []string

	maxLen := 0

	var layer1, layer2, layer3, layer4 []string
	var l1, l2, l3, l4 string
	ops = strings.Fields(lines[4])
	for i := 0; i < len(lines[0]); i++ {

		if string(lines[0][i]) == " " && string(lines[1][i]) == " " && string(lines[2][i]) == " " && string(lines[3][i]) == " " {
			layer1 = append(layer1, l1)
			layer2 = append(layer2, l2)
			layer3 = append(layer3, l3)
			layer4 = append(layer4, l4)
			l1 = ""
			l2 = ""
			l3 = ""
			l4 = ""
			if len(l1) > maxLen {
				maxLen = len(l1)
			}
			if len(l2) > maxLen {
				maxLen = len(l2)
			}
			if len(l3) > maxLen {
				maxLen = len(l3)
			}
			if len(l4) > maxLen {
				maxLen = len(l4)
			}
		} else {
			l1 += string(lines[0][i])
			l2 += string(lines[1][i])
			l3 += string(lines[2][i])
			l4 += string(lines[3][i])
		}
	}
	if l1 != "" {
		layer1 = append(layer1, l1)
		layer2 = append(layer2, l2)
		layer3 = append(layer3, l3)
		layer4 = append(layer4, l4)
	}

	for col := 0; col < len(layer1); col++ {
		l1, l2, l3, l4 := layer1[col], layer2[col], layer3[col], layer4[col]

		var numbers []string
		for d := 0; d < len(l1); d++ {
			numbers = append(numbers, string(l1[d])+string(l2[d])+string(l3[d])+string(l4[d]))
		}

		var t int
		if ops[col] == "*" {
			t = 1
		}

		for _, n := range numbers {
			nmr, _ := strconv.Atoi(strings.TrimSpace(n))
			if ops[col] == "*" {
				t *= nmr
			} else {
				t += nmr
			}
		}
		total += t
	}
	return
}
