package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/STollenaar/AdventOfCode/internal"
)

type State string

var (
	file       State = "FILE"
	empty      State = "EMPTY"
	fileBlocks       = make(map[int]int)
)

func main() {
	lines := internal.Reader()
	line := lines[0]

	var id int
	state := file

	var fileSystem []string
	start := time.Now()
	for _, c := range line {
		nmbr, _ := strconv.Atoi(string(c))
		switch state {
		case file:
			n := strings.TrimSpace(strings.Repeat(fmt.Sprintf("%d ", id), nmbr))
			split := strings.Split(n, " ")
			fileSystem = append(fileSystem, split...)
			state = empty
			fileBlocks[id] = nmbr
			id++
		case empty:
			fileSystem = append(fileSystem, strings.Split(strings.Repeat(".", nmbr), "")...)
			state = file
		}
	}

	part2FileSystem := make([]string, len(fileSystem))
	copy(part2FileSystem, fileSystem)

	fmt.Printf("Part 1 %d, Duration: %v\n", sortPart1(fileSystem), time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2: %d, Duration %v\v", sortPart2(part2FileSystem), time.Since(start))
}

func sortPart1(fileSystem []string) (chk int) {
	var j int
	for i := len(fileSystem) - 1; i >= 0 && i > j; i-- {
		if fileSystem[i] != "." {
			for ; j < i; j++ {
				if fileSystem[j] == "." {
					fileSystem[j] = fileSystem[i]
					fileSystem[i] = "."
					break
				}
			}
		}
	}
	for i, n := range fileSystem {
		if n == "." {
			break
		}
		nmbr, _ := strconv.Atoi(n)
		chk += i * nmbr
	}
	return
}

func sortPart2(fileSystem []string) (chk int) {
	for i := len(fileSystem) - 1; i >= 0; i-- {
		if fileSystem[i] != "." {
			id, _ := strconv.Atoi(fileSystem[i])
			block := fileBlocks[id]
			for j := 0; j < i; j++ {
				if strings.ReplaceAll(strings.Join(fileSystem[j:j+block], ""), ".", "") == "" {
					fileSystem = slices.Replace(fileSystem, j, j+block, fileSystem[i-block+1:i+1]...)
					fileSystem = slices.Replace(fileSystem, i-block+1, i+1, strings.Split(strings.Repeat(".", block), "")...)
					break
				}
			}
		}
	}
	for i, n := range fileSystem {
		nmbr, _ := strconv.Atoi(n)
		chk += i * nmbr
	}
	return
}
