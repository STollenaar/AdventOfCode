package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/STollenaar/AdventOfCode/internal"
)

var (
	workflows          = make(map[string][]Workflow)
	rejected, accepted []Part

	brackets  = regexp.MustCompile("[{}]")
	operants  = regexp.MustCompile(`\w*`)
	operation = regexp.MustCompile(`\W`)
)

type Workflow struct {
	feature     string
	operation   string
	destination string
	value       int
}

type Part struct {
	attribute map[string]int
}

func init() {

}

func main() {
	lines := internal.Reader()

	var doneWorkflows bool
	for _, line := range lines {
		line = brackets.ReplaceAllString(line, " ")
		if line == "" {
			doneWorkflows = true
			continue
		}
		parts := strings.Split(line, " ")
		rules := strings.Split(parts[1], ",")
		part := Part{attribute: make(map[string]int)}
		for _, rule := range rules {
			if doneWorkflows {
				att := strings.Split(rule, "=")
				value, _ := strconv.Atoi(att[1])
				part.attribute[att[0]] = value
			} else {
				if !strings.Contains(rule, ":") {
					workflows[parts[0]] = append(workflows[parts[0]], Workflow{destination: rule})
					continue
				}
				rule = strings.ReplaceAll(rule, ":", " ")
				comp := operants.FindAllStringSubmatch(strings.Split(rule, " ")[0], -1)
				op := operation.FindAllStringSubmatch(strings.Split(rule, " ")[0], -1)
				dest := strings.Split(rule, " ")[1]

				value, _ := strconv.Atoi(comp[1][0])
				workflows[parts[0]] = append(workflows[parts[0]], Workflow{feature: comp[0][0], operation: op[0][0], destination: dest, value: value})
			}
		}
		if doneWorkflows {
			currentWorkflow := "in"
			for currentWorkflow != "" && currentWorkflow != "R" && currentWorkflow != "A" {
				for _, rule := range workflows[currentWorkflow] {
					if rule.operation == "<" {
						if part.attribute[rule.feature] < rule.value {
							currentWorkflow = rule.destination
							break
						}
					} else if rule.operation == ">" {
						if part.attribute[rule.feature] > rule.value {
							currentWorkflow = rule.destination
							break
						}
					} else {
						currentWorkflow = rule.destination
						break
					}
				}
			}
			if currentWorkflow == "A" {
				accepted = append(accepted, part)
			} else {
				rejected = append(rejected, part)
			}
		}
	}
	var total int
	for _, acPart := range accepted {
		for _, v := range acPart.attribute {
			total += v
		}
	}
	fmt.Printf("Solution for Part 1: %d\n", total)
}
