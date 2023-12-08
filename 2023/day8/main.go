package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Node struct {
	leftNode, rightNode string
}

var (
	nodes map[string]Node
	reg   = regexp.MustCompile(`[)(,=]`)
)

func init() {
	nodes = make(map[string]Node)
}

func main() {
	lines := internal.Reader()

	sequence := strings.Split(lines[0], "")
	lines = lines[1:]
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = reg.ReplaceAllString(line, "")
		node := strings.Split(line, " ")[0]
		paths := strings.Split(line, " ")[1:]
		nodes[node] = Node{leftNode: paths[1], rightNode: paths[2]}
	}
	part1(sequence)
	part2(sequence)
}

func part1(sequence []string) {
	currentNode := "AAA"
	var i, steps int
	for currentNode != "ZZZ" {
		if i == len(sequence) {
			i = 0
		}
		if sequence[i] == "L" {
			currentNode = nodes[currentNode].leftNode
		} else {
			currentNode = nodes[currentNode].rightNode
		}
		i++
		steps++
	}
	fmt.Printf("Solution to part1: %d\n", steps)
}

func part2(sequence []string) {
	currentNodes := getStartNodes()

	var Allsteps []int
	for _, currentNode := range currentNodes {
		var seqI, steps int
		for string(currentNode[2]) != "Z" {
			if seqI == len(sequence) {
				seqI = 0
			}
			if sequence[seqI] == "L" {
				currentNode = nodes[currentNode].leftNode
			} else {
				currentNode = nodes[currentNode].rightNode
			}
			seqI++
			steps++
		}
		Allsteps = append(Allsteps, steps)
	}
	lcm := LCM(Allsteps[0], Allsteps[1], Allsteps[2:]...)
	fmt.Printf("Solution to part2: %d\n", lcm)
}
func getStartNodes() (result []string) {
	for k := range nodes {
		if string(k[2]) == "A" {
			result = append(result, k)
		}
	}
	return
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}