package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type Bracket struct {
	opening string
	closing string
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	defer doPart1(lines, start)

}
func doPart1(lines []string, start time.Time) {

	errorScore := 0
	var incompletes [][]Bracket
	for _, line := range lines {
		var stack []Bracket
		breakLoop := false
		for _, char := range strings.Split(line, "") {
			switch char {
			case "<":
				stack = append(stack, Bracket{opening: "<", closing: ">"})
			case ">":
				current := stack[len(stack)-1]
				if char != current.closing {
					errorScore += 25137
					breakLoop = true
				}
				stack = stack[:len(stack)-1]
			case "[":
				stack = append(stack, Bracket{opening: "[", closing: "]"})
			case "]":
				current := stack[len(stack)-1]
				if char != current.closing {
					errorScore += 57
					breakLoop = true
				}
				stack = stack[:len(stack)-1]
			case "{":
				stack = append(stack, Bracket{opening: "{", closing: "}"})
			case "}":
				current := stack[len(stack)-1]
				if char != current.closing {
					errorScore += 1197
					breakLoop = true
				}
				stack = stack[:len(stack)-1]
			case "(":
				stack = append(stack, Bracket{opening: "(", closing: ")"})
			case ")":
				current := stack[len(stack)-1]
				if char != current.closing {
					errorScore += 3
					breakLoop = true
				}
				stack = stack[:len(stack)-1]
			}
			if breakLoop {
				break
			}
		}
		if !breakLoop {
			incompletes = append(incompletes, stack)
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Total error score for part 1:", errorScore)

	start = time.Now()
	doPart2(incompletes, start)
}

func doPart2(lines [][]Bracket, start time.Time) {
	var scores []int
	for _, stack := range lines {
		score := 0
		for len(stack) > 0 {
			bracket := stack[len(stack)-1]
			score *= 5
			switch bracket.closing {
			case ")":
				score += 1
			case "]":
				score += 2
			case "}":
				score += 3
			case ">":
				score += 4
			}
			stack = stack[:len(stack)-1]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	winner := scores[(len(scores)-1)/2]
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Winner error score for part 2:", winner)
}
