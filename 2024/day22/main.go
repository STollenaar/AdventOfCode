package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

func main() {
	lines := internal.Reader()

	var total int
	seq := make(map[int][]int)
	seqOnes := make(map[int][]int)
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		initial := secret
		secretOne := ones(secret)
		seqOnes[initial] = append(seqOnes[initial], secretOne)
		seq[initial] = append(seq[initial], 0)
		for i := 0; i < 2000; i++ {
			secret = sequence(secret)
			nextOne := ones(secret)
			delta := nextOne - secretOne
			seq[initial] = append(seq[initial], delta)
			seqOnes[initial] = append(seqOnes[initial], nextOne)
			secretOne = nextOne
		}
		total += secret
	}
	fmt.Printf("Part 1: %d\n", total)
	maxSeq := findMaxSeq(seqOnes, seq)
	fmt.Printf("Part 2: %d\n", addAll(maxSeq))
}

func sequence(secret int) int {
	secret = prune(mix(secret, secret*64))
	secret = prune(mix(secret, secret/32))
	secret = prune(mix(secret, secret*2048))
	return secret
}

func prune(secret int) int {
	return secret % 16777216
}

func mix(secret, value int) int {
	return secret ^ value
}

func ones(nmbr int) int {
	s := strconv.Itoa(nmbr)
	n, _ := strconv.Atoi(string(s[len(s)-1]))
	return n
}

func findMaxSeq(seqOnes map[int][]int, seqDeltas map[int][]int) []int {
	occ := make(map[string][]int)

	for nmbr, ones := range seqOnes {
		delta := seqDeltas[nmbr]
		seq := make([]int, 4)
		occTemp := make(map[string][]int)
		for d := 3; d < len(delta); d++ {
			seq[0] = delta[d-3]
			seq[1] = delta[d-2]
			seq[2] = delta[d-1]
			seq[3] = delta[d]

			occTemp[sliceToString(seq)] = append(occTemp[sliceToString(seq)], ones[d])
		}
		for k, v := range occTemp {
			occ[k] = append(occ[k], v[0])
		}
	}
	var max int
	var maxSeq string

	for k, v := range occ {
		count := addAll(v)
		if max < count {
			max = count
			maxSeq = k
		}
	}
	return occ[maxSeq]
}

func sliceToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}

func addAll(a []int) (out int) {
	for _, t := range a {
		out += t
	}
	return
}
