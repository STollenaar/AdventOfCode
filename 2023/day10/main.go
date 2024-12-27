package main

import (
	"fmt"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Node struct {
	parent *Node
	x, y   int
}

var (
	grid, blownGrid = new(internal.Grid[string]), new(internal.Grid[string])
	nodes           []*Node
	subtract        *internal.Grid[string]
)

func init() {
	blownGrid.SetEmpty(".")
}

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		row := strings.Split(line, "")
		grid.AddRow(row)
	}
	startX, startY := getStart()
	start := &Node{x: startX, y: startY}
	nodes = append(nodes, start)

	neighBours := getValidNeighboursP1(start, startX, startY)
	neighBours = neighBours[:1]
	nodes = append(nodes, neighBours...)
	for len(neighBours) > 0 {
		current := neighBours[0]
		neighBours = neighBours[1:]

		nbs := getValidNeighboursP1(current, current.x, current.y)
		for _, n := range nbs {
			if (n.x == startX && n.y == startY) || visitedNodes(n.x, n.y) != nil {
				continue
			}
			if current.parent == nil {
				neighBours = append(neighBours, n)
			} else if !(current.parent.x == n.x && current.parent.y == n.y) {
				neighBours = append(neighBours, n)
				nodes = append(nodes, n)
			}
		}
	}
	fmt.Printf("Solution part1: %d\n", len(nodes)/2)
	subtract = grid.Copy()

	for y, row := range subtract.Rows {
		for x := range row {
			if visitedNodes(x, y) == nil {
				subtract.SetUnsafeColumn(".", x, y)
			}
		}
	}
	fillBlowngrid()
	blownGrid.SetUnsafeColumn("O", 0, 0)

	neighBours = []*Node{{x: 0, y: 0}}
	for len(neighBours) > 0 {
		current := neighBours[0]
		neighBours = neighBours[1:]
		nbs := getValidNeighboursP2(current.x, current.y)
		for _, n := range nbs {
			blownGrid.SetUnsafeColumn("O", n.x, n.y)
		}
		neighBours = append(neighBours, nbs...)
	}

	for y, row := range blownGrid.Rows {
		for x := range row {
			fmt.Print(blownGrid.Rows[y][x])
		}
		fmt.Println()
	}
	var nest int
	for y:=0;y<len(blownGrid.Rows);y+=2{
		rows := blownGrid.GetSafeRow(y)
		for x:=0;x<len(rows);x+=2 {
			if rows[x] == "."{
				nest++
			}
		}
	}
	fmt.Printf("Solution part2: %d\n", nest)
}

func getStart() (int, int) {
	for y, row := range grid.Rows {
		for x, c := range row {
			if c == "S" {
				return x, y
			}
		}
	}
	return -1, -1
}

func getValidNeighboursP1(current *Node, sX, sY int) (nbs []*Node) {
	if sX+1 < len(grid.Rows[0]) && connected(sX, sY, sX+1, sY) {
		nbs = append(nbs, &Node{parent: current, x: sX + 1, y: sY})
	}
	if sY+1 < len(grid.Rows) && connected(sX, sY, sX, sY+1) {
		nbs = append(nbs, &Node{parent: current, x: sX, y: sY + 1})
	}
	if sX-1 >= 0 && connected(sX, sY, sX-1, sY) {
		nbs = append(nbs, &Node{parent: current, x: sX - 1, y: sY})
	}
	if sY-1 >= 0 && connected(sX, sY, sX, sY-1) {
		nbs = append(nbs, &Node{parent: current, x: sX, y: sY - 1})
	}
	return
}

func connected(sX, sY, nX, nY int) bool {
	token := grid.Rows[sY][sX]
	switch token {
	case "-":
		return (sY == nY && nX == sX+1) || (sY == nY && nX == sX-1)
	case "L":
		return (sY == nY && nX == sX+1) || (sY-1 == nY && nX == sX)
	case "F":
		return (sY == nY && nX == sX+1) || (sY+1 == nY && nX == sX)
	case "J":
		return (sY == nY && nX == sX-1) || (sY-1 == nY && nX == sX)
	case "|":
		return (sY+1 == nY && nX == sX) || (sY-1 == nY && nX == sX)
	case "7":
		return (sY == nY && nX == sX-1) || (sY+1 == nY && nX == sX)
	case "S":
		return true
	default:
		return false
	}
}

