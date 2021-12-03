package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BitOccurences struct {
	oneOc  int
	zeroOc int
}

type Bit struct {
	bit []string
}

var bitOccurences [12]BitOccurences // For part 1
var bits []Bit

// Gettomg character needed for filter
func getCharacter(bitOccurence BitOccurences, mostSignificant bool) (character string) {
	character = "1"
	if mostSignificant && bitOccurence.zeroOc > bitOccurence.oneOc {
		character = "0"
	} else if !mostSignificant && bitOccurence.zeroOc <= bitOccurence.oneOc {
		character = "0"
	}
	return character
}

// Get the criteria
func getCriteria(bits []Bit, mostSignificant bool, index int) (filteredBits []Bit) {
	var bitOccurence BitOccurences

	for _, bit := range bits {
		if bit.bit[index] == "0" {
			bitOccurence.zeroOc++
		} else {
			bitOccurence.oneOc++
		}
	}

	character := getCharacter(bitOccurence, mostSignificant)

	filteredBits = filterSlice(bits, character, index)
	if len(filteredBits) != 1 {
		filteredBits = getCriteria(filteredBits, mostSignificant, (index + 1))
	}
	return filteredBits
}

// Filter bits with for only character at index
func filterSlice(bits []Bit, character string, index int) (filtered []Bit) {
	for _, bit := range bits {
		if bit.bit[index] == character {
			filtered = append(filtered, bit)
		}
	}
	return filtered
}

func binaryToInt(binary string) int64 {
	inter, _ := strconv.ParseInt(binary, 2, 64)
	return inter
}

// Getting occurences of all
func getGammeEpsilon(bitOccurences [12]BitOccurences) (int64, int64) {
	gammaString, epsilonString := "", ""

	for _, bit := range bitOccurences {
		if bit.oneOc > bit.zeroOc {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}
	gamma := binaryToInt(gammaString)
	epsilon := binaryToInt(epsilonString)
	return gamma, epsilon
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		bits = append(bits, Bit{line})
		for index, bit := range line {
			if bit == "0" {
				bitOccurences[index].zeroOc++
			} else {
				bitOccurences[index].oneOc++
			}
		}
	}
	gamma, epsilon := getGammeEpsilon(bitOccurences)
	fmt.Println(gamma, " ", epsilon, " ", gamma*epsilon)

	O2Rate := binaryToInt(strings.Join(getCriteria(bits, true, 0)[0].bit, ""))
	CO2Rate := binaryToInt(strings.Join(getCriteria(bits, false, 0)[0].bit, ""))
	fmt.Println(O2Rate, " ", CO2Rate, " ", O2Rate*CO2Rate)

}
