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
	X, Y, Z, ID int
}

type Edge struct {
	A, B     *Node
	Distance float64
}

var nodes []*Node

func main() {
	lines := internal.Reader()

	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		nodes = append(nodes, &Node{
			X:  x,
			Y:  y,
			Z:  z,
			ID: i,
		})
	}

	edges := generateEdges(nodes)

	circuits := clusterByProximity(edges)
	X := 3
	part1 := 1
	for i := 0; i < X && i < len(circuits); i++ {
		fmt.Printf("Circuit %d has %d junctions\n", i+1, circuits[i])
		part1 *= circuits[i]
	}
	fmt.Printf("Part1: %d\n", part1)

	last := clusterAll(edges)
	fmt.Printf("Part2: %d\n", last.A.X*last.B.X)
}

func generateEdges(nodes []*Node) (edges []*Edge) {

	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			edges = append(edges, &Edge{
				A:        nodes[i],
				B:        nodes[j],
				Distance: distance(nodes[i], nodes[j]),
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Distance < edges[j].Distance
	})
	return
}

func distance(a, b *Node) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}
func clusterByProximity(edges []*Edge) []int {
	// DSU MUST contain every node, not just those in pairs
	dsu := NewDSU(nodes)

	// Union all pairs
	limit := min(1000, len(edges))
	for i := 0; i < limit; i++ {
		dsu.Union(edges[i].A, edges[i].B)
	}

	circuitMap := make(map[int]int)
	for i := 0; i < len(nodes); i++ {
		root := dsu.Find(i)
		circuitMap[root.ID]++
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

func clusterAll(edges []*Edge) *Edge {
	dsu := NewDSU(nodes)
	var lastMerged *Edge

	for _, e := range edges {
		rootA := dsu.Find(e.A.ID)
		rootB := dsu.Find(e.B.ID)
		if rootA != rootB {
			dsu.Union(e.A, e.B)
			lastMerged = e
		}
	}

	return lastMerged
}

type DSU struct {
	parent map[int]*Node
	size   map[int]int
}

func NewDSU(nodes []*Node) *DSU {
	parent := make(map[int]*Node)
	size := make(map[int]int)

	for _, n := range nodes {
		parent[n.ID] = n
		size[n.ID] = 1
	}

	return &DSU{parent, size}
}

func (d *DSU) Find(id int) *Node {
	if d.parent[id].ID != id {
		d.parent[id] = d.Find(d.parent[id].ID)
	}
	return d.parent[id]
}

func (d *DSU) Union(a, b *Node) {
	ra := d.Find(a.ID)
	rb := d.Find(b.ID)

	if ra == rb {
		return
	}

	// Union by size
	if d.size[ra.ID] < d.size[rb.ID] {
		ra, rb = rb, ra
	}

	d.parent[rb.ID] = ra
	d.size[ra.ID] += d.size[rb.ID]
}

func (d *DSU) ComponentSizes() map[int]int {
	result := make(map[int]int)
	for n := range d.parent {
		r := d.Find(n)
		result[r.ID] = d.size[r.ID]
	}
	return result
}
