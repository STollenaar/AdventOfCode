package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func filterFromSlice(a []string, b []string) (output []string) {
	for _, i := range a {
		exist := false
		for _, j := range b {
			if i == j {
				exist = true
				break
			}
		}
		if !exist {
			output = append(output, i)
		}
	}
	return output
}

func findElementContains(slice []string, values []string) string {
	for _, element := range slice {
		contains := true
		for _, value := range values {
			if !strings.Contains(element, value) {
				contains = false
				break
			}
		}
		if contains {
			return element
		}
	}
	return ""
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

type DigitLarge struct {
	top         []string
	topLeft     []string
	topRight    []string
	middle      []string
	bottom      []string
	bottomLeft  []string
	bottomRight []string

	inputs []string
}

type Input struct {
	patterns []string
	digits   []string
	decoder  *DigitLarge
}

func main() {
	start := time.Now()

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var inputs []*Input
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")

		noise := strings.Split(strings.TrimSpace(line[0]), " ")
		digit := strings.Split(strings.TrimSpace(line[1]), " ")

		inputs = append(inputs, &Input{noise, digit, &DigitLarge{}})
	}
	doPart1(inputs, start)
	doPart2(inputs, start)
}

func doPart1(inputs []*Input, startTime time.Time) {
	occurences := 0
	for _, input := range inputs {
		for _, digit := range input.digits {
			switch len(digit) {
			case 2:
				occurences++
			case 4:
				occurences++
			case 3:
				occurences++
			case 7:
				occurences++
			}
		}
	}

	elapsed := time.Since(startTime)
	fmt.Println("Execution time for part 1: ", elapsed)
	fmt.Println("Occurences of 1, 4, 7 or 8: ", occurences)
}

func doPart2(inputs []*Input, startTime time.Time) {
	total := 0
	for _, input := range inputs {
		input.decoder = putDecoder(input.patterns)

		numberPart := ""
		for _, digit := range input.digits {
			numberPart += decodeString(input.decoder, digit)
		}
		number, _ := strconv.Atoi(numberPart)
		total += number
	}
	elapsed := time.Since(startTime)
	fmt.Println("Execution time for part 2: ", elapsed)
	fmt.Println("Total: ", total)
}

