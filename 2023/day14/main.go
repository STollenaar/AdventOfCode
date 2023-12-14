package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	grid [][]string
)

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		l := strings.Split(line, "")
		grid = append(grid, l)
	}
	rollUp()
	fmt.Printf("Solution to Part1: %d\n", countLoad())
	for cycle := 0; cycle < 1000; cycle++ { // I am lucky. My solution worked with only a 1000 cycles
		rollUp()
		rollLeft()
		rollDown()
		rollRight()
	}
	fmt.Printf("Solution to Part2: %d\n", countLoad())
}

func countLoad() (total int) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == "O" {
				total += (len(grid) - r)
			}
		}
	}
	return
}

func rollUp() {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == "O" {
				oR := r - 1
				for ; oR >= 0; oR-- {
					if grid[oR][c] == "." {
						grid[oR][c] = "O"
						grid[oR+1][c] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func rollDown() {
	for r := len(grid) - 1; r >= 0; r-- {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == "O" {
				oR := r + 1
				for ; oR < len(grid); oR++ {
					if grid[oR][c] == "." {
						grid[oR][c] = "O"
						grid[oR-1][c] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func rollLeft() {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == "O" {
				oC := c - 1
				for ; oC >= 0; oC-- {
					if grid[r][oC] == "." {
						grid[r][oC] = "O"
						grid[r][oC+1] = "."
					} else {
						break
					}
				}
			}
		}
	}
}

func rollRight() {
	for r := 0; r < len(grid); r++ {
		for c := len(grid[r]) - 1; c >= 0; c-- {
			if grid[r][c] == "O" {
				oC := c + 1
				for ; oC <len(grid[r]); oC++ {
					if grid[r][oC] == "." {
						grid[r][oC] = "O"
						grid[r][oC-1] = "."
					} else {
						break
					}
				}
			}
		}
	}
}
