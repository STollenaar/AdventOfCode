package main

import (
	"fmt"
	"maps"
	"sync"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Point struct {
	symbol     string
	x, y, cost int
}

var (
	grid  internal.Grid[string]
	walls []Point
)

func main() {
	lines := internal.Reader()
	var start *Point
	for y, line := range lines {
		for x, c := range line {
			grid.AddSafeToColumn(string(c), y)
			if string(c) == "S" {
				start = &Point{x: x, y: y, symbol: "S", cost: 0}
			}
			if string(c) == "#" && x != 0 && y != 0 {
				walls = append(walls, Point{x: x, y: y})
			}
		}
	}

	baseline, r := solve(start, map[string]*Point{}, 0, 0)
	fmt.Printf("Baseline: %d, %d\n", baseline, r)
	_, cheats := solve(start, map[string]*Point{}, baseline-100, 1)
	fmt.Printf("Part 1: %d\n", cheats)
	_, cheats = solve(start, map[string]*Point{}, baseline-100, 20)
	fmt.Printf("Part 2: %d\n", cheats)

}

func solve(start *Point, visited map[string]*Point, maxCost, maxCheat int) (steps, results int) {
	queue := &internal.Queue[*Point]{}
	queue.EqualFunction = func(a, b *Point) bool {
		return a.x == b.x && a.y == b.y
	}

	moves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	// Helper function to generate a unique key for a point
	key := func(p *Point) string {
		return fmt.Sprintf("%d,%d", p.x, p.y)
	}

	queue.Push(start)
	var current *Point
	var mu sync.Mutex
	var wg sync.WaitGroup
	for len(queue.Elements) > 0 {
		current = queue.Shift()
		visited[key(current)] = current

		if current.symbol == "E" {
			results++
			break
		}

		for _, m := range moves {
			next := &Point{x: current.x + m[0], y: current.y + m[1], cost: current.cost + 1}
			if maxCost != 0 && next.cost > maxCost {
				continue
			}
			if next.x <= 0 || next.y <= 0 || next.x >= len(grid.Rows[0]) || next.y >= len(grid.Rows) {
				continue
			}
			next.symbol = grid.GetSafeColumn(next.x, next.y)
			if grid.GetSafeColumn(next.x, next.y) == "." || grid.GetSafeColumn(next.x, next.y) == "E" {
				if v, ok := visited[key(next)]; !ok || (ok && v.cost > next.cost) {
					if i := queue.FindIndex(next); i != -1 {
						if queue.Elements[i].cost < next.cost {
							continue
						}
						queue.Elements = append(queue.Elements[:i], queue.Elements[i+1:]...)
					}
					queue.Push(next)
				}
			} else if grid.GetSafeColumn(next.x, next.y) == "#" && maxCheat > 0 {
				wg.Add(1)
				v := make(map[string]*Point)
				maps.Copy(v, visited)
				go func(wg *sync.WaitGroup, mu *sync.Mutex) {
					defer wg.Done()
					s, r := solve(next, v, maxCost, maxCheat-1)
					if s <= maxCost {
						mu.Lock()
						results += r
						mu.Unlock()
					}
				}(&wg, &mu)
			}
		}
	}
	steps = current.cost
	wg.Wait()
	return
}
