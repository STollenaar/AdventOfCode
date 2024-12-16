package main

import (
	"fmt"
	"math"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[string]
}
type Queue struct {
	internal.Queue[*Step]
}

type Step struct {
	parent            *Step
	x, y, delta, cost int
	direction         string
}

var (
	grid  Grid
	queue = &Queue{}

	directions = map[string][]int{"<": {-1, 0}, "^": {0, -1}, ">": {1, 0}, "v": {0, 1}}
	costs      = map[string]map[string]int{
		"^": {">": 1, "v": 2001, "<": 1, "^": 1},
		">": {"^": 1, "v": 1, "<": 2001, ">": 1},
		"v": {">": 1, "^": 2001, "<": 1, "v": 1},
		"<": {">": 2001, "v": 1, "^": 1, "<": 1},
	}
	startX, startY, endX, endY int
)

func main() {
	lines := internal.Reader()
	queue.SortFunction = func(i, j int) bool {
		return queue.Elements[i].cost < queue.Elements[j].cost
	}

	for y, line := range lines {
		for x, c := range line {
			grid.AddSafeToColumn(string(c), y)
			if string(c) == "S" {
				startX = x
				startY = y
			}
			if string(c) == "E" {
				endX = x
				endY = y
			}
		}
	}
	queue.Push(&Step{x: startX, y: startY, cost: 1, direction: "S"})
	visisted := make(map[string]*Step)

	key := func(x, y int) string {
		return fmt.Sprintf("%d,%d", x, y)
	}

	var current *Step
	maxCost := math.MaxInt
	var paths []*Step
	for len(queue.Elements) > 0 {
		current = queue.Shift()
		visisted[key(current.x, current.y)] = current
		if grid.GetSafeColumn(current.x, current.y) == "E" {
			if maxCost <= math.MaxInt {
				maxCost = current.cost
			}
			paths = append(paths, current)
			continue
		}

		for dir, d := range directions {
			if grid.GetSafeColumn(current.x+d[0], current.y+d[1]) != "#" {
				step := &Step{x: d[0] + current.x, y: d[1] + current.y, direction: dir, cost: current.cost + costs[current.direction][dir], parent: current}
				step.delta = int(math.Abs(float64(step.x-endX))) + int(math.Abs(float64(step.y-endY)))
				if s, ok := visisted[key(step.x, step.y)]; ok && step.cost > s.cost {
					continue
				}
				if step.cost <= maxCost {
					if index := inQueue(step.x, step.y); index != -1 {
						if queue.Elements[index].cost > step.cost {
							queue.Elements = append(queue.Elements[:index], queue.Elements[index+1:]...)
						}
					}
					if step.direction != current.direction {
						step.cost += 1000
					}
					queue.Push(step)
				}
			}
		}
		queue.Sort()
	}

	fmt.Printf("Part 1: %d\n", paths[0].cost)

	var bestSeats int
	for _, path := range paths {
		current := path
		for current != nil {
			grid.SetSafeColumn("O", current.x, current.y)
			current = current.parent
			bestSeats++
		}
	}
	fmt.Printf("Part 2: %d\n", bestSeats)
	grid.print()
}

func (g *Grid) print() {
	for _, row := range g.Rows {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func inQueue(x, y int) int {
	for i, q := range queue.Elements {
		if q.x == x && q.y == y {
			return i
		}
	}
	return -1
}
