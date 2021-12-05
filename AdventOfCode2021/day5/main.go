package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	gridFlat := make(map[int]map[int]int)
	gridAll := make(map[int]map[int]int)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		positions := strings.Split(line, "->")

		pos1 := strings.Split(strings.TrimSpace(positions[0]), ",")
		x1, _ := strconv.Atoi(pos1[0])
		y1, _ := strconv.Atoi(pos1[1])

		pos2 := strings.Split(strings.TrimSpace(positions[1]), ",")
		x2, _ := strconv.Atoi(pos2[0])
		y2, _ := strconv.Atoi(pos2[1])

		// Horizontal or Vertical vent
		if (x1 == x2 && y1 != y2) || x1 != x2 && y1 == y2 {
			if x2 < x1 {
				tmp := x1
				x1 = x2
				x2 = tmp
			}

			if y2 < y1 {
				tmp := y1
				y1 = y2
				y2 = tmp
			}
			for x := x1; x <= x2; x++ {
				if gridFlat[x] == nil {
					gridFlat[x] = make(map[int]int)
				}
				if gridAll[x] == nil {
					gridAll[x] = make(map[int]int)
				}
				for y := y1; y <= y2; y++ {
					gridFlat[x][y]++
					gridAll[x][y]++
				}
			}
		} else {
			// Diagonal
			modifierX, modifierY := 1, 1
			if x2 < x1 {
				modifierX = -1
			}
			if y2 < y1 {
				modifierY = -1
			}
			for x := x1; x != (x2 + modifierX); x += modifierX {
				if gridAll[x] == nil {
					gridAll[x] = make(map[int]int)
				}
				gridAll[x][y1]++
				y1 += modifierY
			}
		}
	}

	overlapsFlat := 0
	for _, x := range gridFlat {
		for _, point := range x {
			if point > 1 {
				overlapsFlat++
			}
		}
	}

	// Horizontal Vertical print
	fmt.Println(overlapsFlat)

	overlapsAll := 0
	for _, x := range gridAll {
		for _, point := range x {
			if point > 1 {
				overlapsAll++
			}
		}
	}
	fmt.Println(overlapsAll)
}
