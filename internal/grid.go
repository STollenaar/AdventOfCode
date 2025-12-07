package internal

import (
	"fmt"
	"slices"
)

type Row[T comparable] []T

type Grid[T comparable] struct {
	Rows  []Row[T]
	empty T
}

// Adds object to end of the slice
func (g *Grid[T]) AddUnsafeToColumn(object T, row int) bool {
	g.Rows[row] = append(g.Rows[row], object)
	return true
}

// Adds object to beginning of the slice
func (g *Grid[T]) ShiftUnsafeToColumn(object T, row int) bool {
	g.Rows[row] = append([]T{object}, g.Rows[row]...)
	return true
}

// Adds objects safely to end of the slice
func (g *Grid[T]) AddSafeToColumn(object T, row int) bool {
	g.GetSafeRow(row)
	g.AddUnsafeToColumn(object, row)
	return true
}

// Adds objects safely to beginning of the slice
func (g *Grid[T]) ShiftSafeToColumn(object T, row int) bool {
	g.GetSafeRow(row)
	g.ShiftUnsafeToColumn(object, row)
	return true
}

func (g *Grid[T]) SetUnsafeColumn(object T, x, y int) bool {
	g.Rows[y][x] = object
	return true
}

func (g *Grid[T]) SetSafeColumn(object T, x, y int) bool {
	g.GetSafeColumn(x, y)
	g.GetSafeRow(y)[x] = object
	return true
}

// Adds row to end of the slice
func (g *Grid[T]) AddRow(row Row[T]) bool {
	g.Rows = append(g.Rows, row)
	return true
}

// Adds row to beginning of the slice
func (g *Grid[T]) ShiftRow(row Row[T]) bool {
	g.Rows = append([]Row[T]{row}, g.Rows...)
	return true
}

func (g *Grid[T]) GetUnsafeRow(y int) Row[T] {
	return g.Rows[y]
}

// Adds rows if needed, or shifts once and returns the new row
func (g *Grid[T]) GetSafeRow(y int) Row[T] {
	if y < 0 {
		g.ShiftRow(Row[T]{})
		return g.Rows[0]
	}
	for len(g.Rows) <= y {
		g.AddRow(Row[T]{})
	}
	return g.Rows[y]
}

func (g *Grid[T]) GetUnsafeColumn(x, y int) T {
	return g.GetUnsafeRow(y)[x]
}

// Tries to safely get a column, will apply GetSafeRow function. Throws error if the x value is not valid
func (g *Grid[T]) GetSafeColumn(x, y int) T {
	for len(g.GetSafeRow(y)) <= x {
		g.AddUnsafeToColumn(g.empty, y)
	}
	return g.GetSafeRow(y)[x]
}

// Getting the max Y of a non-empty given column
func (g *Grid[T]) GetHeight(x int) int {
	for y := range g.Rows {
		if g.GetSafeColumn(x, y) != g.empty {
			return y
		}
	}
	return len(g.Rows) - 1
}

// Set the empty value
func (g *Grid[T]) SetEmpty(in T) {
	g.empty = in
}

func (g *Grid[T]) Print() {
	for _, row := range g.Rows {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func (g *Grid[T]) Copy() *Grid[T] {
	var duplicate Grid[T]

	for i := range g.Rows {
		row := make([]T, len(g.Rows[i]))
		copy(row, g.Rows[i])
		duplicate.AddRow(row)
	}
	return &duplicate
}

func (g *Grid[T]) Rotate90() *Grid[T] {
	rotated := Grid[T]{empty: g.empty}
	for x := 0; x < len(g.Rows[0]); x++ {
		newRow := Row[T]{}
		for y := len(g.Rows) - 1; y >= 0; y-- {
			newRow = append(newRow, g.GetSafeColumn(x, y))
		}
		rotated.AddRow(newRow)
	}
	return &rotated
}

func (g *Grid[T]) FlipHorizontal() *Grid[T] {
	flipped := g.Copy()
	for y := range flipped.Rows {
		slices.Reverse(flipped.Rows[y])
	}
	return flipped
}

func (g *Grid[T]) FlipVertical() *Grid[T] {
	flipped := Grid[T]{empty: g.empty}
	for y := len(g.Rows) - 1; y >= 0; y-- {
		flipped.AddRow(g.Rows[y])
	}
	return &flipped
}

func (g *Grid[T]) FloodFill(x, y int, target, replacement T) {
	if g.GetSafeColumn(x, y) != target || target == replacement {
		return
	}
	var stack []Point[int]
	stack = append(stack, Point[int]{X: x, Y: y})

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if g.GetSafeColumn(p.X, p.Y) == target {
			g.SetSafeColumn(replacement, p.X, p.Y)

			for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				nx, ny := p.X+d[0], p.Y+d[1]
				if nx >= 0 && ny >= 0 {
					stack = append(stack, Point[int]{X: nx, Y: ny})
				}
			}
		}
	}
}

