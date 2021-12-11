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
	var grid [][]*int
	for scanner.Scan() {
		line := scanner.Text()

		var row []*int
		for _, number := range strings.Split(line, "") {
			energy, _ := strconv.Atoi(number)
			row = append(row, &energy)
		}
		grid = append(grid, row)

		if err != nil {
			log.Fatal(err)
		}
	}

	totalFlashes := new(int)

	for day := 0; ; day++ {
		var flashed [10][10]*bool
		currentAmount := *totalFlashes

		for i, row := range flashed {
			for j := range row {
				flashed[i][j] = new(bool)
			}
		}

		for y := 0; y < len(grid); y++ {
			row := grid[y]
			for x := 0; x < len(row); x++ {
				if !*flashed[y][x] {
					*grid[y][x]++
					if *grid[y][x] > 9 {
						*grid[y][x] = 0
						flashed[y][x] = new(bool)
						*flashed[y][x] = true

						*totalFlashes = *totalFlashes + 1
						doNeightbours(x, y, grid, flashed, totalFlashes)
					}
				}
			}
		}
		if day == 99 {
			elapsed := time.Since(start)
			fmt.Println("Execution time for part 1: ", elapsed)
			fmt.Println("Total amount of flashes part 1: ", *totalFlashes)
		}
		if *totalFlashes-currentAmount == 100 {
			elapsed := time.Since(start)
			fmt.Println("Execution time for part 2: ", elapsed)
			fmt.Println("Total amount of flashes part 2: ", day+1)
			break
		}
	}
}

func doNeightbours(x, y int, grid [][]*int, flashed [10][10]*bool, totalFlashes *int) {
	if x > 0 {
		if !*flashed[y][x-1] {
			*grid[y][x-1]++
			if *grid[y][x-1] > 9 {

				*grid[y][x-1] = 0
				*flashed[y][x-1] = true
				*totalFlashes++
				doNeightbours(x-1, y, grid, flashed, totalFlashes)
			}
		}
	}
	if x+1 < len(grid[y]) {
		if !*flashed[y][x+1] {
			*grid[y][x+1]++
			if *grid[y][x+1] > 9 {

				*grid[y][x+1] = 0
				*flashed[y][x+1] = true
				*totalFlashes++
				doNeightbours(x+1, y, grid, flashed, totalFlashes)
			}
		}
	}

	if y > 0 {
		if !*flashed[y-1][x] {
			*grid[y-1][x]++
			if *grid[y-1][x] > 9 {

				*grid[y-1][x] = 0
				*flashed[y-1][x] = true
				*totalFlashes++
				doNeightbours(x, y-1, grid, flashed, totalFlashes)
			}
		}

		if x > 0 {
			if !*flashed[y-1][x-1] {
				*grid[y-1][x-1]++
				if *grid[y-1][x-1] > 9 {

					*grid[y-1][x-1] = 0
					*flashed[y-1][x-1] = true
					*totalFlashes++
					doNeightbours(x-1, y-1, grid, flashed, totalFlashes)
				}
			}
		}
		if x+1 < len(grid[y-1]) {
			if !*flashed[y-1][x+1] {
				*grid[y-1][x+1]++
				if *grid[y-1][x+1] > 9 {

					*grid[y-1][x+1] = 0
					*flashed[y-1][x+1] = true
					*totalFlashes++
					doNeightbours(x+1, y-1, grid, flashed, totalFlashes)
				}
			}
		}
	}

	if y+1 < len(grid) {
		if !*flashed[y+1][x] {
			*grid[y+1][x]++
			if *grid[y+1][x] > 9 {

				*grid[y+1][x] = 0
				*flashed[y+1][x] = true
				*totalFlashes++
				doNeightbours(x, y+1, grid, flashed, totalFlashes)
			}
		}

		if x > 0 {
			if !*flashed[y+1][x-1] {
				*grid[y+1][x-1]++
				if *grid[y+1][x-1] > 9 {

					*grid[y+1][x-1] = 0
					*flashed[y+1][x-1] = true
					*totalFlashes++
					doNeightbours(x-1, y+1, grid, flashed, totalFlashes)
				}
			}
		}
		if x+1 < len(grid[y+1]) {
			if !*flashed[y+1][x+1] {
				*grid[y+1][x+1]++
				if *grid[y+1][x+1] > 9 {

					*grid[y+1][x+1] = 0
					*flashed[y+1][x+1] = true
					*totalFlashes++
					doNeightbours(x+1, y+1, grid, flashed, totalFlashes)
				}
			}
		}
	}
}
