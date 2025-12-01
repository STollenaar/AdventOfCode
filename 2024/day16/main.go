package main

import (
	"fmt"
	"math"
	"time"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Step struct {
	parent            *Step
	x, y, delta, cost int
	direction         string
}

var (
	grid  internal.Grid[string]
	queue = &internal.Queue[*Step]{}

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
	start := time.Now()
	queue.Push(&Step{x: startX, y: startY, cost: 1, direction: "S"})

	var current *Step
	maxCost := math.MaxInt
	var paths []*Step
	visited := make(map[string]*Step)
	for len(queue.Elements) > 0 {
		current = queue.Shift()
		visited[key(current.x, current.y)] = current
		// Stop if we've exceeded the maxCost
		if current.cost > maxCost {
			break
		}

		if current.parent != nil && current.direction != current.parent.direction {
			np := &Step{direction: current.direction, x: current.parent.x, y: current.parent.y, cost: current.parent.cost + 1000, delta: current.parent.delta, parent: current.parent.parent}
			current.parent = np
			current.cost += 1000
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

				// Create the new step
				step := &Step{
					x:         d[0] + current.x,
					y:         d[1] + current.y,
					direction: dir,
					cost:      stepCost,
					parent:    current,
				}
				if vStep, ok := visited[key(step.x, step.y)]; ok {
					if (vStep.cost < step.cost && vStep.direction == step.direction) || (vStep.cost+1000 < step.cost && step.direction != vStep.direction) {
						continue
					}
				}
				step.delta = int(math.Abs(float64(step.x-endX))) + int(math.Abs(float64(step.y-endY)))

				if step.cost <= maxCost {
					queue.Push(step)
				}
			}
		}
		queue.Sort()
	}

	fmt.Printf("Done doing pathing: %v\n", time.Since(start))
	fmt.Printf("Part 1: %d\n", maxCost)

	// Print all best paths
	var bestSeats int
	counted := make(map[string]bool)
	for _, path := range paths {
		current := path
		for current != nil {
			if _, ok := counted[key(current.x, current.y)]; !ok {
				grid.SetSafeColumn("O", current.x, current.y)
				counted[key(current.x, current.y)] = true
				bestSeats++
			}
			current = current.parent
		}
	}
	fmt.Printf("Part 2: %d\n", bestSeats)
}
