package main

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
	"time"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	stones = make(map[string]int)
)

func main() {
	lines := internal.Reader()
	line := strings.Split(lines[0], " ")
	start := time.Now()

	for _, l := range line {
		stones[l]++
	}
	for i := 0; i < 25; i++ {
		blink()
	}
	fmt.Printf("Part 1: %d, Duration %v\n", getTotal(), time.Since(start))
	start = time.Now()
	for i := 0; i < 50; i++ {
		blink()
	}
	fmt.Printf("Part 2: %d, Duration %v\n", getTotal(), time.Since(start))
}

func blink() {
	tmp := make(map[string]int)
	maps.Copy(tmp, stones)
	for stone, v := range tmp {
		if stone == "0" {
			stones["1"] += v
		} else if len(stone)%2 == 0 {
			mid := len(stone) / 2
			left, right := stone[:mid], stone[mid:]
			lt, _ := strconv.Atoi(left)
			rt, _ := strconv.Atoi(right)
			stones[strconv.Itoa(lt)] += v
			stones[strconv.Itoa(rt)] += v
		} else {
			n, _ := strconv.Atoi(stone)
			n *= 2024
			stones[strconv.Itoa(n)] += v
		}
		stones[stone] -= v
	}
	return
}

func getTotal() (total int) {
	for _, v := range stones {
		total += v
	}
	return
}
