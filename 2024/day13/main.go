package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Machine struct {
	buttonA, buttonB, prize internal.Point[float64]
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
			buttonA: internal.Point[float64]{
				X: aX,
				Y: aY,
			},
			buttonB: internal.Point[float64]{
				X: bX,
				Y: bY,
			},
			prize: internal.Point[float64]{
				X: prizeX,
				Y: prizeY,
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
		machine.prize.X += 10000000000000
		machine.prize.Y += 10000000000000
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
	a = ((machine.prize.X * machine.buttonB.Y) - (machine.prize.Y * machine.buttonB.X)) / ((machine.buttonA.X * machine.buttonB.Y) - (machine.buttonA.Y * machine.buttonB.X))
	b = ((machine.buttonA.X * machine.prize.Y) - (machine.buttonA.Y * machine.prize.X)) / ((machine.buttonA.X * machine.buttonB.Y) - (machine.buttonA.Y * machine.buttonB.X))
	return
}
