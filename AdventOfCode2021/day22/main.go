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

type Cube struct {
	on                                         bool
	xMin, xMax, yMin, yMax, zMin, zMax, volume int64
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func (c Cube) doIntersect(other Cube) Cube {
	var intersect Cube
	intersect.on = !other.on

	intersect.xMin = max(other.xMin, c.xMin)
	intersect.xMax = min(other.xMax, c.xMax)

	intersect.yMin = max(other.yMin, c.yMin)
	intersect.yMax = min(other.yMax, c.yMax)

	intersect.zMin = max(other.zMin, c.zMin)
	intersect.zMax = min(other.zMax, c.zMax)

	intersect.volume = (abs(intersect.xMin-intersect.xMax) + 1) * (abs(intersect.yMin-intersect.yMax) + 1) * (abs(intersect.zMin-intersect.zMax) + 1)

	return intersect
}

func (c Cube) isValidPoint() bool {
	return c.xMin <= c.xMax && c.yMin <= c.yMax && c.zMin <= c.zMax
}

func (c Cube) isInitPoint() bool {
	return c.xMin >= -50 && c.xMax <= 50 && c.yMin >= -50 && c.yMax <= 50 && c.zMin >= -50 && c.zMax <= 50
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var cubes []Cube
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		var on bool
		if line[0] == "on" {
			on = true
		} else {
			on = false
		}
		area := strings.Split(line[1], ",")

		xArea := strings.Split(strings.ReplaceAll(area[0], "x=", ""), "..")
		xMin, _ := strconv.ParseInt(xArea[0], 10, 64)
		xMax, _ := strconv.ParseInt(xArea[1], 10, 64)

		yArea := strings.Split(strings.ReplaceAll(area[1], "y=", ""), "..")
		yMin, _ := strconv.ParseInt(yArea[0], 10, 64)
		yMax, _ := strconv.ParseInt(yArea[1], 10, 64)

		zArea := strings.Split(strings.ReplaceAll(area[2], "z=", ""), "..")
		zMin, _ := strconv.ParseInt(zArea[0], 10, 64)
		zMax, _ := strconv.ParseInt(zArea[1], 10, 64)

		volume := (abs(xMin-xMax) + 1) * (abs(yMin-yMax) + 1) * (abs(zMin-zMax) + 1)
		cubes = append(cubes, Cube{on: on, xMin: xMin, xMax: xMax, yMin: yMin, yMax: yMax, zMin: zMin, zMax: zMax, volume: volume})
	}

	doPart1(cubes, start)
	start = time.Now()
	doPart2(cubes, start)
}

func doPart1(cubes []Cube, start time.Time) {

	var amount int64
	var checkedCubes []Cube

	for _, cube := range cubes {
		if cube.isInitPoint() {
			var nextCubes []Cube
			if cube.on {
				nextCubes = append(nextCubes, cube)
			}
			for _, check := range checkedCubes {
				if intersect := cube.doIntersect(check); intersect.isValidPoint() {
					nextCubes = append(nextCubes, intersect)
				}
			}
			checkedCubes = append(checkedCubes, nextCubes...)
		}
	}

	for _, checkedCube := range checkedCubes {
		if checkedCube.on {
			amount += checkedCube.volume
		} else {
			amount -= checkedCube.volume
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Amount on for part 1: ", amount)

}

func doPart2(cubes []Cube, start time.Time) {

	var amount int64
	var checkedCubes []Cube
	for _, cube := range cubes {
		var nextCubes []Cube
		if cube.on {
			nextCubes = append(nextCubes, cube)
		}
		for _, check := range checkedCubes {
			if intersect := cube.doIntersect(check); intersect.isValidPoint() {
				nextCubes = append(nextCubes, intersect)
			}
		}
		checkedCubes = append(checkedCubes, nextCubes...)
	}

	for _, checkedCube := range checkedCubes {
		if checkedCube.on {
			amount += checkedCube.volume
		} else {
			amount -= checkedCube.volume
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Amount on for part 2: ", amount)
}
