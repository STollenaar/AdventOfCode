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

	key := func(x, y int) string {
		return fmt.Sprintf("%d,%d", x, y)
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

	var current *Step
	maxCost := math.MaxInt
	var paths []*Step

	for len(queue.Elements) > 0 {
		current = queue.Shift()

		// Stop if we've exceeded the maxCost
		if current.cost > maxCost {
			break
		}

		// Check if we've reached the end
		if grid.GetSafeColumn(current.x, current.y) == "E" {
			if current.cost < maxCost {
				maxCost = current.cost
				paths = []*Step{current} // Reset paths since we've found a better cost
			} else if current.cost == maxCost {
				paths = append(paths, current)
			}
			continue
		}

		for dir, d := range directions {
			if grid.GetSafeColumn(current.x+d[0], current.y+d[1]) != "#" {
				// Calculate base cost for the next step
				stepCost := current.cost + costs[current.direction][dir]

				// Apply direction change penalty
				if dir != current.direction {
					stepCost += 1000
				}

				// Create the new step
				step := &Step{
					x:         d[0] + current.x,
					y:         d[1] + current.y,
					direction: dir,
					cost:      stepCost,
					parent:    current,
				}
				step.delta = int(math.Abs(float64(step.x-endX))) + int(math.Abs(float64(step.y-endY)))

				if step.cost <= maxCost {
					if index := inQueue(step.x, step.y); index != -1 {
						if queue.Elements[index].cost > step.cost {
							queue.Elements = append(queue.Elements[:index], queue.Elements[index+1:]...)
						}
					}
					queue.Push(step)
				}
			}
		}
		queue.Sort()
	}

	fmt.Printf("Part 1: %d\n", maxCost)

	// Print all best paths
	var bestSeats int
	counted := make(map[string]bool)
	for _, path := range paths {
		current := path
		for current != nil {
			if _, ok := counted[key(current.x, current.y)]; !ok {
				counted[key(current.x, current.y)] = true
				bestSeats++
			}
			current = current.parent
		}
	}
	fmt.Printf("Part 2: %d\n", bestSeats)
}

func inQueue(x, y int) int {
	for i, q := range queue.Elements {
		if q.x == x && q.y == y {
			return i
		}
	}
	return -1
}
