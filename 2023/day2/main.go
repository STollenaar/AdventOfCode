package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	cubes = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
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
		game, hands := strings.Split(line, ":")[0], strings.Split(strings.Split(line, ":")[1], ";")
		gameNmbr, _ := strconv.Atoi(strings.Split(game, " ")[1])

		br := false
		for _, hand := range hands {
			blocks := strings.Split(strings.TrimSpace(hand), ",")
			for _, block := range blocks {
				amountString, color := strings.Split(strings.TrimSpace(block), " ")[0], strings.Split(strings.TrimSpace(block), " ")[1]
				amount, _ := strconv.Atoi(amountString)
				if cubes[color] < amount {
					br = true
					break
				}
			}
			if br {
				break
			}
		}
		if !br {
			sum += gameNmbr
		}
	}
	fmt.Printf("Solution for Part1: %d\n", sum)
}

func part2(lines []string) {
	var sum int
	for _, line := range lines {
		hands := strings.Split(strings.Split(line, ":")[1], ";")
		cubes := make(map[string]int)

		for _, hand := range hands {
			blocks := strings.Split(strings.TrimSpace(hand), ",")
			for _, block := range blocks {
				amountString, color := strings.Split(strings.TrimSpace(block), " ")[0], strings.Split(strings.TrimSpace(block), " ")[1]
				amount, _ := strconv.Atoi(amountString)
				if cubes[color] < amount {
					cubes[color] = amount
				}
			}
		}
		sum += (cubes["red"] * cubes["blue"] * cubes["green"])
	}
	fmt.Printf("Solution for Part1: %d\n", sum)
}
