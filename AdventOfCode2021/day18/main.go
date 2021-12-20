package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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

	var array string
	for scanner.Scan() {
		line := scanner.Text()

		if array == "" {
			array = line
		} else {
			array = "[" + array + "," + line + "]"
			// Do reduction
		}
		array = doReduction(array)
		if err != nil {
			log.Fatal(err)
		}
	}
	magnitude := getMagnitude(array)
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Finaly string for part 1: ", array)
	fmt.Println("Magnitude for part 1: ", magnitude)
}

func getMagnitude(snail string) (magnitude int) {
	for i := 0; i < len(snail); i++ {
		element := snail[i]
		if _, err := strconv.Atoi(string(element)); err == nil {
			firstPair, _ := strconv.Atoi(string(element))
			secondPair, _ := strconv.Atoi(string(snail[i+2]))
			magnitude += firstPair*3 + secondPair*2
			i += 2
		}
	}
	return magnitude
}

func doReduction(snail string) string {
	checkReduction := true
	for checkReduction {
		checkReduction = false

		snail, checkReduction = doExploders(snail)

		if !checkReduction {
			// Check pair split
			snail, checkReduction = doSplits(snail)
		}
		fmt.Println("Line after reduction step: ", snail)
	}
	return snail
}

func doExploders(snail string) (string, bool) {
	var depth int
	for i, element := range snail {
		if string(element) == "[" && i != 0 {
			depth++
		} else if string(element) == "]" {
			depth--
		} else if string(element) != "," {
			// Explore
			if depth == 4 {
				leftPair, _ := strconv.Atoi(string(element))
				rightPair, _ := strconv.Atoi(string(snail[i+2]))

				// Shove the left pair number to the first number to the left
				var foundLength int
				for j := i - 1; j > 0; j-- {
					if _, err := strconv.Atoi(string(snail[j])); err == nil {
						found, _ := strconv.Atoi(string(snail[j]))
						found += leftPair
						foundLength = len(strconv.Itoa(found)) - 1
						snail = snail[:j] + strconv.Itoa(found) + snail[j+1:]
						break
					}
				}

				// Shove the right pair number to the first number to the right
				for j := i + 3 + foundLength; j < len(snail); j++ {
					if _, err := strconv.Atoi(string(snail[j])); err == nil {
						found, _ := strconv.Atoi(string(snail[j]))
						found += rightPair
						snail = snail[:j] + strconv.Itoa(found) + snail[j+1:]
						break
					}
				}

				// Replacing exploding pair by a 0
				snail = snail[:i-1+foundLength] + "0" + snail[i+foundLength+4:]
				depth--
				return snail, true
			}
		}
	}
	return snail, false
}

func doSplits(snail string) (string, bool) {
	for i, element := range snail {
		// Quick and dirty way to find a number larger than 9
		if _, err := strconv.Atoi(string(element)); err == nil {
			if _, err := strconv.Atoi(string(snail[i+1])); err == nil {
				numberString := snail[i : i+2]
				number, _ := strconv.Atoi(numberString)
				leftPair := int(math.Floor(float64(number) / 2))
				rightPair := int(math.Ceil(float64(number) / 2))
				snail = snail[:i] + "[" + strconv.Itoa(leftPair) + "," + strconv.Itoa(rightPair) + "]" + snail[i+2:]
				return snail, true
			}
		}
	}
	return snail, false
}
