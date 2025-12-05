package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	var ranges [][]int
	var pt1 int
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			l, r := strings.Split(line, "-")[0], strings.Split(line, "-")[1]
			ln, _ := strconv.Atoi(l)
			rn, _ := strconv.Atoi(r)
			ranges = append(ranges, []int{ln, rn})

		} else {
			n, _ := strconv.Atoi(line)
			for _, ra := range ranges {
				if ra[0] <= n && n <= ra[1] {
					pt1++
					break
				}
			}
		}
	}

	pt2 := countUniqueInRanges(ranges)

	fmt.Printf("Part 1: %d\n", pt1)
	fmt.Printf("Part 2: %d\n", pt2)
}

func countUniqueInRanges(ranges [][]int) int {
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by start point
	for i := 0; i < len(ranges)-1; i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[i][0] > ranges[j][0] {
				ranges[i], ranges[j] = ranges[j], ranges[i]
			}
		}
	}

	// Merge overlapping ranges
	var merged [][]int
	current := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] <= current[1]+1 {
			// Overlapping or adjacent, merge
			if ranges[i][1] > current[1] {
				current[1] = ranges[i][1]
			}
		} else {
			// No overlap, save current and start new
			merged = append(merged, current)
			current = ranges[i]
		}
	}
	merged = append(merged, current)

	// Count total unique integers in merged ranges
	count := 0
	for _, r := range merged {
		count += r[1] - r[0] + 1
	}
	return count
}
