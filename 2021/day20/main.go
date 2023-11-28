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
	imageEnhancement := true
	var imageAlgorithm string

	grid := make(map[string]string)
	var rows int
	var colums int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			imageEnhancement = false
		} else if imageEnhancement {
			imageAlgorithm += line
		} else {
			pixels := strings.Split(line, "")
			for x, pixel := range pixels {
				grid[toKey(x, rows)] = pixel
				if x > colums {
					colums = x
				}
			}
			rows++
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	part1Enhance := doEnchancements(2, rows, imageAlgorithm, grid)
	pixelsLitPart1 := getPixelsLit(part1Enhance)
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Amount of pixels lit for part 1: ", pixelsLitPart1)

	start = time.Now()

	part2Enhance := doEnchancements(50, rows, imageAlgorithm, grid)
	pixelsLitPart2 := getPixelsLit(part2Enhance)
	elapsed = time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Amount of pixels lit for part 1: ", pixelsLitPart2)
}

func doEnchancements(steps, rows int, imageAlgorithm string, grid map[string]string) map[string]string {
	for step := 0; step < steps; step++ {
		gridCopy := copyGrid(grid)
		for y := -rows - step; y <= rows+step; y++ {
			for x := -rows - step; x <= rows+step; x++ {
				binaryNumber := getBinaryNumber(x, y, step, grid)
				gridCopy[toKey(x, y)] = string(imageAlgorithm[binaryNumber])
			}
		}
		grid = gridCopy
	}
	return grid
}

func toKey(x, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func getBinaryNumber(x, y, step int, grid map[string]string) int64 {
	var binaryString string
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if grid[toKey(x+j, y+i)] == "#" {
				binaryString += "1"
			} else if grid[toKey(x+j, y+i)] == "." {
				binaryString += "0"
			} else {
				if step%2 == 1 {
					binaryString += "1"
				} else {
					binaryString += "0"
				}
			}
		}
	}

	i, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func copyGrid(grid map[string]string) map[string]string {
	duplicate := make(map[string]string)
	for k, v := range grid {
		duplicate[k] = v
	}
	return duplicate
}

func getPixelsLit(grid map[string]string) (amount int) {
	for _, pixel := range grid {
		if pixel == "#" {
			amount++
		}
	}
	return amount
}
