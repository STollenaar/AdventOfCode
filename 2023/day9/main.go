package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	var part1Total, part2Total int
	for _, line := range lines {
		nmbrs := getNmbrs(strings.Split(line, " "))

		pyrmid := [][]int{nmbrs}

		for j := 0; j < len(pyrmid); j++ {
			slice := pyrmid[j]
			if isAll0(slice) {
				break
			}
			var l []int
			for i := 1; i < len(slice); i++ {
				l = append(l, slice[i]-slice[i-1])
			}
			pyrmid = append(pyrmid, l)
		}
		var prevTotalPart1, prevTotalPart2 int
		for i := len(pyrmid) - 1; i >= 0; i-- {
			prevTotalPart1 += pyrmid[i][len(pyrmid[i])-1]
			prevTotalPart2 = pyrmid[i][0] - prevTotalPart2
		}
		part1Total += prevTotalPart1
		part2Total += prevTotalPart2
	}
	fmt.Printf("Solution for part1: %d\n", part1Total)
	fmt.Printf("Solution for part2: %d\n", part2Total)
}

func isAll0(sl []int) bool {
	for _, l := range sl {
		if l != 0 {
			return false
		}
	}
	return true
}

func getNmbrs(l []string) (result []int) {
	for _, i := range l {
		n, _ := strconv.Atoi(i)
		result = append(result, n)
	}
	return
}
