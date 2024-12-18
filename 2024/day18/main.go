package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[string]
}

type Point struct {
	parent            *Point
	x, y, delta, cost int
}

type Queue struct {
	internal.Queue[*Point]
}

var (
	grid Grid

	fallingBytes []Point

	endX, endY, fallAmount = 70, 70, 1024
)

func main() {
	lines := internal.Reader()

	for y := 0; y <= endY; y++ {
		for x := 0; x <= endX; x++ {
			grid.SetSafeColumn(".", x, y)
		}
	}

	for _, line := range lines {
		xs, ys := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		x, _ := strconv.Atoi(xs)
		y, _ := strconv.Atoi(ys)
		fallingBytes = append(fallingBytes, Point{x: x, y: y})
	}
	for i, b := range fallingBytes {
		if i == fallAmount {
			break
		}
		grid.SetSafeColumn("#", b.x, b.y)
	}

	fmt.Printf("Part 1: %d\n", solve())

	for _, fb := range fallingBytes {
		grid.SetSafeColumn("#", fb.x, fb.y)
		if solve() == -1 {
			fmt.Printf("Part 2: %d,%d\n", fb.x, fb.y)
            break
		}
	}
}

func solve() int {
	visited := make(map[string]*Point)
	queue := &Queue{}
	queue.EqualFunction = func(a, b *Point) bool {
		return a.x == b.x && a.y == b.y
	}
	queue.SortFunction = func(i, j int) bool {
		return queue.Elements[i].cost < queue.Elements[j].cost
	}

	moves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	// Helper function to generate a unique key for a point
	key := func(p *Point) string {
		return fmt.Sprintf("%d,%d", p.x, p.y)
	}

	queue.Push(&Point{x: 0, y: 0})
	var current *Point
	for len(queue.Elements) > 0 {
		current = queue.Shift()
		visited[key(current)] = current

		if current.x == endX && current.y == endY {
			break
		}

		for _, m := range moves {
			next := &Point{x: current.x + m[0], y: current.y + m[1], parent: current}
			// next.delta = int(math.Abs(float64(endX-next.x)) + math.Abs(float64(endY-next.y)))
			next.cost = current.cost + 1
			if next.x >= 0 && next.x <= endX && next.y >= 0 && next.y <= endY && grid.GetSafeColumn(next.x, next.y) == "." {
				if v, ok := visited[key(next)]; !ok || (ok && v.cost > next.cost) {
					if i := queue.FindIndex(next); i != -1 {
						if queue.Elements[i].cost < next.cost {
							continue
						}
						queue.Elements = append(queue.Elements[:i], queue.Elements[i+1:]...)
					}
					queue.Push(next)
				}
			}
		}
		queue.Sort()
	}
	if current.x != endX || current.y != endY {
		return -1
	}
	var steps int
	// grid.Print()
	for current != nil {
		// grid.SetSafeColumn("O", current.x, current.y)
		steps++
		current = current.parent
	}
	steps--
	// grid.Print()
	return steps
}