func (g *Grid[T]) BFS(startX, startY, endX, endY int, walkable T) int {
	type Node struct {
		x, y, dist int
	}

	queue := []Node{{x: startX, y: startY, dist: 0}}
	visited := make(map[string]bool)
	key := func(x, y int) string { return fmt.Sprintf("%d,%d", x, y) }

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.x == endX && node.y == endY {
			return node.dist
		}

		if visited[key(node.x, node.y)] {
			continue
		}
		visited[key(node.x, node.y)] = true

		for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			nx, ny := node.x+d[0], node.y+d[1]
			if nx >= 0 && ny >= 0 && g.GetSafeColumn(nx, ny) == walkable {
				queue = append(queue, Node{x: nx, y: ny, dist: node.dist + 1})
			}
		}
	}
	return -1
}

func (g *Grid[T]) DrawLine(x1, y1, x2, y2 int, symbol T) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := -1
	if x1 < x2 {
		sx = 1
	}
	sy := -1
	if y1 < y2 {
		sy = 1
	}
	err := dx - dy

	for {
		g.SetSafeColumn(symbol, x1, y1)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}
func (g *Grid[T]) ShiftRowsDown() {
	last := g.Rows[len(g.Rows)-1]
	copy(g.Rows[1:], g.Rows[:len(g.Rows)-1])
	g.Rows[0] = last
}

func (g *Grid[T]) ShiftColumnsRight() {
	for y := range g.Rows {
		last := g.Rows[y][len(g.Rows[y])-1]
		copy(g.Rows[y][1:], g.Rows[y][:len(g.Rows[y])-1])
		g.Rows[y][0] = last
	}
}

// GetAdjacent get the direct 4 adjacent tiles
func (g *Grid[T]) GetAdjacent(x, y int) (adj map[string]T) {
	adj = make(map[string]T)
	// Check top-left
	if y-1 >= 0 {
		rows := g.Rows[y-1]
		adj["u"] = rows[x]
	}
	// Check left and right
	rows := g.Rows[y]
	if x-1 >= 0 {
		adj["l"] = rows[x-1]
	}
	if x+1 < len(rows) {
		adj["r"] = rows[x+1]
	}
	// Check bottom-left, bottom, bottom-right
	if y+1 < len(g.Rows) {
		adj["b"] = g.Rows[y+1][x]
	}
	return
}

// GetNeighbours get the octogonal tiles
func (g *Grid[T]) GetNeighbours(x, y int) (adj map[string]T) {
	adj = make(map[string]T)
	// Check top-left
	if y-1 >= 0 {
		rows := g.Rows[y-1]
		adj["u"] = rows[x]
		if x-1 >= 0 {
			adj["ul"] = rows[x-1]
		}
		if x+1 < len(rows) {
			adj["ur"] = rows[x+1]
		}
	}
	// Check left and right
	rows := g.Rows[y]
	if x-1 >= 0 {
		adj["l"] = rows[x-1]
	}
	if x+1 < len(rows) {
		adj["r"] = rows[x+1]
	}
	// Check bottom-left, bottom, bottom-right
	if y+1 < len(g.Rows) {
		rows := g.Rows[y+1]
		adj["b"] = rows[x]
		if x-1 >= 0 {
			adj["bl"] = rows[x-1]
		}
		if x+1 < len(rows) {
			adj["br"] = rows[x+1]
		}
	}
	return
}
