package main

import (
	"fmt"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Node struct {
	plant   string
	visited bool
	x, y    int
}

type Plot struct {
	plant  string
	places []*Node
}

var (
	grid  internal.Grid[*Node]
	plots []*Plot
)

func main() {
	lines := internal.Reader()

	for y, line := range lines {
		for x, c := range line {
			grid.AddSafeToColumn(&Node{plant: string(c), x: x, y: y}, y)
		}
	}

	for y := 0; y < len(grid.Rows); y++ {
		for x := 0; x < len(grid.Rows[y]); x++ {
			node := grid.GetSafeColumn(x, y)
			if node.visited {
				continue
			}
			node.visited = true
			queue := &internal.Queue[*Node]{}
			queue.Enqueue(node)
			plot := &Plot{
				plant: node.plant,
			}
			for len(queue.Elements) > 0 {
				current := queue.Shift()
				plot.places = append(plot.places, current)
				queue.Push(current.getNeighbours()...)
			}
			plots = append(plots, plot)
		}
	}

	var totalPart1, totalPart2 int
	for _, plot := range plots {
		area := plot.getArea()
		per := plot.getPer()
		outer := plot.getInnerWalls()
		totalPart1 += (area * len(per))
		totalPart2 += (area * outer)
	}

	fmt.Printf("Part 1: %d\n", totalPart1)
	fmt.Printf("Part 2: %d\n", totalPart2)

}

// Checks the outer wall only
func (p *Plot) getPer() (per []*Node) {
	for _, place := range p.places {
		// Check each neighbor (up, down, left, right)
		neighbors := []struct{ x, y int }{
			{place.x - 1, place.y},
			{place.x + 1, place.y},
			{place.x, place.y - 1},
			{place.x, place.y + 1},
		}

		for _, neighbor := range neighbors {
			node := &Node{x: neighbor.x, y: neighbor.y}
			if !contains(p.places, node) {
				per = append(per, node)
			}
		}
	}
	return
}

func (p *Plot) getInnerWalls() (total int) {
	for _, place := range p.places {
		if !contains(p.places, &Node{x: place.x + 1, y: place.y}) && !contains(p.places, &Node{x: place.x, y: place.y - 1}) {
			total++
		}
		if !contains(p.places, &Node{x: place.x - 1, y: place.y}) && !contains(p.places, &Node{x: place.x, y: place.y - 1}) {
			total++
		}
		if !contains(p.places, &Node{x: place.x + 1, y: place.y}) && !contains(p.places, &Node{x: place.x, y: place.y + 1}) {
			total++
		}
		if !contains(p.places, &Node{x: place.x - 1, y: place.y}) && !contains(p.places, &Node{x: place.x, y: place.y + 1}) {
			total++
		}
		if contains(p.places, &Node{x: place.x - 1, y: place.y}) && contains(p.places, &Node{x: place.x, y: place.y + 1}) && !contains(p.places, &Node{x: place.x - 1, y: place.y + 1}) {
			total++
		}
		if contains(p.places, &Node{x: place.x + 1, y: place.y}) && contains(p.places, &Node{x: place.x, y: place.y - 1}) && !contains(p.places, &Node{x: place.x + 1, y: place.y - 1}) {
			total++
		}
		if contains(p.places, &Node{x: place.x + 1, y: place.y}) && contains(p.places, &Node{x: place.x, y: place.y + 1}) && !contains(p.places, &Node{x: place.x + 1, y: place.y + 1}) {
			total++
		}
		if contains(p.places, &Node{x: place.x - 1, y: place.y}) && contains(p.places, &Node{x: place.x, y: place.y - 1}) && !contains(p.places, &Node{x: place.x - 1, y: place.y - 1}) {
			total++
		}
	}
	return
}

func (p *Plot) getArea() int {
	return len(p.places)
}

func contains(in []*Node, compare *Node) bool {
	for _, i := range in {
		if i.x == compare.x && i.y == compare.y {
			return true
		}
	}
	return false
}

func (n *Node) getNeighbours() (out []*Node) {
	if n.x-1 >= 0 {
		ne := grid.GetSafeColumn(n.x-1, n.y)
		if !ne.visited && ne.plant == n.plant {
			ne.visited = true
			out = append(out, ne)
		}
	}
	if n.x+1 < len(grid.GetSafeRow(n.y)) {
		ne := grid.GetSafeColumn(n.x+1, n.y)
		if !ne.visited && ne.plant == n.plant {
			ne.visited = true
			out = append(out, ne)
		}
	}
	if n.y-1 >= 0 {
		ne := grid.GetSafeColumn(n.x, n.y-1)
		if !ne.visited && ne.plant == n.plant {
			ne.visited = true
			out = append(out, ne)
		}
	}
	if n.y+1 < len(grid.Rows) {
		ne := grid.GetSafeColumn(n.x, n.y+1)
		if !ne.visited && ne.plant == n.plant {
			ne.visited = true
			out = append(out, ne)
		}
	}
	return
}
