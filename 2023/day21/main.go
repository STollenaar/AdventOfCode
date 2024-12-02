package main

import (
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Grid struct {
	internal.Grid[string]
}

var (
	grid = new(Grid)
)

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		grid.AddRow(strings.Split(line, ""))
	}

	
}
