package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
	"time"

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
	mainLoop                           = make(map[string][]string)
	direction                          string
)

func main() {
	lines := internal.Reader()
	start := time.Now()

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
	maps.Copy(mainLoop, grid.visited)

	fmt.Printf("Part 1: %d, Duration: %v\n", len(grid.visited), time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2: %d, Duration: %v\n", doPart2(), time.Since(start))
}

func doWalk() bool {
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

func doPart2() (total int) {

	for k := range mainLoop {
		t := strings.Split(k, "-")
		x, _ := strconv.Atoi(t[0])
		y, _ := strconv.Atoi(t[1])

		grid.visited = make(map[string][]string)
		currentX = startX
		currentY = startY
		direction = "^"
		if grid.GetSafeColumn(x, y) == "." {
			grid.SetSafeColumn("#", x, y)
			if doWalk() {
				total++
			}
			grid.SetSafeColumn(".", x, y)
		}
	}
	return
}
