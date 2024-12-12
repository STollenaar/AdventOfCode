package main

import (
	"fmt"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[*Node]
}

type Node struct {
	plant   string
	visited bool
	x, y    int
}

type Plot struct {
	plant  string
	places []*Node
}

type Queue struct {
	internal.Queue[*Node]
}

var (
	grid  Grid
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
			queue := &Queue{}
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
		sides := plot.getSides()
		totalPart1 += (area * len(per))
		totalPart2 += (area * sides)
	}

	fmt.Printf("Part 1: %d\n", totalPart1)
	fmt.Printf("Part 2: %d\n", totalPart2)

}

func (p *Plot) isNeighbor(x, y int) bool {
	for _, place := range p.places {
		if place.x-1 == x && place.y == y {
			return true
		} else if place.x+1 == x && place.y == y {
			return true
		} else if place.x == x && place.y-1 == y {
			return true
		} else if place.x == x && place.y+1 == y {
			return true
		}
	}
	return false
}

func (p *Plot) getPer() (per []*Node) {
	for _, place := range p.places {
		if !contains(p.places, &Node{x: place.x - 1, y: place.y}) {
			per = append(per, &Node{x: place.x - 1, y: place.y})
		}
		if !contains(p.places, &Node{x: place.x + 1, y: place.y}) {
			per = append(per, &Node{x: place.x + 1, y: place.y})
		}
		if !contains(p.places, &Node{x: place.x, y: place.y - 1}) {
			per = append(per, &Node{x: place.x, y: place.y - 1})
		}
		if !contains(p.places, &Node{x: place.x, y: place.y + 1}) {
			per = append(per, &Node{x: place.x, y: place.y + 1})
		}
	}
	return
}

func (p *Plot) getSides() (total int) {
	per := p.places
	x, y := per[0].x, per[0].y
	dx, dy := 1, 0
	
	var steps int
	plotSize := len(p.getPer())
	for steps != plotSize {

		if contains(per, &Node{x: x + dx, y: y + dy}) {
			if dx == 1 && contains(per, &Node{x: x + dx, y: y - 1}) {
				total++
				x, y = x+dx, y+dy
				dx, dy = 0, -1
			} else if dx == -1 && contains(per, &Node{x: x + dx, y: y + 1}) {
				total++
				x, y = x+dx, y+dy
				dx, dy = 0, 1
			} else if dy == 1 && contains(per, &Node{x: x + 1, y: y + dy}) {
				total++
				x, y = x+dx, y+dy
				dx, dy = 1, 0
			} else if dy == -1 && contains(per, &Node{x: x - 1, y: y + dy}) {
				total++
				x, y = x+dx, y+dy
				dx, dy = -1, 0
			}
			x, y = x+dx, y+dy
		} else {
			total++
			if dx == 1 {
				dx, dy = 0, 1
			} else if dx == -1 {
				dx, dy = 0, -1
			} else if dy == 1 {
				dx, dy = -1, 0
			} else {
				dx, dy = 1, 0
			}
		}
		steps++
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
