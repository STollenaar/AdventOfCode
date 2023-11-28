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

type Fish struct {
	timer  int
	amount int
}

func getFishAmount(fishes []*Fish) (amount int) {
	for _, fish := range fishes {
		amount += fish.amount
	}
	return amount
}

func merge(slice []*Fish) (merged []*Fish) {
	for i, fish := range slice {
		inSlice := getFish(merged, fish.timer, i)
		if inSlice == nil {
			merged = append(merged, fish)
		} else {
			inSlice.amount += fish.amount
		}
	}
	return merged
}

func getFish(fishes []*Fish, value int, ignoreIndex int) *Fish {
	for i, fish := range fishes {
		if (ignoreIndex != -1 || i != ignoreIndex) && fish.timer == value {
			return fish
		}
	}
	return nil
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var fishes []*Fish
	for scanner.Scan() {
		line := scanner.Text()

		for _, state := range strings.Split(line, ",") {
			initialState, _ := strconv.Atoi(state)

			existingFish := getFish(fishes, initialState, -1)
			if existingFish == nil {
				fishes = append(fishes, &Fish{initialState, 1})
			} else {
				existingFish.amount++
			}
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	for x := 1; x <= 80; x++ {
		newBornAmount := 0
		for _, fish := range fishes {
			if fish.timer == 0 {
				newBornAmount += fish.amount
				fish.timer = 6
			} else {
				fish.timer--
			}
		}
		if newBornAmount != 0 {
			fishes = append(fishes, &Fish{8, newBornAmount})
		}
		fishes = merge(fishes) // This is to keep the slice nice and small.
	}

	// Part 1 print
	elapsed := time.Since(start)
	lapPart1 := time.Now()
	fmt.Println(getFishAmount(fishes))
	fmt.Println("Execution time for part 1: ", elapsed)

	// Looping another 176 days after the initial 80 days for part 1
	for x := 1; x <= 176; x++ {
		newBornAmount := 0
		for _, fish := range fishes {
			if fish.timer == 0 {
				newBornAmount += fish.amount
				fish.timer = 6
			} else {
				fish.timer--
			}
		}
		if newBornAmount != 0 {
			fishes = append(fishes, &Fish{8, newBornAmount})
		}
		fishes = merge(fishes) // This is to keep the slice nice and small.
	}

	// Part 2 print
	par2Time := time.Since(lapPart1)
	fmt.Println(getFishAmount(fishes))
	fmt.Println("Execution time for part 2: ", par2Time)
}
