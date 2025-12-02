package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()
	ids := strings.Split(lines[0], ",")

	var pt1, pt2 int
	for _, id := range ids {
		a, _ := strconv.Atoi(strings.Split(id, "-")[0])
		b, _ := strconv.Atoi(strings.Split(id, "-")[1])

		repeating, sequenced := repeatingTwice(a, b)
		pt1 += sum(repeating)
		pt2 += sum(sequenced)
	}
	fmt.Printf("Solution Part 1: %d\n", pt1)
	fmt.Printf("Solution Part 2: %d\n", pt2)
}
func repeatingTwice(start, end int) (pt1, pt2 []int) {
	for n := start; n <= end; n++ {
		if isDoubleRepeat(strconv.Itoa(n)) {
			pt1 = append(pt1, n)
		}
		if isRepeatingSequence(strconv.Itoa(n)) {
			pt2 = append(pt2, n)
		}
	}
	return
}

func isDoubleRepeat(n string) bool {
	if len(n)%2 == 0 {
		l, r := n[:len(n)/2], n[len(n)/2:]
		return l == r
	}
	return false
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func isRepeatingSequence(n string) bool {
	length := len(n)

	for i := 1; i <= length/2; i++ {
		if length%i == 0 {
			substr := n[:i]
			repeated := strings.Repeat(substr, length/i)
			if repeated == n {
				return true
			}
		}
	}
	return false
}
