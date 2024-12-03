package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

const mulsR = `mul\([0-9]{1,3},[0-9]{1,3}\)`
const doMulR = `mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`

var mulReg = regexp.MustCompile(mulsR)
var doReg = regexp.MustCompile(doMulR)

func main() {
	lines := internal.Reader()
	line := strings.Join(lines, "\n")
	part1Matches := mulReg.FindAllString(line, -1)

	var totalPart1, totalPart2 int
	for _, match := range part1Matches {
		totalPart1 += doMul(match)
	}

	fmt.Printf("Part 1: %d\n", totalPart1)

	part2Matches:= doReg.FindAllString(line, -1)

	mult := true
	for _, match := range part2Matches {
		if match == "do()" {
			mult = true
		} else if match == "don't()" {
			mult = false
		} else if mult{
			totalPart2 += doMul(match)
		}
	}
	fmt.Printf("Part 2: %d\n", totalPart2)
}

func doMul(in string) int {
	t := strings.ReplaceAll(in, "mul(", "")
		t = strings.ReplaceAll(t, ",", " ")
		t = strings.ReplaceAll(t, ")", " ")

		nmbrs := strings.Split(t, " ")

		left, _ := strconv.Atoi(nmbrs[0])
		right, _ := strconv.Atoi(nmbrs[1])
		return left*right
}