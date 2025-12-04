package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	grid = new(internal.Grid[string])
)

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		row := strings.Split(line, "")
		grid.AddRow(row)
	}

	var pt1, pt2 int
	for i := 0; i < len(grid.Rows); i++ {
		row := grid.GetSafeRow(i)
		for j := 0; j < len(row); j++ {
			if row[j] == "@" && getValidNeighbour(j, i) < 4 {
				pt1++
			}
		}
	}
	fmt.Printf("Part 1: %d\n", pt1)
	for i := 0; i < len(grid.Rows); i++ {
		row := grid.GetSafeRow(i)
		for j := 0; j < len(row); j++ {
			if row[j] == "@" && getValidNeighbour(j, i) < 4 {
				pt2++
				grid.SetSafeColumn(".", j, i)
				i = -1
				break
			}
		}
	}
	fmt.Printf("Part 2: %d\n", pt2)
}

func getValidNeighbour(x, y int) (amount int) {
	// Check top-left
	if y-1 >= 0 {
		rows := grid.Rows[y-1]
		if x-1 >= 0 && rows[x-1] == "@" {
			amount++
		}
		// Check top
		if rows[x] == "@" {
			amount++
		}
		// Check top-right
		if x+1 < len(rows) && rows[x+1] == "@" {
			amount++
		}
	}
	// Check left and right
	rows := grid.Rows[y]
	if x-1 >= 0 && rows[x-1] == "@" {
		amount++
	}
	if x+1 < len(rows) && rows[x+1] == "@" {
		amount++
	}
	// Check bottom-left, bottom, bottom-right
	if y+1 < len(grid.Rows) {
		rows := grid.Rows[y+1]
		if x-1 >= 0 && rows[x-1] == "@" {
			amount++
		}
		// Check bottom
		if rows[x] == "@" {
			amount++
		}
		// Check bottom-right
		if x+1 < len(rows) && rows[x+1] == "@" {
			amount++
		}
	}
	return
}
