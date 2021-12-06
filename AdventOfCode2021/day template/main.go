package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Execution time for part : ", elapsed)

}
