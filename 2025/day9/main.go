package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Node struct {
	X, Y int
}

type Edge struct {
	A, B        *Node
	BoundedArea int
	Area        int
}

type BoundingBox struct {
	BottomLeft *Node
	TopRight   *Node
}

var nodes []*Node

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		nodes = append(nodes, &Node{
			X: x,
			Y: y,
		})
	}

	areas := generateAreas(nodes)

	fmt.Printf("Part 1: %d\n", areas[0].Area)

	boundedArea := generateBoundedArea(nodes)

	fmt.Printf("Part 2: %d\n", boundedArea)

}

func generateAreas(nodes []*Node) (edges []*Edge) {
	n := len(nodes)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {

			a := nodes[i]
			b := nodes[j]

			unboundedArea := area(a, b)

			edges = append(edges, &Edge{
				A:    a,
				B:    b,
				Area: unboundedArea,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Area > edges[j].Area
	})
	return
}

func generateBoundedArea(nodes []*Node) (result int) {
	var edges []*Edge
	first, last := nodes[0], nodes[len(nodes)-1]

	for i := 0; i < len(nodes)-1; i++ {
		edges = append(edges, &Edge{
			A: nodes[i],
			B: nodes[i+1],
		})
	}
	edges = append(edges, &Edge{
		A: last,
		B: first,
	})

	intersections := func(minX, minY, maxX, maxY int) bool {
		for _, inter := range edges {
			iMinX, iMaxX := minMax(inter.A.X, inter.B.X)
			iMinY, iMaxY := minMax(inter.A.Y, inter.B.Y)
			if minX < iMaxX && maxX > iMinX && minY < iMaxY && maxY > iMinY {
				return true
			}
		}
		return false
	}

	manhattanDistance := func(a, b *Node) int {
		return abs(a.X, b.X) + abs(a.Y, b.Y)
	}

	for fTIdx := 0; fTIdx < len(nodes)-1; fTIdx++ {
		for tTIdx := fTIdx; tTIdx < len(nodes); tTIdx++ {
			fromTile := nodes[fTIdx]
			toTile := nodes[tTIdx]
			minX, maxX := minMax(fromTile.X, toTile.X)
			minY, maxY := minMax(fromTile.Y, toTile.Y)

			manhattanDistance := manhattanDistance(fromTile, toTile)
			if manhattanDistance*manhattanDistance > result {
				if !intersections(minX, minY, maxX, maxY) {
					area := area(fromTile, toTile)
					if area > result {
						result = area
					}
				}
			}
		}
	}
	return
}

func area(a, b *Node) int {
	return (abs(a.X, b.X) + 1) * (abs(a.Y, b.Y) + 1)
}

func abs(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
