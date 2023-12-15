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

		if err != nil {
			log.Fatal(err)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Execution time for part : ", elapsed)

}
