package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type LowPoint struct {
	x int
	y int
}

func inSlice(slice *[]LowPoint, lowPoint LowPoint) bool {
	for _, point := range *slice {
		if point.x == lowPoint.x && lowPoint.y == point.y {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	var heightMap [][]*int

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var row []*int
		for _, char := range strings.Split(line, "") {
			height, _ := strconv.Atoi(char)
			row = append(row, &height)
		}
		heightMap = append(heightMap, row)
		if err != nil {
			log.Fatal(err)
		}
	}

	doPart2(heightMap, start)
}

func doPart1(heightMap [][]*int, start time.Time) (lowPoints []LowPoint) {
	totalRisk := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			current := *heightMap[i][j]

			// Up Position
			if i > 0 && current >= *heightMap[i-1][j] {
				continue
			}
			// Down Position
			if i < len(heightMap)-1 && current >= *heightMap[i+1][j] {
				continue
			}
			// Left Position
			if j > 0 && current >= *heightMap[i][j-1] {
				continue
			}
			// Right Position
			if j < len(heightMap[i])-1 && current >= *heightMap[i][j+1] {
				continue
			}
			totalRisk += current + 1
			lowPoints = append(lowPoints, LowPoint{x: j, y: i})
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Total risk for part 1: ", totalRisk)
	return lowPoints
}

func doPart2(heightMap [][]*int, start time.Time) {
	lowPoints := doPart1(heightMap, start)
	start = time.Now()
	var basins []int

	for _, lowPoint := range lowPoints {
		var visitedPoints *[]LowPoint
		visitedPoints = new([]LowPoint)
		basins = append(basins, search(lowPoint, heightMap, visitedPoints, 0))
	}
	sort.Ints(basins)

	largest3 := basins[len(basins)-3:]

	totalSize := largest3[0] * largest3[1] * largest3[2]

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Total risk for part 2: ", totalSize)
}

// Classic breathfirst search
func search(lowPoint LowPoint, heightMap [][]*int, visitedPoints *[]LowPoint, basinSize int) int {
	basinSize++
	*visitedPoints = append(*visitedPoints, lowPoint)
	if lowPoint.x-1 >= 0 && *heightMap[lowPoint.y][lowPoint.x] < *heightMap[lowPoint.y][lowPoint.x-1] && *heightMap[lowPoint.y][lowPoint.x-1] < 9 && !inSlice(visitedPoints, LowPoint{x: lowPoint.x - 1, y: lowPoint.y}) {
		basinSize += search(LowPoint{x: lowPoint.x - 1, y: lowPoint.y}, heightMap, visitedPoints, 0)
	}
	if lowPoint.x+1 < len(heightMap[lowPoint.y]) && *heightMap[lowPoint.y][lowPoint.x] < *heightMap[lowPoint.y][lowPoint.x+1] && *heightMap[lowPoint.y][lowPoint.x+1] < 9 && !inSlice(visitedPoints, LowPoint{x: lowPoint.x + 1, y: lowPoint.y}) {
		basinSize += search(LowPoint{x: lowPoint.x + 1, y: lowPoint.y}, heightMap, visitedPoints, 0)
	}
	if lowPoint.y-1 >= 0 && *heightMap[lowPoint.y][lowPoint.x] < *heightMap[lowPoint.y-1][lowPoint.x] && *heightMap[lowPoint.y-1][lowPoint.x] < 9 && !inSlice(visitedPoints, LowPoint{x: lowPoint.x, y: lowPoint.y - 1}) {
		basinSize += search(LowPoint{x: lowPoint.x, y: lowPoint.y - 1}, heightMap, visitedPoints, 0)
	}
	if lowPoint.y+1 < len(heightMap) && *heightMap[lowPoint.y][lowPoint.x] < *heightMap[lowPoint.y+1][lowPoint.x] && *heightMap[lowPoint.y+1][lowPoint.x] < 9 && !inSlice(visitedPoints, LowPoint{x: lowPoint.x, y: lowPoint.y + 1}) {
		basinSize += search(LowPoint{x: lowPoint.x, y: lowPoint.y + 1}, heightMap, visitedPoints, 0)
	}

	return basinSize
}
