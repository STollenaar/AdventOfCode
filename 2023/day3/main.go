package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/STollenaar/AdventOfCode/internal"
	"golang.org/x/exp/slices"
)

type Grid struct {
	internal.Grid[string]
}

var (
	schematic Grid
	reg       = regexp.MustCompile(`[@\/#$%&*\-+=]`)
)

func main() {
	lines := internal.Reader()

	for i, line := range lines {
		for _, c := range line {
			schematic.AddSafeToColumn(string(c), i)
		}
	}
	part1()
	part2()
}

func part1() {
	var sum int
	for i, row := range schematic.Rows {
		var nmbr string
		var isPrtNmbr bool
		for j, column := range row {
			if _, err := strconv.Atoi(column); err == nil {
				nmbr += column
				if !isPrtNmbr {
					isPrtNmbr = checkAdjecency(i, j)
				}
			} else {
				if isPrtNmbr {
					n, _ := strconv.Atoi(nmbr)
					sum += n
					isPrtNmbr = false
				}
				nmbr = ""
			}
		}
		if isPrtNmbr {
			n, _ := strconv.Atoi(nmbr)
			sum += n
		}
	}
	fmt.Printf("Solution for Part1: %d\n", sum)
}

func part2() {
	var sum int
	for i, row := range schematic.Rows {
		for j, column := range row {
			if column == "*" {
				var gears []string
				if i-1 >= 0 {
					gears = append(gears, creepSlice(i-1, j)...)
				}
				if i+1 < len(schematic.Rows) {
					gears = append(gears, creepSlice(i+1, j)...)
				}
				gears = append(gears, creepSlice(i, j)...)
				if len(gears) == 2 {
					gear1, _ := strconv.Atoi(gears[0])
					gear2, _ := strconv.Atoi(gears[1])
					sum += gear1 * gear2
				}
			}

		}
	}
	fmt.Printf("Solution for Part2: %d\n", sum)
}

func checkAdjecency(i, j int) bool {
	if i-1 >= 0 {
		row := schematic.GetUnsafeRow(i - 1)
		if checkNeighbours(j, row, true) {
			return true
		}
	}
	if i+1 < len(schematic.Rows) {
		row := schematic.GetUnsafeRow(i + 1)
		if checkNeighbours(j, row, true) {
			return true
		}
	}
	row := schematic.GetUnsafeRow(i)
	return checkNeighbours(j, row, false)
}

func checkNeighbours(j int, row []string, self bool) bool {
	if j-1 >= 0 && reg.MatchString(row[j-1]) {
		return true
	}
	if self && reg.MatchString(row[j]) {
		return true
	}
	if j+1 < len(row) && reg.MatchString(row[j+1]) {
		return true
	}
	return false
}

func creepSlice(i, j int) (result []string) {
	nmbr := make(map[int]string)
	row := schematic.GetUnsafeRow(i)
	for t := j; t >= 0; t-- {
		if _, err := strconv.Atoi(row[t]); err == nil {
			nmbr[t] = row[t]
		} else {
			break
		}
	}
	for t := j - 1; t >= 0; t-- {
		if _, err := strconv.Atoi(row[t]); err == nil {
			nmbr[t] = row[t]
		} else {
			break
		}
	}
	for t := j; t < len(row); t++ {
		if _, err := strconv.Atoi(row[t]); err == nil {
			nmbr[t] = row[t]
		} else {
			break
		}
	}
	for t := j + 1; t < len(row); t++ {
		if _, err := strconv.Atoi(row[t]); err == nil {
			nmbr[t] = row[t]
		} else {
			break
		}
	}

	var r []int
	for k := range nmbr {
		r = append(r, k)
	}
	slices.Sort[int](r)

	var buffer string
	for i, n := range r {
		if i != 0 {
			if r[i-1] != n-1 {
				result = append(result, buffer)
				buffer = ""
			}
		}
		buffer += nmbr[n]
	}
	if buffer != "" {
		result = append(result, buffer)
	}
	return result
}
