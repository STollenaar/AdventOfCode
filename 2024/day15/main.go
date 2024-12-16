package main

import (
	"fmt"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Point struct {
	x, y int
	kind string
}

type Grid struct {
	internal.Grid[*Point]
}

type Queue struct {
	internal.Queue[*Point]
}

var (
	gridP1, gridP2 Grid

	moves []string

	robotP1, robotP2 *Point
	boxesP1, boxesP2 []*Point
	walls            []*Point
)

func main() {
	lines := internal.Reader()
	loadGrid := true
	for y, line := range lines {
		if line == "" {
			loadGrid = false
			continue
		}
		if loadGrid {
			var i int
			for x, c := range line {
				p := &Point{x: x, y: y, kind: string(c)}
				gridP1.AddSafeToColumn(p, y)
				switch p.kind {
				case "@":
					robotP1 = p
					p2 := &Point{kind: string(c), x: x + i, y: y}
					robotP2 = p2
					p2w := &Point{kind: ".", x: x + i + 1, y: y}
					gridP2.SetSafeColumn(p2, x+i, y)
					gridP2.SetSafeColumn(p2w, x+i+1, y)
				case "O":
					boxesP1 = append(boxesP1, p)
					p2 := &Point{kind: "[", x: x + i, y: y}
					boxesP2 = append(boxesP2, p2)
					p2w := &Point{kind: "]", x: x + i + 1, y: y}
					gridP2.SetSafeColumn(p2, x+i, y)
					gridP2.SetSafeColumn(p2w, x+i+1, y)
				case "#":
					walls = append(walls, p)
					p2 := &Point{kind: string(c), x: x + i, y: y}
					p2w := &Point{kind: "#", x: x + i + 1, y: y}
					gridP2.SetSafeColumn(p2, x+i, y)
					gridP2.SetSafeColumn(p2w, x+i+1, y)
				default:
					p2 := &Point{kind: string(c), x: x + i, y: y}
					p2w := &Point{kind: ".", x: x + i + 1, y: y}
					gridP2.SetSafeColumn(p2, x+i, y)
					gridP2.SetSafeColumn(p2w, x+i+1, y)
				}
				i++
			}
		} else {
			for _, c := range line {
				moves = append(moves, string(c))
			}
		}
	}
	solvePart1()
	solvePart2()
}

func (g *Grid) print() {
	for _, row := range g.Rows {
		for _, c := range row {
			fmt.Print(c.kind)
		}
		fmt.Println()
	}
}

func solvePart1() {
	for _, move := range moves {
		// fmt.Printf("Move: %d/%d\n", i, len(moves))
		var dx, dy int
		switch move {
		case "^":
			dx, dy = 0, -1
		case ">":
			dx, dy = 1, 0
		case "v":
			dx, dy = 0, 1
		case "<":
			dx, dy = -1, 0
		}
		if point := gridP1.GetSafeColumn(robotP1.x+dx, robotP1.y+dy); point.kind == "." {
			gridP1.SetSafeColumn(&Point{kind: ".", x: robotP1.x, y: robotP1.y}, robotP1.x, robotP1.y)
			robotP1.x, robotP1.y = robotP1.x+dx, robotP1.y+dy
			gridP1.SetSafeColumn(robotP1, robotP1.x, robotP1.y)

		} else if point.kind == "O" {
			nextP := gridP1.GetSafeColumn(point.x+dx, point.y+dy)
			mult := 1
			for nextP.kind != "." && nextP.kind != "#" {
				nextP = gridP1.GetSafeColumn(nextP.x+dx, nextP.y+dy)
				mult++
			}
			if nextP.kind == "." {
				currentR := &Point{kind: ".", x: robotP1.x, y: robotP1.y}
				robotP1.x = point.x
				robotP1.y = point.y

				prevNextP := gridP1.GetSafeColumn(nextP.x-dx, nextP.y-dy)
				t := &Point{kind: ".", x: prevNextP.x, y: prevNextP.y}
				prevNextP.x = nextP.x
				prevNextP.y = nextP.y
				gridP1.SetSafeColumn(prevNextP, prevNextP.x, prevNextP.y)
				gridP1.SetSafeColumn(t, t.x, t.y)

				for ; mult > 1; mult-- {
					nextP = t
					prevNextP = gridP1.GetSafeColumn(nextP.x-dx, nextP.y-dy)
					t = &Point{kind: ".", x: prevNextP.x, y: prevNextP.y}
					prevNextP.x = nextP.x
					prevNextP.y = nextP.y
					gridP1.SetSafeColumn(prevNextP, prevNextP.x, prevNextP.y)
					gridP1.SetSafeColumn(t, t.x, t.y)
				}
				gridP1.SetSafeColumn(robotP1, robotP1.x, robotP1.y)
				gridP1.SetSafeColumn(currentR, currentR.x, currentR.y)
			}
		}
	}
	// gridP1.print()
	var totalPart1 int
	for _, box := range boxesP1 {
		totalPart1 += box.y*100 + box.x
	}
	fmt.Printf("Part 1: %d\n", totalPart1)
}

func solvePart2() {
	for _, move := range moves {
		// fmt.Printf("Move: %d/%d\n", i, len(moves))
		// Define direction deltas
			var dx, dy int
		switch move {
		case "^":
			dx, dy = 0, -1
		case ">":
			dx, dy = 1, 0
		case "v":
			dx, dy = 0, 1
		case "<":
			dx, dy = -1, 0
		}

		// Handle movement to an empty cell
		if point := gridP2.GetSafeColumn(robotP2.x+dx, robotP2.y+dy); point.kind == "." {
			gridP2.SetSafeColumn(&Point{kind: ".", x: robotP2.x, y: robotP2.y}, robotP2.x, robotP2.y)
			robotP2.x, robotP2.y = robotP2.x+dx, robotP2.y+dy
			gridP2.SetSafeColumn(robotP2, robotP2.x, robotP2.y)

		} else if point.kind == "[" || point.kind == "]" { // Handle box
			var connectedBoxes []*Point
			if dy != 0 { // Vertical movement
				connectedBoxes = getFullBoxVertical(point, dy)
			} else if dx != 0 { // Horizontal movement
				connectedBoxes = getConnectedBoxesSameRow(point)
			}

			// Validate movement for all boxes
			canMove := true
			for _, box := range connectedBoxes {
				nextPos := gridP2.GetSafeColumn(box.x+dx, box.y+dy)
				if nextPos.kind == "#" {
					canMove = false
					break
				}
			}

			// Move all boxes if valid
			if canMove {
				// Clear current positions
				for _, box := range connectedBoxes {
					gridP2.SetSafeColumn(&Point{kind: ".", x: box.x, y: box.y}, box.x, box.y)
				}

				// Update to new positions
				for i := range connectedBoxes {
					connectedBoxes[i].x += dx
					connectedBoxes[i].y += dy
					gridP2.SetSafeColumn(connectedBoxes[i], connectedBoxes[i].x, connectedBoxes[i].y)
				}

				// Update robot position
				gridP2.SetSafeColumn(&Point{kind: ".", x: robotP2.x, y: robotP2.y}, robotP2.x, robotP2.y)
				robotP2.x, robotP2.y = robotP2.x+dx, robotP2.y+dy
				gridP2.SetSafeColumn(robotP2, robotP2.x, robotP2.y)
			}
		}
		// gridP2.print()
	}

	// Calculate total score
	var totalPart2 int
	for _, box := range boxesP2 {
		if box.kind == "[" {
			totalPart2 += box.y*100 + box.x
		}
	}
	fmt.Printf("Part 2: %d\n", totalPart2)
}

// Helper function to find all connected boxes on the same row
func getConnectedBoxesSameRow(start *Point) []*Point {
	visited := make(map[*Point]bool)
	var result []*Point

	var dfs func(*Point)
	dfs = func(p *Point) {
		if visited[p] || p.y != start.y { // Ensure we stay on the same row
			return
		}
		visited[p] = true
		result = append(result, p)

		// Check adjacent points for connected boxes on the same row
		for _, dir := range [][2]int{{-1, 0}, {1, 0}} { // Only horizontal directions
			adj := gridP2.GetSafeColumn(p.x+dir[0], p.y+dir[1])
			if adj.kind == "[" || adj.kind == "]" {
				dfs(adj)
			}
		}
	}

	dfs(start)
	return result
}

func getFullBoxVertical(start *Point, dy int) []*Point {
	visited := make(map[string]bool) // Track visited points using their "x,y" key
	queue := &Queue{}

	queue.EqualFunction = func(input ...*Point) bool {
		for _, p := range queue.Elements {
			if p == input[0] {
				return true
			}
		}
		return false
	}
	var result []*Point

	// Helper function to generate a unique key for a point
	key := func(p *Point) string {
		return fmt.Sprintf("%d,%d", p.x, p.y)
	}

	// Add a point to the result and mark it as visited
	addPoint := func(p *Point) {
		if !visited[key(p)] {
			visited[key(p)] = true
			queue.PushUnique(p)
			result = append(result, p)
		}
	}
	queue.Push(start)
	if start.kind == "[" {
		queue.Push(gridP2.GetSafeColumn(start.x+1, start.y))
	} else {
		queue.Push(gridP2.GetSafeColumn(start.x-1, start.y))
	}

	for len(queue.Elements) > 0 {
		current := queue.Pop()

		point := gridP2.GetSafeColumn(current.x, current.y+dy)

		// Stop traversal at walls or grid boundaries
		if point.kind == "#" {
			break
		}

		// If a box part (`[` or `]`) is found, include the whole box
		if point.kind == "[" || point.kind == "]" {
			addPoint(point)

			// Include the other half of the box
			if point.kind == "[" {
				right := gridP2.GetSafeColumn(point.x+1, point.y)
				addPoint(right)
			} else {
				left := gridP2.GetSafeColumn(point.x-1, point.y)
				addPoint(left)
			}
		}
	}

	// Include the starting point and its horizontal box part (if any)
	addPoint(start)
	if start.kind == "[" {
		right := gridP2.GetSafeColumn(start.x+1, start.y)
		addPoint(right)
	} else {
		left := gridP2.GetSafeColumn(start.x-1, start.y)
		addPoint(left)
	}

	return result
}
