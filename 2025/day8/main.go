package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Node struct {
	X, Y, Z int
}

type Pair struct {
	A, B     *Node
	Distance float64
}

var nodes []*Node

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		nodes = append(nodes, &Node{
			X: x,
			Y: y,
			Z: z,
		})
	}

	pairs := computeClosestPairs(nodes)
	// for _, p := range pailrs {
	// 	fmt.Printf("Nodes %d & %d => %d\n", p.A, p.B, p.Distance)
	// }

	circuits := clusterByProximity(pairs)
	X := 3
	for i := 0; i < X && i < len(circuits); i++ {
		fmt.Printf("Circuit %d has %d junctions\n", i+1, circuits[i])
	}
}

func computeClosestPairs(nodes []*Node) []*Pair {
	pairs := make([]*Pair, 0, len(nodes))

	for _, a := range nodes {
		var best *Node
		minDist := math.MaxFloat64

		for _, b := range nodes {
			if a == b {
				continue
			}
			d := distance(a, b)
			if d < minDist {
				minDist = d
				best = b
			}
		}
		if !contains(pairs, a, best) {
			pairs = append(pairs, &Pair{
				A:        a,
				B:        best,
				Distance: minDist,
			})
		}
	}

	return pairs
}

func contains(pairs []*Pair, a, b *Node) bool {
	for _, pair := range pairs {
		if pair.A == b && pair.B == a {
			return true
		}
	}
	return false
}

func distance(a, b *Node) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func clusterByProximity(pairs []*Pair) []int {
	// DSU MUST contain every node, not just those in pairs
	dsu := NewDSU(nodes)

	// Union all pairs
	for _, p := range pairs {
		dsu.Union(p.A, p.B)
	}

	// Count components
	sizes := dsu.ComponentSizes()

	// Convert & sort
	var circuits []int
	for _, size := range sizes {
		circuits = append(circuits, size)
	}

	sort.Slice(circuits, func(i, j int) bool { return circuits[i] > circuits[j] })
	return circuits
}

type DSU struct {
	parent map[*Node]*Node
	size   map[*Node]int
}

func NewDSU(nodes []*Node) *DSU {
	parent := make(map[*Node]*Node)
	size := make(map[*Node]int)

	for _, n := range nodes {
		parent[n] = n
		size[n] = 1
	}

	return &DSU{parent, size}
}

func (d *DSU) Find(n *Node) *Node {
	if d.parent[n] != n {
		d.parent[n] = d.Find(d.parent[n])
	}
	return d.parent[n]
}

func (d *DSU) Union(a, b *Node) {
	ra := d.Find(a)
	rb := d.Find(b)

	if ra == rb {
		return
	}

	// Union by size
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}

	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
}

func (d *DSU) ComponentSizes() map[*Node]int {
	result := make(map[*Node]int)
	for n := range d.parent {
		r := d.Find(n)
		result[r] = d.size[r]
	}
	return result
}
