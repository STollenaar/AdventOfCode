package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	var total, totalCards int
	cards := make(map[int]int)
	for i := range lines {
		cards[i] = 1
	}

	for i, line := range lines {
		var matching, totalScore int
		totalCards += cards[i]
		numbers := strings.Split(line, ":")[1]
		numbers = strings.TrimSpace(numbers)

		winning, actual := strings.Split(numbers, "|")[0], strings.Split(numbers, "|")[1]
		winning = strings.TrimSpace(winning)
		actual = strings.TrimSpace(actual)

		winningSl, actualSl := strings.Split(winning, " "), strings.Split(actual, " ")
		actualSl = removeEmpty(actualSl)
		actualSl = dedub(actualSl)
		for _, a := range actualSl {
			if inSlice(a, winningSl) {
				matching++
			}
		}
		fmt.Printf("Total winning numbers: %d, ", matching)
		if matching > 0 {
			totalScore = int(math.Pow(2, float64(matching)-1))
			for x := 1; x < matching+1; x++ {
				cards[x+i] += cards[i]
			}
		}
		fmt.Printf("winning value: %d\n", totalScore)
		total += matching
	}
	fmt.Printf("Solution part1: %d\n", total)
	fmt.Printf("Solution part2: %d\n", totalCards)
}

func removeEmpty(slice []string) (result []string) {
	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}
	return
}

func inSlice(search string, slice []string) bool {
	for _, s := range slice {
		if s == search {
			return true
		}
	}
	return false
}

func dedub(slice []string) (result []string) {
	for i, s := range slice {
		if i == slices.Index[[]string](slice, s) {
			result = append(result, s)
		}
	}
	return
}
