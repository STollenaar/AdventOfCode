package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[int]
}

type Queue struct {
	internal.Queue[*Node]
}

type Node struct {
	x, y, distance, totalH int
	parent                 *Node
}

var (
	grid        Grid
	queue       Queue
	tmpGrid     [][]string
	nodes       [][]bool
	badDistance int
)

func init() {
	queue.SortFunction = sort
	queue.EqualFunction = contains
}

func main() {
	lines := internal.Reader()

	for r, line := range lines {
		var row []string
		for _, c := range strings.Split(line, "") {
			n, _ := strconv.Atoi(c)
			grid.AddSafeToColumn(n, r)
			row = append(row, ".")
		}
		tmpGrid = append(tmpGrid, row)
		nodes = append(nodes, make([]bool, len(row)))
	}
	start := &Node{x: 0, y: 0, distance: 0}
	queue.Enqueue(start)
	badDistance = getBadDistance()

	var current *Node
	var iterations int
	for len(queue.Elements) > 0 {
		current = queue.Shift()
		nodes[current.y][current.x] = true
		if current.x == len(grid.Rows[0])-1 && current.y == len(grid.Rows)-1 {
			break
		}
		queue.PushUnique(getNeigbours(current)...)
		queue.Sort()
		// if len(queue.Elements) > 1000 {
		// 	queue.Elements = queue.Elements[:1000]
		// }
		iterations++
	}
	// var total int
	fmt.Printf("Solution to part 1: %d\n", current.totalH)
}

func getNeigbours(current *Node) (res []*Node) {
	var buffer []*Node

	if current.y+1 < len(grid.Rows) {
		totalH := current.totalH + grid.Rows[current.y+1][current.x]
		buffer = append(buffer, &Node{x: current.x, y: current.y + 1, totalH: totalH, parent: current})
	}
	if current.y-1 >= 0 {
		totalH := current.totalH + grid.Rows[current.y-1][current.x]
		buffer = append(buffer, &Node{x: current.x, y: current.y - 1, totalH: totalH, parent: current})
	}
	if current.x+1 < len(grid.Rows[0]) {
		totalH := current.totalH + grid.Rows[current.y][current.x+1]
		buffer = append(buffer, &Node{x: current.x + 1, y: current.y, totalH: totalH, parent: current})
	}
	if current.x-1 >= 0 {
		totalH := current.totalH + grid.Rows[current.y][current.x-1]
		buffer = append(buffer, &Node{x: current.x - 1, y: current.y, totalH: totalH, parent: current})
	}

	if current.parent == nil {
		return buffer
	}
	for _, b := range buffer {
		if visited(b) || b.totalH >= badDistance {
			continue
		}
		tmp := current.parent
		if tmp.parent != nil && tmp.parent.parent != nil {
			tmpP := tmp.parent
			tmpPP := tmpP.parent
			if b.x == current.x && current.x == tmp.x && tmp.x == tmpP.x && tmpP.x == tmpPP.x {
				continue
			}
			if b.y == current.y && current.y == tmp.y && tmp.y == tmpP.y && tmpP.y == tmpPP.y {
				continue
			}
			b.distance = getDistance(b)
			res = append(res, b)
		} else {
			b.distance = getDistance(b)
			res = append(res, b)
		}
	}
	return
}

func getDistance(node *Node) int {
	abX := math.Abs(float64((len(grid.Rows[0]) - 1) - node.x))
	abY := math.Abs(float64((len(grid.Rows) - 1) - node.y))
	return int(abX + abY)
}

func getBadDistance() (total int) {
	for x := 1; x < len(grid.Rows[0]); x++ {
		total += grid.Rows[0][x]
	}
	for y := 1; y < len(grid.Rows); y++ {
		total += grid.Rows[y][len(grid.Rows[0])-1]
	}
	return
}

func visited(node *Node) bool {
	return nodes[node.y][node.x]
}

func sort(a, b int) bool {
	return queue.Elements[a].totalH+queue.Elements[a].distance < queue.Elements[b].totalH+queue.Elements[b].distance
}

func contains(nodes ...*Node) bool {
	for _, n := range nodes {
		for _, q := range queue.Elements {
			if q.x == n.x && q.y == n.y {
				// if q.totalH >= n.totalH {
				return true
				// }
			}
		}
	}
	return false
}

func printGrid() {
	for _, r := range nodes {
		for _, c := range r {
			if c {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
