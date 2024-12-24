package main

import (
	"fmt"
	"maps"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
	"github.com/emicklei/dot"
)

type Gate struct {
	a, b, out, kind string
}

var (
	cables = make(map[string]bool)
	gates  []Gate

	graph   = make(map[string]dot.Node)
)

func main() {
	lines := internal.Reader()
	g := dot.NewGraph(dot.Directed)

	cableInit := true
	for _, line := range lines {
		if line == "" {
			cableInit = false
			continue
		}
		if cableInit {
			l := strings.Split(line, ": ")
			cables[l[0]] = l[1] == "1"
			graph[l[0]] = g.Node(l[0]).Box()
		} else {
			l := strings.Split(line, " ")
			gate := Gate{
				a:    l[0],
				b:    l[2],
				out:  l[4],
				kind: l[1],
			}
			gates = append(gates, gate)
			n :=g.Node(fmt.Sprintf("%s -> %s",gate.kind,gate.out))
			graph[gate.out] = n
		}
	}
	for _, gate := range gates {
		g.Edge(graph[gate.a], graph[gate.out])
		g.Edge(graph[gate.b], graph[gate.out])
	}
	err := os.WriteFile("graph.dot", []byte(g.String()), 0644)
	if err != nil {
		panic(err)
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
