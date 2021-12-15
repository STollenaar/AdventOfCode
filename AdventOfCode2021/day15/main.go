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

type Point struct {
	x        int
	y        int
	value    int
	weight   int
	distance int

	previous *Point
}

type Queue []*Point

func (q *Queue) Push(p *Point) {
	*q = append(*q, p)
}

func (q *Queue) Sort() {
	sort.Slice(*q, func(i, j int) bool {
		if (*q)[i].weight == (*q)[j].weight {
			return (*q)[i].value < (*q)[j].value
		}
		return (*q)[i].weight < (*q)[j].weight
	})
}

func (q *Queue) Delete(point *Point) {
	for i, iQ := range *q {
		if iQ == point {
			*q = append((*q)[:i], (*q)[i+1:]...)
			break
		}
	}
}

func (q *Queue) Pop() *Point {
	last := (*q)[0]
	*q = (*q)[1:]
	return last
}

func (q *Queue) CullDeadBranches() {
	closest := (*q)[0]
	for _, point := range *q {
		if point.distance > closest.distance+5 {
			q.Delete(point)
		}
	}
}

func inSlice(slice []*Point, currentPoint *Point) bool {
	for _, point := range slice {
		if point.x == currentPoint.x && currentPoint.y == point.y {
			return true
		}
	}
	return false
}

func inSliceQueue(slice []*Point, currentPoint *Point) *Point {
	for _, point := range slice {
		if point.x == currentPoint.x && currentPoint.y == point.y {
			return point
		}
	}
	return nil
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()

		risks := strings.Split(line, "")
		var row []int
		for _, c := range risks {
			risk, _ := strconv.Atoi(c)
			row = append(row, risk)
		}
		grid = append(grid, row)
		if err != nil {
			log.Fatal(err)
		}
	}
	path := search(grid)

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Total risk for part 1: ", path.weight)

	start = time.Now()
	// Now increasing the grid size
	grid = expandGrid(grid)
	elapsed = time.Since(start)
	fmt.Println("Expanding the grid for part 2 took: ", elapsed)
	// Searching again
	path = search(grid)

	elapsed = time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Total risk for part 2: ", path.weight)

}

// Classic dijkstra
// might be able to be optimized even more.
func search(grid [][]int) (endPoint *Point) {
	maxY := len(grid)
	maxX := len(grid[maxY-1])

	var queue Queue
	queue = append(queue, &Point{x: 0, y: 0, value: 0, distance: maxX + maxY, previous: nil})

	var visitedPoints []*Point
	for len(queue) > 0 {
		currentPoint := queue.Pop()
		visitedPoints = append(visitedPoints, currentPoint)
		if currentPoint.y == len(grid)-1 && currentPoint.x == len(grid[len(grid)-1])-1 {
			endPoint = currentPoint
			break
		}

		if currentPoint.x-1 >= 0 && !inSlice(visitedPoints, &Point{x: currentPoint.x - 1, y: currentPoint.y}) {
			point := &Point{x: currentPoint.x - 1, y: currentPoint.y, value: grid[currentPoint.y][currentPoint.x-1], previous: currentPoint}
			point.distance = getDistance(point, maxX, maxY)
			point.weight = currentPoint.weight + point.value

			if q := inSliceQueue(queue, point); q != nil && point.weight < q.weight {
				queue.Delete(q)
				queue.Push(point)
			} else if q == nil {
				queue.Push(point)
			}
		}
		if currentPoint.x+1 < len(grid[currentPoint.y]) && !inSlice(visitedPoints, &Point{x: currentPoint.x + 1, y: currentPoint.y}) {
			point := &Point{x: currentPoint.x + 1, y: currentPoint.y, value: grid[currentPoint.y][currentPoint.x+1], previous: currentPoint}
			point.distance = getDistance(point, maxX, maxY)
			point.weight = currentPoint.weight + point.value

			if q := inSliceQueue(queue, point); q != nil && point.weight < q.weight {
				queue.Delete(q)
				queue.Push(point)
			} else if q == nil {
				queue.Push(point)
			}
		}
		if currentPoint.y-1 >= 0 && !inSlice(visitedPoints, &Point{x: currentPoint.x, y: currentPoint.y - 1}) {
			point := &Point{x: currentPoint.x, y: currentPoint.y - 1, value: grid[currentPoint.y-1][currentPoint.x], previous: currentPoint}
			point.distance = getDistance(point, maxX, maxY)
			point.weight = currentPoint.weight + point.value

			if q := inSliceQueue(queue, point); q != nil && point.weight < q.weight {
				queue.Delete(q)
				queue.Push(point)
			} else if q == nil {
				queue.Push(point)
			}
		}
		if currentPoint.y+1 < len(grid) && !inSlice(visitedPoints, &Point{x: currentPoint.x, y: currentPoint.y + 1}) {
			point := &Point{x: currentPoint.x, y: currentPoint.y + 1, value: grid[currentPoint.y+1][currentPoint.x], previous: currentPoint}
			point.distance = getDistance(point, maxX, maxY)
			point.weight = currentPoint.weight + point.value

			if q := inSliceQueue(queue, point); q != nil && point.weight < q.weight {
				queue.Delete(q)
				queue.Push(point)
			} else if q == nil {
				queue.Push(point)
			}
		}
		queue.Sort()
		if len(queue) > 100 {
			queue.CullDeadBranches()
		}
	}
	return endPoint
}

func getDistance(point *Point, x, y int) int {
	return (x - point.x) + (y - point.y)
}

func expandGrid(grid [][]int) [][]int {
	ySize := len(grid) - 1
	xSize := len(grid[ySize]) - 1

	offSet := 0
	for len(grid)+1 <= (ySize+1)*5 {
		for yLevel := 0; yLevel <= ySize; yLevel++ {
			rowTmp := make([]int, len(grid[yLevel+offSet]))
			copy(rowTmp, grid[yLevel+offSet])

			for len(grid[yLevel+offSet])-1 < xSize*5 {
				for i := 0; i <= xSize; i++ {
					rowTmp[i]++
					if rowTmp[i] > 9 {
						rowTmp[i] = 1
					}
				}
				grid[yLevel+offSet] = append(grid[yLevel+offSet], rowTmp...)
			}
			grid = append(grid, grid[yLevel+offSet][xSize+1:(xSize+1)*2])
		}
		offSet += ySize + 1
	}
	for yLevel := 0; yLevel <= ySize; yLevel++ {
		rowTmp := make([]int, len(grid[yLevel+offSet]))
		copy(rowTmp, grid[yLevel+offSet])

		for len(grid[yLevel+offSet])-1 < xSize*5 {
			for i := 0; i <= xSize; i++ {
				rowTmp[i]++
				if rowTmp[i] > 9 {
					rowTmp[i] = 1
				}
			}
			grid[yLevel+offSet] = append(grid[yLevel+offSet], rowTmp...)
		}
	}
	return grid
}
