package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Point struct {
	x, y float64
}

type Machine struct {
	buttonA, buttonB, prize Point
}

const (
	epsilon = 1e-9
)

var (
	machines []Machine
)

func main() {
	lines := internal.Reader()
	for i := 0; i < len(lines); i += 4 {
		a := strings.Split(lines[i], " ")
		b := strings.Split(lines[i+1], " ")
		prize := strings.Split(lines[i+2], " ")
		aX, _ := strconv.ParseFloat(strings.Split(strings.Split(a[2], "+")[1], ",")[0], 64)
		aY, _ := strconv.ParseFloat(strings.Split(strings.Split(a[3], "+")[1], ",")[0], 64)
		bX, _ := strconv.ParseFloat(strings.Split(strings.Split(b[2], "+")[1], ",")[0], 64)
		bY, _ := strconv.ParseFloat(strings.Split(strings.Split(b[3], "+")[1], ",")[0], 64)
		prizeX, _ := strconv.ParseFloat(strings.Split(strings.Split(prize[1], "=")[1], ",")[0], 64)
		prizeY, _ := strconv.ParseFloat(strings.Split(strings.Split(prize[2], "=")[1], ",")[0], 64)

		machines = append(machines, Machine{
			buttonA: Point{
				x: aX,
				y: aY,
			},
			buttonB: Point{
				x: bX,
				y: bY,
			},
			prize: Point{
				x: prizeX,
				y: prizeY,
			},
		})
	}

	var totalPart1, totalPart2 int
	for _, machine := range machines {
		a, b := solve(machine)
		if isInteger(a) && isInteger(b) {
			totalPart1 += (int(a)*3 + int(b))
		}
	}
	fmt.Printf("Part 1: %d\n", totalPart1)

	for _, machine := range machines {
		machine.prize.x += 10000000000000
		machine.prize.y += 10000000000000
		a, b := solve(machine)
		if isInteger(a) && isInteger(b) {
			totalPart2 += (int(a)*3 + int(b))
		}
	}
	fmt.Printf("Part 2: %d\n", totalPart2)
}

// Function to deal with stupid float precious
func isInteger(value float64) bool {
	return math.Abs(value-math.Round(value)) < epsilon
}

// I am bad at Math. Thanks internet strangers for equation, failed my linear algebra courses a couple of times
func solve(machine Machine) (a, b float64) {
	a = ((machine.prize.x * machine.buttonB.y) - (machine.prize.y * machine.buttonB.x)) / ((machine.buttonA.x * machine.buttonB.y) - (machine.buttonA.y * machine.buttonB.x))
	b = ((machine.buttonA.x * machine.prize.y) - (machine.buttonA.y * machine.prize.x)) / ((machine.buttonA.x * machine.buttonB.y) - (machine.buttonA.y * machine.buttonB.x))
	return
}
