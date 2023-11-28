package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var horizontal, depth, aim int

func main() {
	horizontal, depth, aim = 0, 0, 0
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		value, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		switch command {
		case "up":
			{
				aim -= value
				break
			}
		case "down":
			{
				aim += value
				break
			}
		case "forward":
			{
				horizontal += value
				depth += aim * value
				break
			}
		}
	}
	// Part 1 print
	fmt.Println((horizontal * aim))

	// Part 2 print
	fmt.Println((horizontal * depth))
}
