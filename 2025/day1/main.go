package main

import (
	"fmt"
	"strconv"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()
	pos := 50
	var pt1, pt2 int
	// fmt.Printf("Start Pos: %d\n", pos)
	for _, line := range lines {
		dir := string(line[0])
		amount, _ := strconv.Atoi(line[1:])

		delta := amount
		if dir == "L" {
			delta = -amount
		}

		oldPos := pos
		raw := oldPos + delta
		pos = ((raw % 100) + 100) % 100

		if pos == 0 {
			pt1++
		}

		if delta > 0 {
			for i := 1; i <= delta; i++ {
				newPos := (oldPos + i) % 100
				if newPos == 0 {
					pt2++
				}
			}
		} else {
			for i := 1; i <= -delta; i++ {
				newPos := ((oldPos-i)%100 + 100) % 100
				if newPos == 0 {
					pt2++
				}
			}
		}

		// fmt.Printf("Step: %s, OldPos: %d, Pos: %d\n", line, oldPos, pos)
	}

	fmt.Printf("Solution Part1: %d\n", pt1)
	fmt.Printf("Solution Part2: %d\n", pt2)
}
