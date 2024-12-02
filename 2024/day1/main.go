package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	leftList, rightList := []int{}, []int{}
	lists := [][]int{leftList, rightList}

	for _, line := range lines {
		l := strings.Split(line, " ")
		left, _ := strconv.Atoi(l[0])
		right, _ := strconv.Atoi(l[len(l)-1])
		lists[0] = append(lists[0], left)
		lists[1] = append(lists[1], right)
	}
	sort.Ints(lists[0])
	sort.Ints(lists[1])

	var total, sim int
	for index := range lists[0] {
		total += int(math.Abs(float64(lists[0][index]) - float64(lists[1][index])))
	}

	fmt.Printf("Part 1: %d\n", total)

	for _, search := range lists[0] {
		occ := findOcc(search, lists[1])
		sim += (search * occ)

	}
	fmt.Printf("Part 2: %d\n", sim)

}

func findOcc(search int, list []int) (total int) {
	for _, v := range list {
		if v == search {
			total++
		}
	}

	return
}
