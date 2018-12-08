package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// expected answer = CABDFE
	input := []string{
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
	}

	rules := parseInput(input)

	remaining := getPossibleSteps(rules)
	length := len(remaining)

	var sequence string

	for i := 0; i < length; i++ {
		nextStep := getNextStep(rules, remaining)
		remaining = strings.Replace(remaining, nextStep, "", 1)
		sequence += nextStep
	}

	fmt.Printf("steps: %s", sequence)
}

type rule struct {
	before string
	after  string
}

func parseInput(str []string) (output []rule) {
	regex, _ := regexp.Compile("Step ([A-Z]) must be finished before step ([A-Z]) can begin.")
	for _, s := range str {
		matches := regex.FindAllStringSubmatch(s, 5)
		output = append(output, rule{before: matches[0][1], after: matches[0][2]})
	}
	return
}

func getPossibleSteps(rules []rule) (remaining string) {
	stepMap := make(map[string]int)

	// prepare remaining
	for _, r := range rules {
		stepMap[r.before]++
		stepMap[r.after]++
	}

	for key := range stepMap {
		remaining += key
	}
	return
}

func getNextStep(rules []rule, remaining string) string {
	var possibilities []string

	for _, r := range remaining {
		usable := true
		for _, rule := range rules {
			if rule.after == string(r) && strings.Contains(remaining, rule.before) {
				usable = false
			}
		}
		if usable {
			possibilities = append(possibilities, string(r))
		}
	}

	sort.Strings(possibilities)
	return possibilities[0]
}
