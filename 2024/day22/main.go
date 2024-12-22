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
	fmt.Println(total)
	fmt.Println(seq)
	fmt.Println(findMaxSeq(seq[123]))
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

func findMaxSeq(delta []int) []int {
	seq := make([]int, 4)
	occ := make(map[string]int)

	for d := 3; d < len(delta); d++ {
		seq[0] = delta[d-3]
		seq[1] = delta[d-2]
		seq[2] = delta[d-1]
		seq[3] = delta[d]

		occ[sliceToString(seq)]++
	}

	var max int
	var maxKey string
	fmt.Println(occ["-2,1,-1,3"])
	for k, v := range occ {
		if max < v {
			maxKey = k
			max = v
		}
	}
	return stringToSlice(maxKey)
}

func sliceToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}

func stringToSlice(in string) (out []int) {
	a := strings.Split(in, ",")
	for _, i := range a {
		t, _ := strconv.Atoi(i)
		out = append(out, t)
	}
	return
}
