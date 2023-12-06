package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Catalog struct {
	srcRange, destRange, sizeRange []int
}

var (
	seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHum, humToLoc *Catalog
	seeds                                                                               []int
)

func init() {
	seedToSoil = &Catalog{}
	soilToFert = &Catalog{}
	fertToWater = &Catalog{}
	waterToLight = &Catalog{}
	lightToTemp = &Catalog{}
	tempToHum = &Catalog{}
	humToLoc = &Catalog{}
}

func main() {
	lines := internal.Reader()
	sds := lines[0]
	nmbrs := strings.Split(sds, " ")
	for _, n := range nmbrs {
		if tn, err := strconv.Atoi(strings.TrimSpace(n)); err == nil {
			seeds = append(seeds, tn)
		}
	}

	lines = lines[1:]
	startTime := time.Now()
	var mapFilling string
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			mapFilling = line
		} else {
			l := strings.Split(line, " ")
			dest, _ := strconv.Atoi(l[0])
			src, _ := strconv.Atoi(l[1])
			r, _ := strconv.Atoi(l[2])
			switch mapFilling {
			case "seed-to-soil map:":
				mapper(dest, src, r, seedToSoil)
			case "soil-to-fertilizer map:":
				mapper(dest, src, r, soilToFert)
			case "fertilizer-to-water map:":
				mapper(dest, src, r, fertToWater)
			case "water-to-light map:":
				mapper(dest, src, r, waterToLight)
			case "light-to-temperature map:":
				mapper(dest, src, r, lightToTemp)
			case "temperature-to-humidity map:":
				mapper(dest, src, r, tempToHum)
			case "humidity-to-location map:":
				mapper(dest, src, r, humToLoc)
			}
		}
	}
	doneInit := time.Now()
	fmt.Printf("Done doing the initialization, took: %v\n", doneInit.Sub(startTime))
	startTime = time.Now()
	part1(seeds)
	endTime := time.Now()
	fmt.Printf("Done doing part 1, took: %v\n", endTime.Sub(startTime))
	startTime = endTime
	part2(seeds)
	endTime = time.Now()
	fmt.Printf("Done doing part 2, took: %v\n", endTime.Sub(startTime))
}

func part1(seeds []int) {
	var locations []int
	for _, seed := range seeds {
		soil := getDest(seed, seedToSoil)
		fert := getDest(soil, soilToFert)
		water := getDest(fert, fertToWater)
		light := getDest(water, waterToLight)
		temp := getDest(light, lightToTemp)
		hum := getDest(temp, tempToHum)
		loc := getDest(hum, humToLoc)
		locations = append(locations, loc)
	}
	slices.Sort[[]int](locations)
	fmt.Printf("Solution for Part1: %d\n", locations[0])
}

func part2(seedPairs []int) {
	minLocation := -1

	for i := 8; i < len(seedPairs); i += 2 {
		seeds := make([]int, seedPairs[i+1])
		j := 0
		startTime := time.Now()
		fmt.Printf("Generating seeds for Part2 iteration: %d/%d\n", i, len(seedPairs))
		for k := seedPairs[i]; k < seedPairs[i]+seedPairs[i+1]; k++ {
			seeds[j] = k
			j++
		}

		endTime := time.Now()
		fmt.Printf("Done generating seeds for Part2 iteration %d/%d. Took %v\n", i, len(seedPairs), endTime.Sub(startTime))
		startTime = endTime
		for _, seed := range seeds {
			soil := getDest(seed, seedToSoil)
			fert := getDest(soil, soilToFert)
			water := getDest(fert, fertToWater)
			light := getDest(water, waterToLight)
			temp := getDest(light, lightToTemp)
			hum := getDest(temp, tempToHum)
			loc := getDest(hum, humToLoc)
			if minLocation == -1 || loc < minLocation {
				minLocation = loc
			}
		}
		endTime = time.Now()
		fmt.Printf("Done checking seeds for Part2 iteration %d/%d. Took %v\n", i, len(seedPairs), endTime.Sub(startTime))
	}
	fmt.Printf("Solution for Part2: %d\n", int(minLocation))
}

func mapper(dest, src, r int, current *Catalog) {
	current.destRange = append(current.destRange, dest)
	current.srcRange = append(current.srcRange, src)
	current.sizeRange = append(current.sizeRange, r)
}

func getDest(src int, c *Catalog) int {
	for i := range c.srcRange {
		if c.srcRange[i] <= src && src < c.srcRange[i]+c.sizeRange[i] {
			d := src - c.srcRange[i]
			return c.destRange[i] + d
		}
	}
	return src
}
