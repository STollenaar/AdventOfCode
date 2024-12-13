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

	var total int
	for _, machine := range machines {
		a := (machine.prize.x - (machine.prize.y * machine.buttonB.x / machine.buttonB.y)) / (machine.buttonA.x - (machine.buttonA.y * machine.buttonB.x / machine.buttonB.y))
		b := (machine.prize.y - machine.buttonA.y*a) / machine.buttonB.y
		fmt.Println(a, isInteger(a), b, isInteger(b))
		fmt.Printf("For Prize: %v, ButtonA: %f, ButtonB: %f\n", machine.prize, a, b)
		if isInteger(a) && isInteger(b) {
			fmt.Printf("For Prize: %v, ButtonA: %.0f, ButtonB: %.0f\n", machine.prize, a, b)
			total += (int(a)*3 + int(b))
		}
	}
	fmt.Printf("Total tokens: %d\n", total)
}

// Function to check if a number is close to an integer within a tolerance
func isInteger(value float64) bool {
	return math.Abs(value-math.Round(value)) < epsilon
}
