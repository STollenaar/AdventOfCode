package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Step struct {
	onOff                              bool
	xMin, xMax, yMin, yMax, zMin, zMax int
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var steps []Step
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		var onOff bool
		if line[0] == "on" {
			onOff = true
		} else {
			onOff = false
		}
		area := strings.Split(line[1], ",")

		xArea := strings.Split(strings.ReplaceAll(area[0], "x=", ""), "..")
		xMin, _ := strconv.Atoi(xArea[0])
		xMax, _ := strconv.Atoi(xArea[1])

		yArea := strings.Split(strings.ReplaceAll(area[1], "y=", ""), "..")
		yMin, _ := strconv.Atoi(yArea[0])
		yMax, _ := strconv.Atoi(yArea[1])

		zArea := strings.Split(strings.ReplaceAll(area[2], "z=", ""), "..")
		zMin, _ := strconv.Atoi(zArea[0])
		zMax, _ := strconv.Atoi(zArea[1])

		steps = append(steps, Step{onOff: onOff, xMin: xMin, xMax: xMax, yMin: yMin, yMax: yMax, zMin: zMin, zMax: zMax})
	}

	doPart1(steps, start)
	start = time.Now()
	doPart2(steps, start)
}

func doPart1(steps []Step, start time.Time) {
	reactor := make(map[string]bool)

	for _, step := range steps {
		for x := step.xMin; x <= step.xMax; x++ {
			if x < -50 || x > 50 {
				break
			}
			for y := step.yMin; y <= step.yMax; y++ {
				if y < -50 || y > 50 {
					break
				}

				for z := step.zMin; z <= step.zMax; z++ {
					if z < -50 || z > 50 {
						break
					}
					reactor[strconv.Itoa(x)+"_"+strconv.Itoa(y)+"_"+strconv.Itoa(z)] = step.onOff
				}
			}
		}
	}
	var amount int
	for _, value := range reactor {
		if value {
			amount++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Amount on for part 1: ", amount)

}

func doPart2(steps []Step, start time.Time) {
	reactor := make(map[string]bool)

	for _, step := range steps {
		for x := step.xMin; x <= step.xMax; x++ {
			for y := step.yMin; y <= step.yMax; y++ {
				for z := step.zMin; z <= step.zMax; z++ {
					reactor[strconv.Itoa(x)+"_"+strconv.Itoa(y)+"_"+strconv.Itoa(z)] = step.onOff
				}
			}
		}
	}
	var amount int
	for _, value := range reactor {
		if value {
			amount++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Amount on for part 2: ", amount)

}
