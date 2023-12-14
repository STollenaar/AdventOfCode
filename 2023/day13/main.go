package main

import (
	"fmt"
	"slices"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	grids [][]string
)

func main() {
	lines := internal.Reader()

	var grid []string
	for _, line := range lines {
		if line == "" {
			grids = append(grids, grid)
			grid = []string{}
		} else {
			grid = append(grid, line)
		}
	}
	grids = append(grids, grid)

	var p1Total, p2Total int

	for _, grid := range grids {
		p1Total += (getMirrorCount(grid, slices.Equal) * 100) + getMirrorCount(transpose(grid), slices.Equal)
		p2Total += (getMirrorCount(grid, checkSmudge) * 100) + getMirrorCount(transpose(grid), checkSmudge)
	}
	fmt.Printf("Solution to part1: %d\n", p1Total)
	fmt.Printf("Solution to part2: %d\n", p2Total)
}

func getMirrorCount(grid []string, isMirrored func([]string, []string) bool) int {
	for i := 1; i < len(grid); i++ {
		j := slices.Min([]int{i, len(grid) - i})
		mirrored := grid[i : i+j]
		orig := slices.Clone(grid[i-j : i])
		slices.Reverse(orig)
		if isMirrored(orig, mirrored) {
			return i
		}
	}

	return 0
}

func checkSmudge(orig, mirr []string) bool {
	var diffs int
	for i := range orig {
		for j := range orig[i] {
			if orig[i][j] != mirr[i][j] {
				diffs++
			}
		}
	}
	return diffs == 1
}

func transpose(grid []string) (result []string) {
	for i := 0; i < len(grid[0]); i++ {
		var buffer string
		for j := 0; j < len(grid); j++ {
			buffer += string(grid[j][i])
		}
		result = append(result, buffer)
	}
	return
}
