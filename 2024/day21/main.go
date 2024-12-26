package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Queue struct {
	internal.Queue[*Point]
}

type Point struct {
	parent      *Point
	cost, x, y  int
	symbol, dir string
}

var (
	pathCache    = make(map[string]*Point)
	segmentCache = make(map[string]string)

	directions = map[string][]int{
		"^": {0, -1},
		">": {1, 0},
		"v": {0, 1},
		"<": {-1, 0},
	}

	costMap = map[string]map[string]int{
		"^": {
			">": 2,
			"v": 1,
			"<": 2,
			"A": 1,
		},
		">": {
			"^": 2,
			"v": 1,
			"<": 2,
			"A": 1,
		},
		"v": {
			">": 1,
			"^": 1,
			"<": 1,
			"A": 2,
		},
		"<": {
			">": 2,
			"v": 1,
			"^": 2,
			"A": 3,
		},
		"A": {
			">": 1,
			"v": 2,
			"<": 3,
			"^": 1,
		},
	}

	keypad = map[string][]int{
		"7": {0, 0},
		"8": {1, 0},
		"9": {2, 0},
		"4": {0, 1},
		"5": {1, 1},
		"6": {2, 1},
		"1": {0, 2},
		"2": {1, 2},
		"3": {2, 2},
		"X": {0, 3},
		"0": {1, 3},
		"A": {2, 3},
	}

	robot = map[string][]int{
		"X": {0, 0},
		"A": {2, 0},
		"^": {1, 0},
		"<": {0, 1},
		"v": {1, 1},
		">": {2, 1},
	}
)

func main() {
	lines := internal.Reader()

	var part1Total, part2Total int
	for _, line := range lines {
		steps := findSteps(line, 3)
		nmbr, _ := strconv.Atoi(strings.ReplaceAll(line, "A", ""))
		part1Total += nmbr * len(steps)
		fmt.Println(nmbr, len(steps), steps)
	}
	fmt.Printf("Part1: %d\n", part1Total)
	// continue for infinite runtime
	for _, line := range lines {
		steps := findSteps(line, 25)
		nmbr, _ := strconv.Atoi(strings.ReplaceAll(line, "A", ""))
		part2Total += nmbr * len(steps)
		fmt.Println(nmbr, len(steps))
	}
	fmt.Printf("Part2: %d\n", part2Total)
}

func findSteps(code string, maxLayer int) string {
	layerSteps := []string{code}
	for i := 0; i < maxLayer; i++ {
		layer := layerSteps[i]
		start := "A"
		var steps []string
		for _, c := range layer {
			moves := findMoves(start, string(c), i, maxLayer)
			path := reconstructPath(moves)
			steps = append(steps, path...)
			start = string(c)
		}
		layerStep := strings.Join(steps, "")
		layerSteps = append(layerSteps, layerStep)
	}

	return layerSteps[len(layerSteps)-1]
}

// Find the smallest path for each layer recursively
func findMoves(start, end string, layer, maxLayer int) *Point {
	cacheKey := fmt.Sprintf("%s-%s", start, end)
	if cached, ok := pathCache[cacheKey]; ok {
		return cached
	}

	var grid map[string][]int
	if layer == 0 {
		grid = keypad
	} else {
		grid = robot
	}

	var queue Queue
	queue.SortFunction = func(i, j int) bool {
		if queue.Elements[i].cost == queue.Elements[j].cost {
			return queue.Elements[i].symbol < queue.Elements[j].symbol
		}
		return queue.Elements[i].cost < queue.Elements[j].cost
	}

	queue.Push(&Point{symbol: start, cost: 0, x: grid[start][0], y: grid[start][1]})

	var paths []*Point
	visited := make(map[string]int)
	key := func(p *Point) string {
		return fmt.Sprintf("%d,%d,%s", p.x, p.y, p.symbol)
	}

	for len(queue.Elements) > 0 {
		current := queue.Shift()

		if c, ok := visited[key(current)]; ok && c < current.cost {
			continue
		}
		visited[key(current)] = current.cost

		currentPos := grid[current.symbol]
		if current.symbol == end {
			paths = append(paths, current)
			continue
		}

		for dir, v := range directions {
			nextPos := []int{currentPos[0] + v[0], currentPos[1] + v[1]}
			if grid["X"][0] == nextPos[0] && grid["X"][1] == nextPos[1] {
				continue
			}

			for s, p := range grid {
				if p[0] == nextPos[0] && p[1] == nextPos[1] {
					cost := current.cost + costMap[current.dir][dir]
					nextPoint := &Point{
						parent: current,
						symbol: s,
						dir:    dir,
						cost:   cost,
						x:      nextPos[0],
						y:      nextPos[1],
					}

					if c, ok := visited[key(nextPoint)]; ok && c < nextPoint.cost {
						continue
					}
					queue.Push(nextPoint)
				}
			}
		}
		queue.Sort()
	}

	// Handle if no path was found
	if len(paths) == 0 {
		return &Point{
			cost: 999999999,
		}
	}

	if layer+1 < maxLayer {
		for _, path := range paths {
			p := reconstructPath(path)
			st := "A"
			for _, step := range p {
				refinedPath := findMoves(st, step, layer+1, maxLayer)
				st = step
				path.cost += refinedPath.cost
			}
		}
	}

	// Get the best path for this layer
	smallest := paths[0]
	for _, s := range paths {
		if s.cost < smallest.cost {
			smallest = s
		}
	}
	pathCache[cacheKey] = smallest

	return smallest
}

func reconstructPath(path *Point) (out []string) {
	for path != nil && path.dir != "" {
		out = append(out, path.dir)
		path = path.parent
	}
	slices.Reverse(out)
	out = append(out, "A") // End with 'A' as required
	return
}
