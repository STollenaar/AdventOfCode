package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

type Module struct {
	kind         string
	destinations []string
	inputs       map[string]bool
	active       bool
}

type Pulse struct {
	origin      string
	strength    string
	destination string
}

var (
	modules = make(map[string]*Module)
	pulses  []*Pulse

	processedPulses = map[string]int{"high": 0, "low": 0}
)

func main() {
	lines := internal.Reader()

	for _, line := range lines {
		line = strings.ReplaceAll(line, " -> ", " ")
		line = strings.ReplaceAll(line, ", ", " ")
		att := strings.Split(line, " ")
		kind := strings.Split(att[0], "")[0]
		if kind != "%" && kind != "&" {
			kind = ""
		}
		var name string
		if kind == "" {
			name = att[0]
		} else {
			name = strings.Join(strings.Split(att[0], "")[1:], "")
		}
		modules[name] = &Module{kind: kind, destinations: att[1:], inputs: make(map[string]bool)}
	}
	for k, m := range modules {
		if m.kind == "&" {
			for kT, mT := range modules {
				if slices.Contains(mT.destinations, k) {
					m.inputs[kT] = false
				}
			}
		}
	}

	var buttonPresses int
	for {
		if buttonPresses == 1000 {
			fmt.Printf("Solution for part 1: %d\n", processedPulses["high"]*processedPulses["low"])
		}
		// fmt.Printf("Doing press %d/%d\n", i, 1000)
		pulses = []*Pulse{{strength: "low", destination: "broadcaster", origin: "button"}}
		processedPulses["low"]++

		for len(pulses) > 0 {
			var processed []*Pulse
			for _, pulse := range pulses {
				if pulse.destination == "rx" && pulse.strength == "low" {
					fmt.Printf("Solution for part 2: %d\n", buttonPresses)
					break
				}
				pulseProcess := modules[pulse.destination].processPulse(pulse)
				processed = append(processed, pulseProcess...)
			}
			pulses = processed
		}
		buttonPresses++
	}
}

func (m *Module) processPulse(in *Pulse) (out []*Pulse) {
	if m == nil {
		return nil
	}

	switch m.kind {
	case "%":
		if in.strength == "high" {
			return nil
		}
		m.active = !m.active
		if m.active {
			in.strength = "high"
		} else {
			in.strength = "low"
		}
	case "&":
		if in.strength == "low" {
			m.inputs[in.origin] = false
		} else {
			m.inputs[in.origin] = true
		}
		var values []bool
		for _, input := range m.inputs {
			values = append(values, input)
		}
		if allHighValue(values) {
			in.strength = "low"
		} else {
			in.strength = "high"
		}
	}
	for _, des := range m.destinations {
		processedPulses[in.strength]++
		// fmt.Printf("%s -%s-> %s\n", in.destination, in.strength, des)
		out = append(out, &Pulse{origin: in.destination, destination: des, strength: in.strength})
	}
	return
}

func allHighValue(a []bool) bool {
	for _, v := range a {
		if !v {
			return false
		}
	}
	return true
}
