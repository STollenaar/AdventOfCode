package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	reg       = regexp.MustCompile("[a-z]")
	numberMap = map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "4",
		"five":  "5e",
		"six":   "6",
		"seven": "7n",
		"eight": "e8t",
		"nine":  "9e",
	}
)

func main() {
	lines := internal.Reader()

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var sum int
	for _, line := range lines {
		numbers := reg.ReplaceAllString(line, "")

		number := string(numbers[0]) + string(numbers[len(numbers)-1])
		n, _ := strconv.Atoi(number)
		sum += n
	}
	fmt.Printf("Solution to part1: %d\n", sum)
}

func part2(lines []string) {
	var sum int
	for _, line := range lines {
		for k, v := range numberMap {
			line = strings.Replace(line, k, v, -1)
		}
		numbers := reg.ReplaceAllString(line, "")

		number := string(numbers[0]) + string(numbers[len(numbers)-1])
		fmt.Println(number)
		n, _ := strconv.Atoi(number)
		sum += n
	}
	fmt.Printf("Solution to part2: %d\n", sum)
}
