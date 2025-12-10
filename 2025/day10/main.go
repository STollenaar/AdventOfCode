package main

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Machine struct {
	endMask   int
	moveMasks []int
	joltages  []int
	bitCount  int
}

type variable struct {
	expr linear
	free bool
	val  int
	max  int
}

const N = 13

type linear struct {
	a [N]float64
	b float64
}

const EPS = 1e-8

var machines []Machine

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		machines = append(machines, ParseMachine(line))
	}
	var totalLightSteps, totalJoltageSteps int
	for _, machine := range machines {
		totalLightSteps += LightSteps(machine)
	}

	fmt.Printf("Part1: %d\n", totalLightSteps)
	for _, machine := range machines {
		totalJoltageSteps += JoltageSteps(machine)
	}
	fmt.Printf("Part2: %d\n", totalJoltageSteps)
}

func LightSteps(m Machine) int {
	start := 0
	target := m.endMask

	if start == target {
		return 0
	}

	queue := list.New()
	queue.PushBack(start)

	visited := map[int]bool{start: true}
	steps := map[int]int{start: 0}

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(int)

		if cur == target {
			return steps[cur]
		}

		for _, mv := range m.moveMasks {
			next := cur ^ mv // XOR flips bits

			if !visited[next] {
				visited[next] = true
				steps[next] = steps[cur] + 1
				queue.PushBack(next)
			}
		}
	}

	return -1
}

// WHY, DO I SUFFER, WHY DO I USE LINEAR ALGEBRA
// WHY, AM I RAMBLING
func JoltageSteps(m Machine) int {
	vars := make([]variable, len(m.moveMasks))
	for i := range vars {
		vars[i].max = math.MaxInt
	}

	eqs := make([]linear, len(m.joltages))
	for i, joltages := range m.joltages {
		eq := linear{b: float64(-joltages)}
		for j, b := range m.moveMasks {
			if b&(1<<i) != 0 {
				eq.a[j] = 1
				vars[j].max = min(vars[j].max, joltages)
			}
		}
		eqs[i] = eq
	}

	for i := range vars {
		vars[i].free = true

		for _, eq := range eqs {
			if expr, ok := extract(eq, i); ok {
				vars[i].free = false
				vars[i].expr = expr

				for j := range eqs {
					eqs[j] = substitute(eqs[j], i, expr)
				}

				break
			}
		}
	}

	free := []int(nil)
	for i, v := range vars {
		if v.free {
			free = append(free, i)
		}
	}

	best, _ := evalRecursive(vars, free, 0)
	return best
}

func extract(lin linear, index int) (linear, bool) {
	a := -lin.a[index]
	if math.Abs(a) < EPS {
		return linear{}, false
	}

	r := linear{b: lin.b / a}
	for i := 0; i < N; i++ {
		if i != index {
			r.a[i] = lin.a[i] / a
		}
	}
	return r, true
}

func substitute(lin linear, index int, expr linear) linear {
	r := linear{}

	a := lin.a[index]
	lin.a[index] = 0

	for i := 0; i < N; i++ {
		r.a[i] = lin.a[i] + a*expr.a[i]
	}
	r.b = lin.b + a*expr.b
	return r
}

func evalRecursive(vars []variable, free []int, index int) (int, bool) {
	if index == len(free) {
		vals := [N]int{}
		total := 0

		for i := len(vars) - 1; i >= 0; i-- {
			x := eval(vars[i], vals)
			if x < -EPS || math.Abs(x-math.Round(x)) > EPS {
				return 0, false
			}
			vals[i] = int(math.Round(x))
			total += vals[i]
		}

		return total, true
	}

	best, found := math.MaxInt, false
	for x := 0; x <= vars[free[index]].max; x++ {
		vars[free[index]].val = x
		total, ok := evalRecursive(vars, free, index+1)

		if ok {
			found = true
			best = min(best, total)
		}
	}

	if found {
		return best, true
	} else {
		return 0, false
	}
}

func eval(v variable, vals [N]int) float64 {
	if v.free {
		return float64(v.val)
	}

	x := v.expr.b
	for i := 0; i < N; i++ {
		x += v.expr.a[i] * float64(vals[i])
	}
	return x
}

func ParseMachine(input string) Machine {
	m := Machine{}

	start := strings.Index(input, "[")
	end := strings.Index(input, "]")

	endStateStr := input[start+1 : end]
	m.bitCount = len(endStateStr)

	for i, r := range endStateStr {
		if r == '#' {
			m.endMask |= 1 << uint(i)
		}
	}

	rest := input[end+1:]
	leftCurly := strings.Index(rest, "{")
	movesStr := strings.TrimSpace(rest[:leftCurly])

	for _, part := range strings.Fields(movesStr) {
		if !strings.HasPrefix(part, "(") || !strings.HasSuffix(part, ")") {
			continue
		}
		inner := part[1 : len(part)-1]
		if len(inner) == 0 {
			continue
		}

		mask := 0
		for _, ns := range strings.Split(inner, ",") {
			i, _ := strconv.Atoi(ns)
			mask |= 1 << uint(i)
		}
		m.moveMasks = append(m.moveMasks, mask)
	}

	start = strings.Index(input, "{")
	end = strings.Index(input, "}")

	joltStr := input[start+1 : end]
	if strings.TrimSpace(joltStr) != "" {
		for _, ns := range strings.Split(joltStr, ",") {
			n, _ := strconv.Atoi(strings.TrimSpace(ns))
			m.joltages = append(m.joltages, n)
		}
	}

	return m
}
