package main

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Gate struct {
	a, b, out, kind string
}

var (
	cables = make(map[string]bool)
	gates  []Gate
)

func main() {
	lines := internal.Reader()

	cableInit := true
	for _, line := range lines {
		if line == "" {
			cableInit = false
			continue
		}
		if cableInit {
			l := strings.Split(line, ": ")
			cables[l[0]] = l[1] == "1"
		} else {
			l := strings.Split(line, " ")
			gate := Gate{
				a:    l[0],
				b:    l[2],
				out:  l[4],
				kind: l[1],
			}
			gates = append(gates, gate)
		}
	}
	for {
		before := make(map[string]bool)
		maps.Copy(before, cables)
		for _, gate := range gates {
			gate.DoGate()
		}
		if reflect.DeepEqual(before, cables) {
			break
		}
	}

	var outputs []string
	for k := range cables {
		if strings.HasPrefix(k, "z") {
			outputs = append(outputs, k)
		}
	}
	slices.Sort(outputs)
	slices.Reverse(outputs)
	var binaryString string
	for _, output := range outputs {
		if cables[output] {
			binaryString += "1"
		} else {
			binaryString += "0"
		}
	}
	i, _ := strconv.ParseInt(binaryString, 2, 64)
	fmt.Printf("Part1: %d\n", i)
}

func (g *Gate) DoGate() {
	switch g.kind {
	case "AND":
		cables[g.out] = cables[g.a] && cables[g.b]
	case "OR":
		cables[g.out] = cables[g.a] || cables[g.b]
	case "XOR":
		cables[g.out] = cables[g.a] != cables[g.b]
	}
}
