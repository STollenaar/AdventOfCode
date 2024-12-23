package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	connections := make(map[string][]string)

	for _, line := range lines {
		left, right := strings.Split(line, "-")[0], strings.Split(line, "-")[1]
		connections[left] = append(connections[left], right)
		slices.Sort(connections[left])
		connections[right] = append(connections[right], left)
		slices.Sort(connections[right])
	}

	var part1Total int
	for _, connection := range findThreeway(connections) {
		if slices.ContainsFunc(connection, func(in string) bool {
			return strings.HasPrefix(in, "t")
		}) {
			part1Total++
		}
	}
	fmt.Printf("Part1: %d\n", part1Total)

	allConnected := findLargestClique(connections)
	var flattened []string
	for _, conn := range allConnected {
		flattened = append(flattened, conn...)
	}
	slices.Sort(flattened)
	fmt.Println(strings.Join(flattened, ","))
}

func findThreeway(connections map[string][]string) (result [][]string) {
	visited := make(map[string]bool)

	for computer := range connections {
		for _, firstConnection := range connections[computer] {
			for _, secondConnection := range connections[firstConnection] {
				if computer != secondConnection {
					for _, thirdConnection := range connections[secondConnection] {
						conns := []string{computer, firstConnection, secondConnection}
						slices.Sort(conns)
						if thirdConnection == computer && !visited[strings.Join(conns, "-")] {
							result = append(result, conns)
							visited[strings.Join(conns, "-")] = true
						}
					}
				}
			}
		}
	}
	return
}

func findLargestClique(allConnections map[string][]string) [][]string {
	var largestClique [][]string
	connections := make([]string, 0, len(allConnections))
	for connection := range allConnections {
		connections = append(connections, connection)
	}

	for i := 0; i < len(connections); i++ {
		for j := i + 1; j < len(connections); j++ {
			clique := []string{connections[i], connections[j]}
			if isClique(clique, allConnections) {
				for k := j + 1; k < len(connections); k++ {
					clique = append(clique, connections[k])
					if !isClique(clique, allConnections) {
						clique = clique[:len(clique)-1]
					}
				}
				if len(clique) > len(largestClique) {
					largestClique = make([][]string, len(clique))
					for idx, computer := range clique {
						largestClique[idx] = []string{computer}
					}
				}
			}
		}
	}

	return largestClique
}

func isClique(clique []string, allConnections map[string][]string) bool {
	for i := 0; i < len(clique); i++ {
		for j := i + 1; j < len(clique); j++ {
			if !isConnected(clique[i], clique[j], allConnections) {
				return false
			}
		}
	}
	return true
}

func isConnected(computer1, computer2 string, allConnections map[string][]string) bool {
	for _, neighbor := range allConnections[computer1] {
		if neighbor == computer2 {
			return true
		}
	}
	return false
}