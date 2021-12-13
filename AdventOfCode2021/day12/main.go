package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Node struct {
	isBig     bool
	label     string
	connected []*Node
}

func inSlice(slice []*Node, node *Node) bool {
	for _, val := range slice {
		if val.label == node.label {
			return true
		}
	}
	return false
}

func amountInSlice(slice []*Node, node *Node) (amount int) {
	for _, val := range slice {
		if val.label == node.label {
			amount++
		}
	}

	return amount
}

func getHighestCountSmallCaves(slice []*Node) (amount int) {
	var counted []*Node
	for _, val := range slice {
		if !val.isBig && !inSlice(counted, val) {
			counted = append(counted, val)
			valAmount := amountInSlice(slice, val)
			if valAmount > amount {
				amount = valAmount
			}
		}
	}
	return amount
}

func getNode(label string, nodes []*Node) *Node {
	for _, node := range nodes {
		if node.label == label {
			return node
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

	var nodes []*Node

	for scanner.Scan() {
		line := scanner.Text()

		points := strings.Split(line, "-")

		node1 := getNode(points[0], nodes)
		node2 := getNode(points[1], nodes)

		if node1 == nil {
			node1 = &Node{label: points[0], isBig: strings.ToUpper(points[0]) == points[0]}
		}

		if node2 == nil {
			node2 = &Node{label: points[1], isBig: strings.ToUpper(points[1]) == points[1]}
		}

		node1.connected = append(node1.connected, node2)
		node2.connected = append(node2.connected, node1)

		nodes = append(nodes, node1, node2)
	}
	doPart1(nodes, start)
	doPart2(nodes, start)
}

func doPart1(nodes []*Node, start time.Time) {
	startNode := getNode("start", nodes)
	if startNode == nil {
		log.Fatal("CANNOT FIND START NODE")
	}
	var total *int
	total = new(int)

	totalPaths := traversePart1(startNode, []*Node{}, total)

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Total paths possible for part 1: ", *totalPaths)
}

func traversePart1(node *Node, visited []*Node, ends *int) *int {
	visited = append(visited, node)

	if node.label == "end" {
		*ends++
		return ends
	}
	for _, connected := range node.connected {
		wasVisited := inSlice(visited, connected)
		if (wasVisited && connected.isBig) || !wasVisited {
			traversePart1(connected, visited, ends)
		}
	}
	return ends
}

func doPart2(nodes []*Node, start time.Time) {
	startNode := getNode("start", nodes)
	if startNode == nil {
		log.Fatal("CANNOT FIND START NODE")
	}
	var total *int
	total = new(int)

	totalPaths := traversePart2(startNode, []*Node{}, total)

	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Total paths possible for part 2: ", *totalPaths)
}

func traversePart2(node *Node, visited []*Node, ends *int) *int {
	visited = append(visited, node)

	if node.label == "end" {
		*ends++
		return ends
	}
	for _, connected := range node.connected {
		if connected.label == "start" {
			continue
		}
		wasVisited := inSlice(visited, connected)
		if !wasVisited || (wasVisited && connected.isBig) || (wasVisited && !connected.isBig && getHighestCountSmallCaves(visited) < 2) {
			traversePart2(connected, visited, ends)
		}
	}
	return ends
}