func getValidNeighboursP2(x, y int) (nbs []*Node) {
	if y-1 >= 0 {
		rows := blownGrid.Rows[y-1]
		if x-1 >= 0 && (rows[x-1] == ".") {
			nbs = append(nbs, &Node{y: y - 1, x: x - 1})
		}
		if x+1 < len(rows) && (rows[x+1] == ".") {
			nbs = append(nbs, &Node{y: y - 1, x: x + 1})
		}
		if rows[x] == "." {
			nbs = append(nbs, &Node{y: y - 1, x: x})
		}
	}
	if y+1 < len(blownGrid.Rows) {
		rows := blownGrid.Rows[y+1]
		if x-1 >= 0 && (rows[x-1] == ".") {
			nbs = append(nbs, &Node{y: y + 1, x: x - 1})
		}
		if x+1 < len(rows) && (rows[x+1] == ".") {
			nbs = append(nbs, &Node{y: y + 1, x: x + 1})
		}
		if rows[x] == "." {
			nbs = append(nbs, &Node{y: y + 1, x: x})
		}
	}
	rows := blownGrid.Rows[y]
	if x-1 >= 0 && (rows[x-1] == ".") {
		nbs = append(nbs, &Node{y: y, x: x - 1})
	}
	if x+1 < len(rows) && (rows[x+1] == ".") {
		nbs = append(nbs, &Node{y: y, x: x + 1})
	}
	if rows[x] == "." {
		nbs = append(nbs, &Node{y: y, x: x})
	}
	return
}

func fillBlowngrid() {
	for y, row := range subtract.Rows {
		for x, c := range row {
			bX, bY := x*2, y*2
			switch c {
			case "|":
				blownGrid.SetSafeColumn("|", bX, bY)
				blownGrid.SetSafeColumn(".", bX+1, bY)
				blownGrid.SetSafeColumn("|", bX, bY+1)
				blownGrid.SetSafeColumn(".", bX+1, bY+1)
			case "-":
				blownGrid.SetSafeColumn("-", bX, bY)
				blownGrid.SetSafeColumn("-", bX+1, bY)
				blownGrid.SetSafeColumn(".", bX, bY+1)
				blownGrid.SetSafeColumn(".", bX+1, bY+1)
			case "L":
				blownGrid.SetSafeColumn("L", bX, bY)
				blownGrid.SetSafeColumn("-", bX+1, bY)
				blownGrid.SetSafeColumn(".", bX, bY+1)
				blownGrid.SetSafeColumn(".", bX+1, bY+1)
			case "J":
				blownGrid.SetSafeColumn("J", bX, bY)
				blownGrid.SetSafeColumn(".", bX+1, bY)
				blownGrid.SetSafeColumn(".", bX, bY+1)
				blownGrid.SetSafeColumn(".", bX+1, bY+1)
			case "F":
				blownGrid.SetSafeColumn("F", bX, bY)
				blownGrid.SetSafeColumn("-", bX+1, bY)
				blownGrid.SetSafeColumn("|", bX, bY+1)
				blownGrid.SetSafeColumn(".", bX+1, bY+1)
			case "7":
				blownGrid.SetSafeColumn("7", bX, bY)
				blownGrid.SetSafeColumn(".", bX+1, bY)
				blownGrid.SetSafeColumn("|", bX, bY+1)
				blownGrid.SetSafeColumn(".", bX+1, bY+1)
			case "S":
				blownGrid.SetSafeColumn("S", bX, bY)
				blownGrid.SetSafeColumn("S", bX+1, bY)
				blownGrid.SetSafeColumn("S", bX, bY+1)
				blownGrid.SetSafeColumn("S", bX+1, bY+1)
			default:
				blownGrid.SetSafeColumn(".", bX, bY)
				blownGrid.SetSafeColumn(".", bX+1, bY)
				blownGrid.SetSafeColumn(".", bX, bY+1)
				blownGrid.SetSafeColumn(".", bX+1, bY+1)
			}
		}
	}
}

func visitedNodes(nX, nY int) *Node {
	for _, c := range nodes {
		if c.x == nX && c.y == nY {
			return c
		}
	}
	return nil
}
