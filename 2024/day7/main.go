package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Queue struct {
	internal.Queue[int]
}

var (
	eqs = make(map[int]*Queue)
)

func main() {
	lines := internal.Reader()

	var totalP1, totalP2 int
	for _, line := range lines {
		l := strings.Split(line, ": ")
		total, _ := strconv.Atoi(l[0])
		eqs[total] = &Queue{}
		ll := strings.Split(l[1], " ")
		for _, nm := range ll {
			nmbrs, _ := strconv.Atoi(nm)
			eqs[total].Push(nmbrs)
		}
	}
	for key, v := range eqs {
		startAcc := v.Shift()
		p1 := *v
		p2 := *v
		validP1, validP2 := checkP1(key, startAcc, &p1), checkP2(key, startAcc, &p2)
		if validP1 {
			totalP1 += key
		}
		if validP2 {
			totalP2 += key
		}
	}
	fmt.Printf("Part 1: %d\n", totalP1)
	fmt.Printf("Part 2: %d\n", totalP2)
}

func checkP1(total, acc int, nmbrs *Queue) bool {
	if len(nmbrs.Elements) == 0 {
		return total == acc
	}

	current := nmbrs.Shift()
	a := *nmbrs
	b := *nmbrs
	var mult, add bool
	if current*acc <= total {
		mult = checkP1(total, current*acc, &a)
	}
	if current+acc <= total {
		add = checkP1(total, current+acc, &b)
	}
	return mult || add
}

func checkP2(total, acc int, nmbrs *Queue) bool {
	if len(nmbrs.Elements) == 0 {
		return total == acc
	}

	current := nmbrs.Shift()
	a := *nmbrs
	b := *nmbrs
	c := *nmbrs
	var mult, add, conc bool
	if current*acc <= total {
		mult = checkP2(total, current*acc, &a)
	}
	if current+acc <= total {
		add = checkP2(total, current+acc, &b)
	}
	if concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", acc, current)); concat <= total {
		conc = checkP2(total, concat, &c)
	}
	return mult || add || conc
}
