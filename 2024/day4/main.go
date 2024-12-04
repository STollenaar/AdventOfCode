package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	grid [][]string
	xmas = []string{"X", "M", "A", "S"}
)

func main() {
	lines := internal.Reader()

	var totalPart1, totalPart2 int
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	for y, line := range grid {
		for x, _ := range line {
			totalPart1 += part1StartSearch(x, y)
			totalPart2 += part2StartSearch(x, y)
		}
	}
	fmt.Printf("Part 1: %d\n", totalPart1)
	fmt.Printf("Part 2: %d\n", totalPart2)
}

func part1StartSearch(x, y int) (total int) {

	if grid[y][x] == "X" {
		if x-1 >= 0 {
			if y-1 >= 0 {
				total += doSearch(x-1, y-1, -1, -1, 1)
			}
			if y+1 < len(grid) {
				total += doSearch(x-1, y+1, -1, 1, 1)
			}
			total += doSearch(x-1, y, -1, 0, 1)
		}
		if x+1 < len(grid[0]) {
			if y-1 >= 0 {
				total += doSearch(x+1, y-1, 1, -1, 1)
			}
			if y+1 < len(grid) {
				total += doSearch(x+1, y+1, 1, 1, 1)
			}
			total += doSearch(x+1, y, 1, 0, 1)
		}
		if y-1 >= 0 {
			total += doSearch(x, y-1, 0, -1, 1)
		}
		if y+1 < len(grid) {
			total += doSearch(x, y+1, 0, 1, 1)
		}
	}
	return total
}

func part2StartSearch(x, y int) (total int) {
	if grid[y][x] == "A" {
		if y-1 >= 0 && x-1 >= 0 && y+1 < len(grid) && x+1 < len(grid[0]) {
			if ((grid[y-1][x-1] == "S" && grid[y+1][x+1] == "M") || (grid[y-1][x-1] == "M" && grid[y+1][x+1] == "S")) && ((grid[y-1][x+1] == "S" && grid[y+1][x-1] == "M") || (grid[y-1][x+1] == "M" && grid[y+1][x-1] == "S")) {
				return 1
			}
		}
	}
	return total
}

func doSearch(x, y, dx, dy, i int) int {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return 0
	}
	if grid[y][x] == xmas[i] {
		if i+1 == len(xmas) {
			return 1
		}
		return doSearch(x+dx, y+dy, dx, dy, i+1)
	}
	return 0
}
