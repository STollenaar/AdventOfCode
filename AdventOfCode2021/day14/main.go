package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type Pair struct {
	pair, creationChar string
	count              int
}

type Pairs []*Pair

func (p *Pairs) addPair(pair *Pair) {
	*p = append(*p, pair)
}

func (p *Pairs) createPair(pair *Pair) []Pair {
	pairChars := strings.Split(pair.pair, "")

	pair1 := pairChars[0] + pair.creationChar
	pair2 := pair.creationChar + pairChars[1]

	pairs1 := Pair{pair: pair1, count: pair.count}
	pairs2 := Pair{pair: pair2, count: pair.count}
	pairReduce := *pair
	pairReduce.count = -pair.count

	return []Pair{pairs1, pairs2, pairReduce}
}

// Get amount of character occurences for all pairs
func (p *Pairs) getOccurences(char string) (occurences int) {
	oc := make(map[string]int)

	for _, pair := range *p {
		elements := strings.Split(pair.pair, "")
		oc[elements[0]]++
		oc[elements[1]]++
	}
	return oc[char]
}

func (p *Pairs) getMinElement() (occurence int) {
	elements := p.getElementsAmount()

	occurence = math.MaxInt32
	for _, el := range elements {
		if el < occurence {
			occurence = el
		}
	}
	return occurence
}

func (p *Pairs) getElementsAmount() (elements map[string]int) {
	elements = make(map[string]int)
	elementOccurences := make(map[string]int)
	for _, pair := range *p {
		pairChars := strings.Split(pair.pair, "")
		elements[pairChars[0]] += pair.count
		elements[pairChars[1]] += pair.count

		if elementOccurences[pairChars[0]] == 0 {
			elementOccurences[pairChars[0]] = p.getOccurences(pairChars[0])
		}
		if elementOccurences[pairChars[1]] == 0 {
			elementOccurences[pairChars[1]] = p.getOccurences(pairChars[1])
		}
	}
	for key := range elements {
		elements[key] /= elementOccurences[key]
	}

	return elements
}

func (p *Pairs) getMaxElement() (occurence int) {
	elements := p.getElementsAmount()

	for _, el := range elements {
		if el > occurence {
			occurence = el
		}
	}
	return occurence
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// var polymer []string
	pairs := new(Pairs)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "->") {
			insertions := strings.Split(line, " -> ")

			if pair := inSlice(pairs, insertions[0]); pair != nil && pair.creationChar == "" {
				pair.creationChar = insertions[1]
			} else if pair == nil {
				pairs.addPair(&Pair{pair: insertions[0], creationChar: insertions[1]})
			}
		} else {
			polymer := strings.Split(line, "")
			for i := 0; i < len(polymer)-1; i++ {
				if pair := inSlice(pairs, polymer[i]+polymer[i+1]); pair == nil {
					pairs.addPair(&Pair{pair: polymer[i] + polymer[i+1], count: 1})
				} else {
					pair.count++
				}
			}
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	for steps := 0; steps < 40; steps++ {
		var changes []Pair
		for i, pair := range *pairs {
			if pair.count > 0 {
				changes = append(changes, pairs.createPair((*pairs)[i])...)
			}
		}
		for _, pair := range changes {
			inSlice(pairs, pair.pair).count += pair.count
		}

		if steps == 9 {
			mostCommon := pairs.getMaxElement()
			leastCommon := pairs.getMinElement()
			elapsed := time.Since(start)
			fmt.Println("Execution time for part 1: ", elapsed)
			fmt.Println("quantity of polymer for par 1: ", mostCommon-leastCommon)
		}
	}
	mostCommon := pairs.getMaxElement()
	leastCommon := pairs.getMinElement()
	elapsed := time.Since(start)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("quantity of polymer for par 2: ", mostCommon-leastCommon)

}

func inSlice(slice *Pairs, pair string) *Pair {
	for _, el := range *slice {
		if el.pair == pair {
			return el
		}
	}
	return nil
}
