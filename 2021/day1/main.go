package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getSliceSize(slice []int) int {
	amount := 0
	for _, element := range slice {
		amount += element
	}
	return amount
}

func subSliceLarger(subA []int, subB []int) bool {
	return getSliceSize(subA) > getSliceSize(subB)
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var measurements []int

	depthIncreasesPart1, depthIncreasesPart2 := 0, 0
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if len(measurements) != 0 && measurement > measurements[len(measurements)-1] {
			depthIncreasesPart1++
		}
		measurements = append(measurements, measurement)
		if len(measurements) > 3 && subSliceLarger(measurements[len(measurements)-3:], measurements[len(measurements)-4:len(measurements)-1]) {
			depthIncreasesPart2++
		}
	}
	fmt.Println(depthIncreasesPart1)
	fmt.Println(depthIncreasesPart2)
}
