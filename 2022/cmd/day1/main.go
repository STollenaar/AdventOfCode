package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	internal "github.com/STollenaar/AdventOfCode/internal"
)

type Inventory struct {
	calories []int
	total    int
}

var (
	inventories []Inventory
)

func main() {

	lines := internal.Reader()
	current := Inventory{}

	for _, line := range lines {
		if line == "" {
			inventories = append(inventories, current)
			current = Inventory{}
		} else {
			amount, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			current.calories = append(current.calories, int(amount))
			current.total += int(amount)
		}
	}
	inventories = append(inventories, current)
	sort.Slice(inventories, func(i, j int) bool {
		return inventories[i].total > inventories[j].total
	})

	fmt.Printf("Maximum amount of calories carried by 1 elf (problem 1): %d\n", inventories[0].total)
	top3 := inventories[:3]

	var top3Total int
	for _, inv := range top3 {
		top3Total += inv.total
	}
	fmt.Printf("Maximum amount of calories carried by top 3 elves (problem 2): %d\n", top3Total)
}
