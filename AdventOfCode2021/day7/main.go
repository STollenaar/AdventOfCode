package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func getMin(slice []int) (min int) {
	min = slice[0]
	for _, val := range slice {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var crabs []int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		for _, crab := range strings.Split(line, ",") {
			crabPos, _ := strconv.Atoi(crab)
			crabs = append(crabs, crabPos)
		}

	}
	sort.Ints(crabs)
	start := time.Now()
	doPart1(crabs, start)
	start = time.Now()
	doPart2(crabs, start)
}

func doPart1(crabs []int, startTime time.Time) {
	var fuelCosts []int
	for pos := range crabs {
		var cost int
		for _, crab := range crabs {
			cost += int(math.Abs(float64(pos) - float64(crab)))
		}
		fuelCosts = append(fuelCosts, cost)
	}

	minAmount := getMin(fuelCosts)
	elapsed := time.Since(startTime)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Part 1 answer: ", minAmount)
}

func doPart2(crabs []int, startTime time.Time) {
	var fuelCosts []int
	for pos := range crabs {
		var cost int
		for _, crab := range crabs {
			diff := int(math.Abs(float64(crab) - float64(pos)))
			sum := (diff * (diff + 1)) / 2
			cost += sum
		}
		fuelCosts = append(fuelCosts, cost)
	}

	minAmount := getMin(fuelCosts)
	elapsed := time.Since(startTime)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Part 2 answer: ", minAmount)
}
