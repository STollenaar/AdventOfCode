package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	time, dist := lines[0], lines[1]

	part1(time, dist)
	part2(time, dist)

}

func part1(time, dist string) {
	var times, dists []int
	for _, t := range strings.Split(time, " ") {
		if v, err := strconv.Atoi(strings.TrimSpace(t)); err == nil {
			times = append(times, v)
		}
	}
	for _, t := range strings.Split(dist, " ") {
		if v, err := strconv.Atoi(strings.TrimSpace(t)); err == nil {
			dists = append(dists, v)
		}
	}

	wins := 1
	for i, t := range times {

		var win int
		for j := 0; j <= t; j++ {
			remaining := t - j
			travelled := j * remaining
			if travelled > dists[i] {
				win++
			}
		}
		wins *= win
	}

	fmt.Printf("Solution to part 1: %d\n", wins)
}

func part2(time, dist string) {
	time = strings.ReplaceAll(time, " ", "")
	dist = strings.ReplaceAll(dist, " ", "")

	var times, dists []int
	for _, t := range strings.Split(time, ":") {
		if v, err := strconv.Atoi(strings.TrimSpace(t)); err == nil {
			times = append(times, v)
		}
	}
	for _, t := range strings.Split(dist, ":") {
		if v, err := strconv.Atoi(strings.TrimSpace(t)); err == nil {
			dists = append(dists, v)
		}
	}

	wins := 1
	for i, t := range times {

		var win int
		for j := 0; j <= t; j++ {
			remaining := t - j
			travelled := j * remaining
			if travelled > dists[i] {
				win++
			}
		}
		wins *= win
	}

	fmt.Printf("Solution to part 2: %d\n", wins)
}