func putDecoder(patterns []string) *DigitLarge {
	decoderMap := make(map[int]*DigitLarge)

	for _, pattern := range patterns {
		decoder := decoderMap[len(pattern)]
		if decoder == nil {
			decoder = &DigitLarge{}
			decoderMap[len(pattern)] = decoder
		}
		decoder.inputs = append(decoder.inputs, pattern)
		switch len(pattern) {

		// Digit 1
		case 2:
			decoder.topRight = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.bottomRight = append(decoder.bottomRight, strings.Split(pattern, "")...)
		// Digit 4
		case 4:
			decoder.topLeft = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.topRight = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.middle = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.bottomRight = append(decoder.topRight, strings.Split(pattern, "")...)
		// Digit 7
		case 3:
			decoder.top = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.topRight = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.bottomRight = append(decoder.topRight, strings.Split(pattern, "")...)
		// Digits 2,3,5,6,8,9,0
		default:
			decoder.topLeft = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.top = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.topRight = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.middle = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.bottomLeft = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.bottom = append(decoder.topRight, strings.Split(pattern, "")...)
			decoder.bottomRight = append(decoder.topRight, strings.Split(pattern, "")...)
		}
		decoder.topLeft = unique(decoder.topLeft)
		decoder.top = unique(decoder.top)
		decoder.topRight = unique(decoder.topRight)
		decoder.middle = unique(decoder.middle)
		decoder.bottomLeft = unique(decoder.bottomLeft)
		decoder.bottom = unique(decoder.bottom)
		decoder.bottomRight = unique(decoder.bottomRight)
	}

	for i := 3; i < 8; i++ {
		decoder := decoderMap[i]
		decoder.topRight = decoderMap[2].topRight
		decoder.bottomRight = decoderMap[2].bottomRight

		switch i {
		case 3:
			decoder.top = filterFromSlice(decoder.top, decoderMap[2].topRight)
		case 4:
			decoder.topRight = decoderMap[2].topRight
			decoder.bottomRight = decoderMap[2].bottomRight
			decoder.topLeft = filterFromSlice(decoder.topLeft, decoderMap[2].bottomRight)
			decoder.middle = filterFromSlice(decoder.middle, decoderMap[2].bottomRight)
		case 5:
			{
				decoder.middle = decoderMap[4].middle
				decoder.top = decoderMap[3].top
				decoder.topLeft = decoderMap[4].topLeft

				decoder.bottom = filterFromSlice(decoder.bottom, decoder.top)
				decoder.bottom = filterFromSlice(decoder.bottom, decoder.middle)
				decoder.bottom = filterFromSlice(decoder.bottom, decoder.topRight)
				decoder.bottom = filterFromSlice(decoder.bottom, decoder.bottomRight)

				decoder.bottomLeft = filterFromSlice(decoder.bottomLeft, decoder.top)
				decoder.bottomLeft = filterFromSlice(decoder.bottomLeft, decoder.middle)
				decoder.bottomLeft = filterFromSlice(decoder.bottomLeft, decoder.topRight)
				decoder.bottomLeft = filterFromSlice(decoder.bottomLeft, decoder.bottomRight)
			}
		case 6:
			{
				decoder.middle = decoderMap[4].middle
				decoder.top = decoderMap[3].top
				decoder.topLeft = decoderMap[4].topLeft

				decoder.bottom = filterFromSlice(decoder.bottom, decoder.top)
				decoder.bottom = filterFromSlice(decoder.bottom, decoder.middle)
				decoder.bottom = filterFromSlice(decoder.bottom, decoder.topRight)
				decoder.bottom = filterFromSlice(decoder.bottom, decoder.bottomRight)

				// With help of the 2,3,5 decoder we derive the others
				others := decoderMap[5]

				// Deriving the bottomLeft value from a nine
				nine := findElementContains(decoder.inputs, append(decoder.topLeft, decoder.topRight...))
				decoder.bottomLeft = filterFromSlice(decoder.bottomLeft, strings.Split(nine, ""))
				decoder.bottom = filterFromSlice(decoder.bottom, decoder.bottomLeft)

				// Deriving the missing values for 2,3,5
				two := findElementContains(others.inputs, decoder.bottomLeft)
				// Deriving the bottom right from 2 since it doesn't have it
				decoder.bottomRight = filterFromSlice(decoder.bottomRight, strings.Split(two, ""))

				// Filtering out the bottom right value for the topright
				decoder.topRight = filterFromSlice(decoder.topRight, decoder.bottomRight)

				// Deriving the topleft from 2 since it doesn't have it
				decoder.topLeft = filterFromSlice(decoder.topLeft, strings.Split(two, ""))
				// finally setting the middle
				decoder.middle = filterFromSlice(decoder.middle, decoder.topLeft)
			}
		}
	}
	return decoderMap[6]
}

func decodeString(decoder *DigitLarge, digit string) string {
	switch len(digit) {
	// Digit 1
	case 2:
		return "1"
	// Digit 4
	case 4:
		return "4"
	// Digit 7
	case 3:
		return "7"
	// Digit 8
	case 7:
		return "8"

	// Digits 2,3,5
	case 5:
		{
			// The middle bridge thingy
			if strings.Contains(digit, decoder.topRight[0]) {
				if strings.Contains(digit, decoder.bottomLeft[0]) {
					return "2"
				} else {
					return "3"
				}
			} else {
				return "5"
			}
		}
	// Digit 6,9,0
	case 6:
		{
			// The middle bridge thingy
			if strings.Contains(digit, decoder.middle[0]) {
				if strings.Contains(digit, decoder.topRight[0]) {
					return "9"
				} else {
					return "6"
				}
			} else {
				return "0"
			}
		}
	}
	return "-1"
}
