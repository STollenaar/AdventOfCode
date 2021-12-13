package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	grid := make(map[int]map[int]*string)

	calledDots := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if !strings.Contains(line, "fold") {
			points := strings.Split(line, ",")
			x, _ := strconv.Atoi(points[0])
			y, _ := strconv.Atoi(points[1])

			if grid[y] == nil {
				grid[y] = make(map[int]*string)
			}
			grid[y][x] = new(string)
			*grid[y][x] = "#"
		} else {
			fillMap(grid)
			cord := strings.Split(strings.ReplaceAll(line, "fold along ", ""), "=")
			position, _ := strconv.Atoi(cord[1])

			if cord[0] == "x" {
				var maxColumns int
				for _, row := range grid {
					maxRow := maxKeyRow(row)
					if maxRow > maxColumns {
						maxColumns = maxRow
					}
				}
				for y := 0; y <= maxKeyGrid(grid); y++ {
					for x := 0; x <= maxColumns; x++ {
						if x <= position {
							continue
						}
						mirrorX := position - (x - position)
						if *grid[y][x] == "#" {
							grid[y][mirrorX] = grid[y][x]
						}
						delete(grid[y], x)
					}
				}
			} else {
				var maxColumns int
				for _, row := range grid {
					maxRow := maxKeyRow(row)
					if maxRow > maxColumns {
						maxColumns = maxRow
					}
				}
				for y := 0; y <= maxKeyGrid(grid); y++ {
					if y < position {
						continue
					}
					mirrorY := position - (y - position)
					for x := 0; x <= maxColumns; x++ {
						if *grid[y][x] == "#" {
							grid[mirrorY][x] = grid[y][x]
						}
						delete(grid[y], x)
					}
					delete(grid, y)
				}
			}
			if !calledDots {
				elapsed := time.Since(start)
				fmt.Println("Execution time for part 1: ", elapsed)
				fmt.Println("Amount of dots for part 1: ", getAmountDots(grid))
				calledDots = true
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	var maxColumns int
	for _, row := range grid {
		maxRow := maxKeyRow(row)
		if maxRow > maxColumns {
			maxColumns = maxRow
		}
	}
	for y := 0; y <= maxKeyGrid(grid); y++ {
		stringBuilder := ""
		for x := 0; x <= maxColumns; x++ {
			stringBuilder += *grid[y][x]
		}
		fmt.Println(stringBuilder)
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
}

func getAmountDots(grid map[int]map[int]*string) (dots int) {
	for _, row := range grid {
		for _, column := range row {
			if column == nil {
				continue
			}
			if *column == "#" {
				dots++
			}
		}
	}
	return dots
}

func fillMap(grid map[int]map[int]*string) {
	var maxColumns int
	for _, row := range grid {
		maxRow := maxKeyRow(row)
		if maxRow > maxColumns {
			maxColumns = maxRow
		}
	}

	for y := 0; y <= maxKeyGrid(grid); y++ {
		if grid[y] == nil {
			grid[y] = make(map[int]*string)
		}
		for x := 0; x <= maxColumns; x++ {
			if grid[y][x] == nil {
				grid[y][x] = new(string)
				*grid[y][x] = "."
			}
		}
	}
}

func maxKeyGrid(grid map[int]map[int]*string) int {
	var maxNumber int
	for n := range grid {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

func maxKeyRow(row map[int]*string) int {
	var maxNumber int
	for n := range row {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}
