package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[string]
}

type Galaxy struct {
	x, y, trueX, trueY int
}

var (
	grid = new(Grid)
)

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		var rows internal.Row[string]
		rows = append(rows, strings.Split(line, "")...)
		grid.AddRow(rows)
	}
	// expandGalaxy(1)

	for _, row := range grid.Rows {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println()
	}
	var galaxies []*Galaxy
	for y, row := range grid.Rows {
		for x, c := range row {
			if c == "#" {
				g := &Galaxy{x: x, y: y}
				getTrueGalaxyCoords(2, g)
				galaxies = append(galaxies, g)
			}
		}
	}
	var total int
	for i := 0; i < len(galaxies); i++ {
		galaxy := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			oGalaxy := galaxies[j]
			distance := int(math.Abs(float64(galaxy.trueX-oGalaxy.trueX)) + math.Abs(float64(galaxy.trueY-oGalaxy.trueY)))
			total += distance
		}
	}
	fmt.Printf("Solution for part1: %d\n", total)
	for _, g := range galaxies {
		getTrueGalaxyCoords(1000000, g)
	}

	total = 0
	for i := 0; i < len(galaxies); i++ {
		galaxy := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			oGalaxy := galaxies[j]
			distance := int(math.Abs(float64(galaxy.trueX-oGalaxy.trueX)) + math.Abs(float64(galaxy.trueY-oGalaxy.trueY)))
			total += distance
		}
	}
	fmt.Printf("Solution for part2: %d\n", total)
}

func getTrueGalaxyCoords(scalar int, galaxy *Galaxy) {
	var emptyRows, emptyColumns int

	for y, row := range grid.Rows {
		if !strings.Contains(strings.Join(row, ""), "#") && y < galaxy.y {
			emptyRows++
		}
	}
	row := grid.Rows[0]
	for x := 0; x < len(row) && x < galaxy.x; x++ {
		if row[x] == "." {
			empty := true
			for _, row := range grid.Rows {
				if row[x] != "." {
					empty = false
					break
				}
			}
			if empty {
				emptyColumns++
			}
		}
	}

	galaxy.trueX = galaxy.x + (emptyColumns * (scalar - 1))
	galaxy.trueY = galaxy.y + (emptyRows * (scalar - 1))
}
