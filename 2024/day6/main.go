package main

import (
	"fmt"
	"slices"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[string]

	visited map[string][]string
}

var (
	grid = Grid{
		visited: make(map[string][]string),
	}
	currentX, currentY, startX, startY int
	direction          string
)

func main() {
	lines := internal.Reader()

	for y, line := range lines {
		for x, c := range line {
			grid.AddSafeToColumn(string(c), y)
			if string(c) == "^" {
				currentX = x
				currentY = y
				startX = x
				startY = y
				grid.visited[fmt.Sprintf("%d-%d", x, y)] = []string{"^"}
				direction = string(c)
			}
		}
	}
	doWalk()

	fmt.Printf("Part 1: %d\n", len(grid.visited))
	fmt.Printf("Part 2: %d\n", doPart2())
}

func doWalk() bool{
	for {
		switch direction {
		case "^":
			if currentY-1 < 0 {
				return false
			}
			if grid.GetSafeColumn(currentX, currentY-1) == "#" {
				direction = ">"
			} else {
				currentY--
			}
		case "V":
			if currentY+1 >= len(grid.Rows) {
				return false
			}
			if grid.GetSafeColumn(currentX, currentY+1) == "#" {
				direction = "<"
			} else {
				currentY++
			}
		case ">":
			if currentX+1 >= len(grid.Rows[0]) {
				return false
			}
			if grid.GetSafeColumn(currentX+1, currentY) == "#" {
				direction = "V"
			} else {
				currentX++
			}

		case "<":
			if currentX-1 < 0 {
				return false
			}
			if grid.GetSafeColumn(currentX-1, currentY) == "#" {
				direction = "^"
			} else {
				currentX--
			}
		}
		if !slices.Contains(grid.visited[fmt.Sprintf("%d-%d", currentX, currentY)], direction) {
			grid.visited[fmt.Sprintf("%d-%d", currentX, currentY)] = append(grid.visited[fmt.Sprintf("%d-%d", currentX, currentY)], direction)
		} else {
			return true
		}
	}
}

func doPart2() (total int){

	for y, row := range grid.Rows {
		for x, c := range row {
			grid.visited = make(map[string][]string)
			currentX = startX
			currentY = startY
			direction = "^"
			if string(c) == "." {
				grid.SetSafeColumn("#", x,y)
				if doWalk() {
					total++
				}
				grid.SetSafeColumn(".", x,y)
			}
		}
	}
	return
}