package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	var totalVars int
	for _, line := range lines {
		springs, groups := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		orig := springs
		oGroups := groups
		for i := 0; i < 5; i++ {
			springs += "?" + orig
			groups += "," + oGroups
		}
		groupings := getGroupings(groups)
		totalVars += generateSpringsP2([]byte(springs), 0, groupings)
	}
	fmt.Printf("Solution for Part1: %d\n", totalVars)
}

func validVariation(spring []byte, groupings []int) bool {
	g := 0
	for s := 0; s < len(spring); s++ {
		if spring[s] == '#' {
			startSpring := s
			// Advancing the current spring check
			for ; s < len(spring) && spring[s] == '#'; s++ {
			}
			if g < len(groupings) && s-startSpring != groupings[g] {
				return false
			} else {
				g++
			}
		}
	}
	return g == len(groupings)
}

func generateSpringsP1(springs []byte, index int, groups []int) int {
	for i := index; i < len(springs); i++ {
		if springs[i] == '?' {
			springsCp1, springsCp2 := make([]byte, len(springs)), make([]byte, len(springs))
			copy(springsCp1, springs)
			copy(springsCp2, springs)
			springsCp1[i] = '.'
			springsCp2[i] = '#'
			return generateSpringsP1(springsCp1, i+1, groups) + generateSpringsP1(springsCp2, i+1, groups)
		}
	}
	if validVariation(springs, groups) {
		return 1
	}
	return 0
}

func generateSpringsP2(springs []byte, index int, groups []int) int {
	for i := index; i < len(springs); i++ {
		if springs[i] == '?' {
			springsCp1, springsCp2 := make([]byte, len(springs)), make([]byte, len(springs))
			copy(springsCp1, springs)
			copy(springsCp2, springs)
			springsCp1[i] = '.'
			springsCp2[i] = '#'
			return generateSpringsP2(springsCp1, i+1, groups) + generateSpringsP2(springsCp2, i+1, groups)
		}
	}
	if validVariation(springs, groups) {
		return 1
	}
	return 0
}

func getGroupings(groups string) (result []int) {
	for _, g := range strings.Split(groups, ",") {
		i, _ := strconv.Atoi(g)
		result = append(result, i)
	}
	return
}
