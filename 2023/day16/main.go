package main

import (
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Node struct {
	x, y      int
	direction []int
}

var (
	grid                   [][]string
	queue                  chan int
	totals                 []int
	totalGroups, totalDone int
)

func init() {
	queue = make(chan int)
}

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	// total := loopBeams([]*Node{{x: 0, y: 0, direction: []int{0, 1}}})
	// fmt.Printf("Solution for part 1: %d\n", total)
	go queueHandler()
	var waitGroup sync.WaitGroup
	var mu sync.Mutex
	for x := 0; x < len(grid[0]); x++ {
		waitGroup.Add(8)
		totalGroups += 8
		go func(x int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: x, y: 0, direction: []int{0, 1}}})
			waitGroup.Done()
		}(x, &waitGroup, &mu)
		go func(x int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: x, y: 0, direction: []int{0, -1}}})

			waitGroup.Done()
		}(x, &waitGroup, &mu)
		go func(x int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: x, y: 0, direction: []int{1, 0}}})

			waitGroup.Done()
		}(x, &waitGroup, &mu)
		go func(x int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: x, y: 0, direction: []int{-1, 0}}})

			waitGroup.Done()
		}(x, &waitGroup, &mu)
		go func(x int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: x, y: len(grid) - 1, direction: []int{0, 1}}})

			waitGroup.Done()
		}(x, &waitGroup, &mu)
		go func(x int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: x, y: len(grid) - 1, direction: []int{0, -1}}})

			waitGroup.Done()
		}(x, &waitGroup, &mu)
		go func(x int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: x, y: len(grid) - 1, direction: []int{1, 0}}})

			waitGroup.Done()
		}(x, &waitGroup, &mu)
		go func(x int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: x, y: len(grid) - 1, direction: []int{-1, 0}}})

			waitGroup.Done()
		}(x, &waitGroup, &mu)
		// waitGroup.Wait()
	}
	for y := 0; y < len(grid); y++ {
		waitGroup.Add(8)
		totalGroups += 8
		go func(y int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: 0, y: y, direction: []int{0, 1}}})

			waitGroup.Done()
		}(y, &waitGroup, &mu)
		go func(y int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: 0, y: y, direction: []int{0, -1}}})

			waitGroup.Done()
		}(y, &waitGroup, &mu)
		go func(y int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: 0, y: y, direction: []int{1, 0}}})

			waitGroup.Done()
		}(y, &waitGroup, &mu)
		go func(y int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: 0, y: y, direction: []int{-1, 0}}})

			waitGroup.Done()
		}(y, &waitGroup, &mu)
		go func(y int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: len(grid[0]) - 1, y: y, direction: []int{0, 1}}})

			waitGroup.Done()
		}(y, &waitGroup, &mu)
		go func(y int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: len(grid[0]) - 1, y: y, direction: []int{0, -1}}})

			waitGroup.Done()
		}(y, &waitGroup, &mu)
		go func(y int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: len(grid[0]) - 1, y: y, direction: []int{1, 0}}})

			waitGroup.Done()
		}(y, &waitGroup, &mu)
		go func(y int, wg *sync.WaitGroup, mu *sync.Mutex) {
			queue <- loopBeams([]*Node{{x: len(grid[0]) - 1, y: y, direction: []int{-1, 0}}})

			waitGroup.Done()
		}(y, &waitGroup, &mu)
		// waitGroup.Wait()
	}
	waitGroup.Wait()
	fmt.Printf("Solution for part 2: %d\n", slices.Max(totals))
}

func queueHandler() {
	for {
		i := <-queue
		totalDone++
		fmt.Printf("Remaining groups %d/%d\n",totalDone, totalGroups)
		totals = append(totals, i)
	}
}

func loopBeams(beams []*Node) (total int) {
	var energy [][]int
	for i := range grid {
		energy = append(energy, make([]int, len(grid[i])))
	}

	for len(beams) > 0 {
		var infLoop bool
		for i := 0; i < len(beams); i++ {
			beam := beams[i]
			energy[beam.y][beam.x]++
			beams = travel(beams, i)

			if energy[beam.y][beam.x] == len(grid)*len(grid[0]) {
				infLoop = true
			}
		}
		if infLoop {
			break
		}
	}
	for _, row := range energy {
		for _, c := range row {
			if c != 0 {
				total++
			}
		}
	}
	return
}

func travel(beams []*Node, index int) []*Node {
	beam := beams[index]
	switch grid[beam.y][beam.x] {
	case "/":
		if beam.direction[1] == 1 {
			beam.direction[1] = 0
			beam.direction[0] = -1
		} else if beam.direction[1] == -1 {
			beam.direction[1] = 0
			beam.direction[0] = 1
		} else if beam.direction[0] == 1 {
			beam.direction[0] = 0
			beam.direction[1] = -1
		} else {
			beam.direction[0] = 0
			beam.direction[1] = 1
		}
	case "\\":
		if beam.direction[1] == 1 {
			beam.direction[1] = 0
			beam.direction[0] = 1
		} else if beam.direction[1] == -1 {
			beam.direction[1] = 0
			beam.direction[0] = -1
		} else if beam.direction[0] == 1 {
			beam.direction[0] = 0
			beam.direction[1] = 1
		} else {
			beam.direction[0] = 0
			beam.direction[1] = -1
		}
	case "-":
		if beam.direction[0] != 0 {
			beam.direction[0] = 0
			beam.direction[1] = -1
			beamCopy := &Node{x: beam.x, y: beam.y, direction: slices.Clone(beam.direction)}
			beamCopy.direction[1] = 1
			beams = append(beams, beamCopy)
		}
	case "|":
		if beam.direction[1] != 0 {
			beam.direction[1] = 0
			beam.direction[0] = -1
			beamCopy := &Node{x: beam.x, y: beam.y, direction: slices.Clone(beam.direction)}
			beamCopy.direction[0] = 1
			beams = append(beams, beamCopy)
		}
	}

	if beam.x+beam.direction[1] >= 0 && beam.x+beam.direction[1] < len(grid[0]) {
		beam.x += beam.direction[1]
	} else {
		if index+1 < len(beams) {
			beams = append(beams[:index], beams[index+1:]...)
		} else {
			beams = beams[:index]
		}
	}
	if beam.y+beam.direction[0] >= 0 && beam.y+beam.direction[0] < len(grid) {
		beam.y += beam.direction[0]
	} else {
		if index+1 < len(beams) {
			beams = append(beams[:index], beams[index+1:]...)
		} else {
			beams = beams[:index]
		}
	}
	return beams
}
