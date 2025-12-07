package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	grid internal.Grid[string]
)

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		grid.AddRow(strings.Split(line, ""))
	}
	var splits int
	for y := 0; y < len(grid.Rows); y++ {
		for x := 0; x < len(grid.Rows[y]); x++ {
			if y > 0 {
				if grid.GetUnsafeColumn(x, y-1) == "|" || grid.GetUnsafeColumn(x, y-1) == "S" {
					if grid.GetUnsafeColumn(x, y) == "." {
						grid.SetUnsafeColumn("|", x, y)
					} else if grid.GetUnsafeColumn(x, y) == "^" {
						splits++
						if x > 0 && grid.GetUnsafeColumn(x-1, y) == "." {
							grid.SetUnsafeColumn("|", x-1, y)
						}
						if x+1 < len(grid.Rows[y]) && grid.GetUnsafeColumn(x+1, y) == "." {
							grid.SetUnsafeColumn("|", x+1, y)
						}
					}
				}
			}
		}
	}
	// grid.Print()

	fmt.Printf("Part 1: %d\n", splits)

	sx, sy := -1, -1
	for y := range grid.Rows {
		for x := range grid.Rows[y] {
			if grid.GetUnsafeColumn(x, y) == "S" {
				sx, sy = x, y
				break
			}
		}
		if sx != -1 {
			break
		}
	}
	if sx == -1 {
		fmt.Println("Part 2: start 'S' not found")
		return
	}

	memo := map[string]int{}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if y == len(grid.Rows)-1 {
			return 1
		}
		key := fmt.Sprintf("%d,%d", x, y)
		if v, ok := memo[key]; ok {
			return v
		}
		total := 0

		if y+1 < len(grid.Rows) {
			below := grid.GetSafeColumn(x, y+1)
			switch below {
			case "|":
				total += dfs(x, y+1)
			case "^":
				if x-1 >= 0 && grid.GetSafeColumn(x-1, y+1) == "|" {
					total += dfs(x-1, y+1)
				}
				if x+1 < len(grid.Rows[y+1]) && grid.GetSafeColumn(x+1, y+1) == "|" {
					total += dfs(x+1, y+1)
				}
			}
		}

		memo[key] = total
		return total
	}

	paths := dfs(sx, sy)
	fmt.Printf("Part 2: %d\n", paths)
}
