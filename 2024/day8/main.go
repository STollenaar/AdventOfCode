package main

import (
	"fmt"
	"slices"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[string]
}

var (
	grid     = &Grid{}
	antennas = make(map[string][][]int)
)

func main() {
	lines := internal.Reader()

	for y, line := range lines {
		for x, c := range line {
			grid.AddSafeToColumn(string(c), y)
			if string(c) != "." {
				antennas[string(c)] = append(antennas[string(c)], []int{x, y})
			}
		}
	}
	var uniqueNodesPart1, uniqueNodesPart2 []string
	for _, v := range antennas {
		antiNodesP1 := genAntiNodesPart1(v)
		antiNodesP2 := genAntiNodesPart2(v)
		for _, n := range antiNodesP1 {
			ta := fmt.Sprintf("%d-%d", n[0], n[1])
			if !slices.Contains(uniqueNodesPart1, ta) {
				uniqueNodesPart1 = append(uniqueNodesPart1, ta)
			}
		}
		for _, n := range antiNodesP2 {
			ta := fmt.Sprintf("%d-%d", n[0], n[1])
			if !slices.Contains(uniqueNodesPart2, ta) {
				uniqueNodesPart2 = append(uniqueNodesPart2, ta)
			}
		}
	}
	fmt.Printf("Part 1: %d\n", len(uniqueNodesPart1))
	fmt.Printf("Part 2: %d\n", len(uniqueNodesPart2))
}

func genAntiNodesPart1(in [][]int) (antiNodes [][]int) {
	for i := 0; i < len(in); i++ {
		a := in[i]
		for j := i + 1; j < len(in); j++ {
			b := in[j]
			deltaX := a[0] - b[0]
			deltaY := a[1] - b[1]

			dA := []int{a[0] + deltaX, a[1] + deltaY}
			if dA[0] >= 0 && dA[0] < len(grid.Rows[0]) && dA[1] >= 0 && dA[1] < len(grid.Rows) {
				antiNodes = append(antiNodes, dA)
			}
			dB := []int{b[0] - deltaX, b[1] - deltaY}
			if dB[0] >= 0 && dB[0] < len(grid.Rows[0]) && dB[1] >= 0 && dB[1] < len(grid.Rows) {
				antiNodes = append(antiNodes, dB)
			}
		}
	}
	return
}

func genAntiNodesPart2(in [][]int) (antiNodes [][]int) {
	for i := 0; i < len(in); i++ {
		a := in[i]
		for j := i + 1; j < len(in); j++ {
			b := in[j]
			deltaX := a[0] - b[0]
			deltaY := a[1] - b[1]
			tAX := a[0]
			tAY := a[1]
			tBX := b[0]
			tBY := b[1]

			for tAX >= 0 && tAX < len(grid.Rows[0]) && tAY >= 0 && tAY < len(grid.Rows) {
				antiNodes = append(antiNodes, []int{tAX, tAY})
				tAX += deltaX
				tAY += deltaY
			}
			for tBX >= 0 && tBX < len(grid.Rows[0]) && tBY >= 0 && tBY < len(grid.Rows) {
				antiNodes = append(antiNodes, []int{tBX, tBY})
				tBX -= deltaX
				tBY -= deltaY
			}
		}
	}
	return
}
