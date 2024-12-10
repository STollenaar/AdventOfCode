package main

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[int]
}

var (
	grid   Grid
	starts [][]int
	trails = make(map[string][]string)
)

func main() {
	lines := internal.Reader()
	
	for y, line := range lines {
		for x, c := range line {
			n, _ := strconv.Atoi(string(c))
			grid.AddSafeToColumn(n, y)
			if n == 0 {
				starts = append(starts, []int{x, y})
			}
		}
	}
	start := time.Now()
	var totalPart1, totalPart2 int
	for _, st := range starts {
		score := doWalk(st[0], st[1], st[0], st[1], true)
		totalPart1 += score
	}
	fmt.Printf("Part 1: %d, Duration: %v\n", totalPart1, time.Since(start))
	start = time.Now()
	for _, st := range starts {
		score := doWalk(st[0], st[1], st[0], st[1], false)
		totalPart2 += score
	}
	fmt.Printf("Part 2: %d, Duration: %v\n", totalPart2, time.Since(start))
}

func doWalk(sx, sy, x, y int, checkTrails bool) (total int) {
	current := grid.GetSafeColumn(x, y)
	if current == 9 {
		if checkTrails && !slices.Contains(trails[fmt.Sprintf("%d-%d", sx, sy)], fmt.Sprintf("%d-%d", x, y)) {
			trails[fmt.Sprintf("%d-%d", sx, sy)] = append(trails[fmt.Sprintf("%d-%d", sx, sy)], fmt.Sprintf("%d-%d", x, y))
			return 1
		} else if checkTrails {
			return 0
		} else {
			return 1
		}
	}
	if x-1 >= 0 && grid.GetSafeColumn(x-1, y) == current+1 {
		total += doWalk(sx, sy, x-1, y, checkTrails)
	}
	if x+1 < len(grid.Rows[0]) && grid.GetSafeColumn(x+1, y) == current+1 {
		total += doWalk(sx, sy, x+1, y, checkTrails)

	}
	if y-1 >= 0 && grid.GetSafeColumn(x, y-1) == current+1 {
		total += doWalk(sx, sy, x, y-1, checkTrails)

	}
	if y+1 < len(grid.Rows) && grid.GetSafeColumn(x, y+1) == current+1 {
		total += doWalk(sx, sy, x, y+1, checkTrails)
	}
	return
}
