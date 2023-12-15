package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type lens struct {
	label string
	focal int
}

var (
	boxmap = make(map[int][]lens)
	op     = regexp.MustCompile("[-=]")
)

func main() {
	lines := internal.Reader()
	lines = strings.Split(lines[0], ",")

	var acc, total, focalPower int

	for _, sequence := range lines {
		acc = hash(sequence)
		total += acc
		seq := op.Split(sequence, -1)
		acc = hash(seq[0])
		if seq[1] == "" {
			// is a - op
			lenses := boxmap[acc]
			if i := contains(seq[0], lenses); i != -1 {
				lenses = append(lenses[:i], lenses[i+1:]...)
				boxmap[acc] = lenses
			}
		} else {
			// is a = op
			focal, _ := strconv.Atoi(seq[1])
			lenses := boxmap[acc]
			if i := contains(seq[0], lenses); i != -1 {
				lenses[i] = lens{label: seq[0], focal: focal}
			} else {
				lenses = append(lenses, lens{label: seq[0], focal: focal})
			}
			boxmap[acc] = lenses
		}
		acc = 0
	}
	fmt.Printf("Solution to Part1: %d\n", total)

	for box, lenses := range boxmap {
		for i, lens := range lenses {
			focalPower += (box + 1) * (i+1) * lens.focal
		}
	}
	fmt.Printf("Solution to Part2: %d\n", focalPower)
}

func hash(sequence string) (acc int) {
	for _, char := range sequence {
		acc += int(char)
		acc *= 17
		acc = acc % 256
	}
	return
}

func contains(label string, lenses []lens) int {
	for i, lens := range lenses {
		if lens.label == label {
			return i
		}
	}
	return -1
}
