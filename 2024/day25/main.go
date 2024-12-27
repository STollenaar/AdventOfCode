package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	var current internal.Grid[string]

	var all []internal.Grid[string]
	var keys, locks [][]int
	for _, line := range lines {
		if line == "" {
			all = append(all, current)
			current = internal.Grid[string]{}
		} else {
			current.AddRow(strings.Split(line, ""))
		}
	}
	all = append(all, current)
	for _, schem := range all {
		var profile []int
		for x := 0; x < 5; x++ {
			var h int
			for ; h < 6; h++ {
				if schem.GetSafeColumn(x, 0) != schem.GetSafeColumn(x, h) {
					break
				}
			}
			profile = append(profile, h-1)
		}

		if schem.GetSafeColumn(0, 0) == "#" {
			locks = append(locks, profile)
		} else {
			for i, p := range profile {
				profile[i] = 5 - p
			}
			keys = append(keys, profile)
		}
	}

	var part1Total int
	for _, lock := range locks {
	keysLoop:
		for _, key := range keys {
			for i := 0; i < 5; i++ {
				if lock[i]+key[i] >= 6 {
					continue keysLoop
				}
			}
			part1Total++
		}
	}
	fmt.Printf("Part1: %d\n", part1Total)
}
