package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Report struct {
	levels     []int
	unsafes    int
	decreasing bool
}

type Pair struct {
	a, b, index, diff int
}

func main() {
	lines := internal.Reader()

	var safe, unsafe []Report
	for _, line := range lines {
		l := strings.Split(line, " ")
		var levels []int
		for _, i := range l {
			r, _ := strconv.Atoi(i)
			levels = append(levels, r)
		}
		report := Report{
			levels:     levels,
			decreasing: levels[0]-levels[1] > 0,
		}

		for i := 0; i < len(levels)-1; i++ {
			switch report.decreasing {
			case true:
				if !(levels[i]-levels[i+1] >= 1 && levels[i]-levels[i+1] <= 3) {
					report.unsafes++
				}
			case false:
				if !(levels[i]-levels[i+1] <= -1 && levels[i]-levels[i+1] >= -3) {
					report.unsafes++
				}
			}
		}
		if report.unsafes == 0 {
			safe = append(safe, report)
		} else {
			unsafe = append(unsafe, report)
		}
	}
	fmt.Printf("Part 1: %d\n", len(safe))

	for _, report := range unsafe {
		report.unsafes = 0
		pairs := genPairs(report.levels)
		for i, pair := range pairs {
            rmA, rmB := make([]int, len(report.levels)),make([]int, len(report.levels))
            copy(rmA, report.levels)
            copy(rmB, report.levels)
			if i != 0 && (pair.diff > 0) != (pairs[i-1].diff > 0) {
				rmA = append(rmA[:pair.index], rmA[pair.index+1:]...)
				rmAPairs := genPairs(rmA)
				if !hasUnsafe(rmAPairs) {
					safe = append(safe, report)
					break
				}
				rmB = append(rmB[:pair.index+1], rmB[pair.index+2:]...)
				rmBPairs := genPairs(rmB)
				if !hasUnsafe(rmBPairs) {
					safe = append(safe, report)
					break
				}
			} else if i+1 < len(pairs) && (pair.diff > 0) != (pairs[i+1].diff > 0) {
				rmA = append(rmA[:pair.index], rmA[pair.index+1:]...)
				rmAPairs := genPairs(rmA)
				if !hasUnsafe(rmAPairs) {
					safe = append(safe, report)
					break
				}
				rmB = append(rmB[:pair.index+1], rmB[pair.index+2:]...)
				rmBPairs := genPairs(rmB)
				if !hasUnsafe(rmBPairs) {
					safe = append(safe, report)
					break
				}
			} else if pair.diff == 0 || pair.diff > 3 || pair.diff < -3 {
				rmA = append(rmA[:pair.index], rmA[pair.index+1:]...)
				rmAPairs := genPairs(rmA)
				if !hasUnsafe(rmAPairs) {
					safe = append(safe, report)
					break
				}
				rmB = append(rmB[:pair.index+1], rmB[pair.index+2:]...)
				rmBPairs := genPairs(rmB)
				if !hasUnsafe(rmBPairs) {
					safe = append(safe, report)
					break
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", len(safe))
}

func genPairs(in []int) (out []Pair) {
	for i := 0; i < len(in)-1; i++ {
		out = append(out, Pair{a: in[i], b: in[i+1], diff: in[i] - in[i+1], index: i})
	}
	return
}

func hasUnsafe(pairs []Pair) bool {
	for i, pair := range pairs {
		if (i != 0 && (pair.diff > 0) != (pairs[i-1].diff > 0)) || (i+1 < len(pairs) && (pair.diff > 0) != (pairs[i+1].diff > 0)) || pair.diff > 3 || pair.diff == 0 || pair.diff < -3 {
			return true
		}
	}
	return false
}
