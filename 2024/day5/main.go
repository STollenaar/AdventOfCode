package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	pages map[int][]int
)

func init() {
	pages = make(map[int][]int)
}

func main() {
	lines := internal.Reader()
	initializing := true
	var valid, invalid [][]int
	for _, line := range lines {
		if line == "" {
			initializing = false
			continue
		}

		if initializing {
			s := strings.Split(line, "|")
			l, _ := strconv.Atoi(s[0])
			r, _ := strconv.Atoi(s[1])
			pages[l] = append(pages[l], r)
		} else {
			ps := strings.Split(line, ",")
			p := toIntSlice(ps)
			var temp []int
			for i, n := range p {
				if i == 0 {
					temp = append(temp, n)
				} else {
					if sliceContainsElems(temp, pages[n]) {
						invalid = append(invalid, p)
						break
					}else{
						temp = append(temp, n)
					}
				}
			}
			if len(temp) == len(p) {
				valid = append(valid, p)
			}
		}
	}
	var totalPart1, totalPart2 int
	for _, v := range valid {
		totalPart1 += v[len(v)/2]
	}
	fmt.Printf("Part 1: %d\n", totalPart1)

	for _, iv := range invalid {
		fixed := fixInvalid(iv)
		totalPart2 += fixed[len(fixed)/2]
	}
	fmt.Printf("Part 2: %d\n", totalPart2)

}

func toIntSlice(in []string) (out []int) {
	for _, i := range in {
		t, _ := strconv.Atoi(i)
		out = append(out, t)
	}
	return
}

func sliceContainsElems(a, b []int) bool {
	for _, i := range a {
		if slices.Contains(b, i) {
			return true
		}
	}
	return false
}

func fixInvalid(in []int) (out []int){

	for i, n := range in {
		if i == 0 {
			out = append(out, n)
		}else{
			var inserted bool
			for j:=0;j<len(out);j++ {
				if slices.Contains(pages[n], out[j]) {
					tmp := make([]int, len(out[j:]))
					copy(tmp, out[j:])
					out = append(out[:j], n)
					out = append(out, tmp...)
					inserted = true
					break
				}
			}
			if !inserted {
				out = append(out, n)
			}
		}
	}
	return
}