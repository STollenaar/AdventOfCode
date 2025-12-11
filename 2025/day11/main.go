package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var nodes map[string][]string

func main() {
	nodes = make(map[string][]string)
	lines := internal.Reader()

	for _, line := range lines {
		key := strings.Split(line, ":")[0]
		values := strings.Split(strings.Split(line, ": ")[1], " ")
		nodes[key] = values
	}

	paths := allPaths(nodes, "you", "out")
	fmt.Printf("Part1: %d\n", len(paths))

	amount := countPathsIncluding(nodes, "svr", "out", "dac", "fft")
	fmt.Printf("Part2: %d\n", amount)

}

func allPaths(nodes map[string][]string, start, finish string) [][]string {
	var result [][]string
	var path []string

	var dfs func(curr string, visited map[string]bool)
	dfs = func(curr string, visited map[string]bool) {
		path = append(path, curr)

		if curr == finish {
			// Copy the path so subsequent modifications don't mutate result
			cp := make([]string, len(path))
			copy(cp, path)
			result = append(result, cp)

			path = path[:len(path)-1]
			return
		}

		visited[curr] = true

		for _, next := range nodes[curr] {
			if !visited[next] {
				dfs(next, clone(visited))
			}
		}

		path = path[:len(path)-1]
	}

	dfs(start, make(map[string]bool))
	return result
}

func countPathsIncluding(nodes map[string][]string, start, finish, a, b string) int {
	memo := make(map[string]map[string]int)

	var count func(src, dst string) int
	count = func(src, dst string) int {
		if _, ok := memo[src]; !ok {
			memo[src] = make(map[string]int)
		}
		if v, ok := memo[src][dst]; ok {
			return v
		}

		if src == dst {
			return 1
		}

		total := 0
		for _, next := range nodes[src] {
			total += count(next, dst)
		}

		memo[src][dst] = total
		return total
	}

	p1 := count(start, a) * count(a, b) * count(b, finish)
	p2 := count(start, b) * count(b, a) * count(a, finish)

	return p1 + p2
}

func clone(m map[string]bool) map[string]bool {
	cp := make(map[string]bool, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func filter(in [][]string) (out [][]string) {
	for _, i := range in {
		if slices.Contains(i, "dac") && slices.Contains(i, "fft") {
			out = append(out, i)
		}
	}
	return
}
