package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	patterns, towels []string
)

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, ",") {
			patterns = strings.Split(line, ", ")
		} else {
			towels = append(towels, line)
		}
	}

	cache := make(map[string]int)
	var solve func(towel string)int
	solve = func(towel string) int {
		if _, ok := cache[towel]; !ok {
			if len(towel) == 0 {
				return 1
			}
			var total int

			for _, pattern := range patterns {
				if strings.HasPrefix(towel, pattern) {
					total += solve(towel[len(pattern):])
				}
			}
			cache[towel] = total
		}
		return cache[towel]
	}
	var part1,part2 int

	for _, towel := range towels {
		out := solve(towel)
		if out > 0 {
			part1++
		}
		part2+=out
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
