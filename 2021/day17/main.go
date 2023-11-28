package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	var targetArea []string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "target area: ", "")

		targetArea = strings.Split(line, ", ")
	}

	xArea := strings.Split(targetArea[0], "..")
	xMin, _ := strconv.Atoi(strings.ReplaceAll(xArea[0], "x=", ""))
	xMax, _ := strconv.Atoi(xArea[1])

	yArea := strings.Split(targetArea[1], "..")
	yMin, _ := strconv.Atoi(strings.ReplaceAll(yArea[0], "y=", ""))
	yMax, _ := strconv.Atoi(yArea[1])

	doPart1(xMin, xMax, yMin, yMax, start)
	start = time.Now()
	doPart2(xMin, xMax, yMin, yMax, start)

}

func doPart1(xMin, xMax, yMin, yMax int, start time.Time) {
	var maxY int
	for attemptY := 0; attemptY < 1000; attemptY++ {
		attemptMax := calcTrickShot(attemptY, 0, 0, yMin, yMax, true)
		if attemptMax > maxY {
			maxY = attemptMax
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Highest Y jump for part 1: ", maxY)
}

func doPart2(xMin, xMax, yMin, yMax int, start time.Time) {
	var validAttempts int

	for attemptX := 0; attemptX <= xMax; attemptX++ {
		for attemptY := yMin; attemptY < 1000; attemptY++ {
			if calcVelocities(attemptX, attemptY, 0, 0, xMin, xMax, yMin, yMax) {
				validAttempts++
			}
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Amount of valid starting positions for part 2: ", validAttempts)
}

func calcTrickShot(vel int, currentPos, max, minPos, maxPos int, canGoNegative bool) int {
	currentPos += vel
	vel--

	if minPos <= currentPos && currentPos <= maxPos {
		return max
	} else if currentPos < minPos {
		return -math.MaxInt32
	}

	if currentPos > max {
		max = currentPos
	}
	if vel > 0 || (vel <= 0 && canGoNegative) {
		return calcTrickShot(vel, currentPos, max, minPos, maxPos, canGoNegative)
	} else {
		return max
	}
}

// Gotta make sure we travel correctly, hence more parameters to track
func calcVelocities(velX, velY, currentPosX, currentPosY, minPosX, maxPosX, minPosY, maxPosY int) bool {
	currentPosY += velY
	velY--

	if velX > 0 {
		currentPosX += velX
		velX--
	}

	// We reached the target area
	if minPosX <= currentPosX && currentPosX <= maxPosX && minPosY <= currentPosY && currentPosY <= maxPosY {
		return true
	} else if currentPosX > maxPosX || currentPosY < minPosY { // Can't return to these overshot positions
		return false
	}

	return calcVelocities(velX, velY, currentPosX, currentPosY, minPosX, maxPosX, minPosY, maxPosY)
}
