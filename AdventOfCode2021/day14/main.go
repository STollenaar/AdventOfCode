package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type InsertionRule struct {
	par1, par2, insertChar string
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var polymer []string
	var insertionRules []InsertionRule
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "->") {
			insertions := strings.Split(line, " -> ")
			searchPar := strings.Split(insertions[0], "")
			insertionRules = append(insertionRules, InsertionRule{par1: searchPar[0], par2: searchPar[1], insertChar: insertions[1]})
		} else {
			polymer = strings.Split(line, "")
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	for steps := 0; steps < 40; steps++ {
		fmt.Println("Doing step: ", steps)
		polymerCopy := new(sync.Map)
		partitions := chunks(polymer, 100000)
		fmt.Println("Looping through partitions on step: ", steps)
		// var waitGroup sync.WaitGroup
		for p, partition := range partitions {
			// waitGroup.Add(1)
			// go func(polymerCopy *sync.Map, p int, partition []string, waitGroup *sync.WaitGroup) {
			// 	defer waitGroup.Done()
			for i := 0; i < len(partition)-1; i++ {
				slice, _ := polymerCopy.LoadOrStore(p, []string{})

				slice = append(slice.([]string), partition[i])
				polymerCopy.Store(p, slice)
				for _, insertionRule := range insertionRules {
					if partition[i] == insertionRule.par1 && partition[i+1] == insertionRule.par2 {
						slice = append(slice.([]string), insertionRule.insertChar)

						polymerCopy.Store(p, slice)
						break
					}
				}
			}
			// }(polymerCopy, p, partition, &waitGroup)
		}
		// waitGroup.Wait()
		var polCopy []string

		fmt.Println("Merging into polymer slice for step: ", steps)
		for i := 0; i < len(partitions); i++ {
			slice, _ := polymerCopy.Load(i)
			polCopy = append(polCopy, slice.([]string)...)
		}
		polymer = polCopy

		if steps == 9 {
			mostCommon := getMostCommonAmount(polymer)
			leastCommon := getLeastCommonAmount(polymer, mostCommon)
			elapsed := time.Since(start)
			fmt.Println("Execution time for part 1: ", elapsed)
			fmt.Println("quantity of polymer for par 1: ", mostCommon-leastCommon)
		}
	}
	mostCommon := getMostCommonAmount(polymer)
	leastCommon := getLeastCommonAmount(polymer, mostCommon)
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("quantity of polymer for par 2: ", mostCommon-leastCommon)

}

func chunks(xs []string, chunkSize int) [][]string {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][]string, (len(xs)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(xs) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = xs[prev : next+1]
		prev = next
		i++
	}
	divided[i] = xs[prev:]
	return divided
}

func getMostCommonAmount(slice []string) (amount int) {
	var seenElements []string
	for _, el := range slice {
		if !inSlice(seenElements, el) {
			occurences := getOccurences(slice, el)
			if occurences > amount {
				amount = occurences
			}
			seenElements = append(seenElements, el)
		}
	}
	return amount
}

func getLeastCommonAmount(slice []string, mostCommon int) (amount int) {
	amount = mostCommon
	var seenElements []string
	for _, el := range slice {
		if !inSlice(seenElements, el) {
			occurences := getOccurences(slice, el)
			if occurences < amount {
				amount = occurences
			}
			seenElements = append(seenElements, el)
		}
	}
	return amount
}

func inSlice(slice []string, par string) bool {
	for _, el := range slice {
		if el == par {
			return true
		}
	}
	return false
}

func getOccurences(slice []string, par string) (amount int) {
	for _, el := range slice {
		if el == par {
			amount++
		}
	}
	return amount
}
