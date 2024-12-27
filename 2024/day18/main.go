package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	grid internal.Grid[string]

	fallingBytes []internal.Point[int]

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
		fallingBytes = append(fallingBytes, internal.Point[int]{X: x, Y: y})
	}
	for i, b := range fallingBytes {
		if i == fallAmount {
			break
		}
		grid.SetSafeColumn("#", b.X, b.Y)
	}

	fmt.Printf("Part 1: %d\n", solve())

	for _, fb := range fallingBytes {
		grid.SetSafeColumn("#", fb.X, fb.Y)
		if solve() == -1 {
			fmt.Printf("Part 2: %d,%d\n", fb.X, fb.Y)
			break
		}
	}
}

func solve() int {
	visited := make(map[string]*internal.Point[int])
	queue := &	internal.Queue[*internal.Point[int]]{}
	queue.EqualFunction = func(a, b *internal.Point[int]) bool {
		return a.X == b.X && a.Y == b.Y
	}
	queue.SortFunction = func(i, j int) bool {
		return queue.Elements[i].Cost < queue.Elements[j].Cost
	}

	moves := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	// Helper function to generate a unique key for a point
	key := func(p *internal.Point[int]) string {
		return fmt.Sprintf("%d,%d", p.X, p.Y)
	}

	queue.Push(&internal.Point[int]{X: 0, Y: 0})
	var current *internal.Point[int]
	for len(queue.Elements) > 0 {
		current = queue.Shift()
		visited[key(current)] = current

		if current.X == endX && current.Y == endY {
			break
		}

		for _, m := range moves {
			next := &internal.Point[int]{X: current.X + m[0], Y: current.Y + m[1], Parent: current}
			// next.delta = int(math.Abs(float64(endX-next.X)) + math.Abs(float64(endY-next.Y)))
			next.Cost = current.Cost + 1
			if next.X >= 0 && next.X <= endX && next.Y >= 0 && next.Y <= endY && grid.GetSafeColumn(next.X, next.Y) == "." {
				if v, ok := visited[key(next)]; !ok || (ok && v.Cost > next.Cost) {
					if i := queue.FindIndex(next); i != -1 {
						if queue.Elements[i].Cost < next.Cost {
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
	if current.X != endX || current.Y != endY {
		return -1
	}
	var steps int
	// grid.Print()
	for current != nil {
		// grid.SetSafeColumn("O", current.X, current.Y)
		steps++
		current = current.Parent
	}
	steps--
	// grid.Print()
	return steps
}
