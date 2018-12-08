package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const numWorkers = 2 // 2 for example. 5 for actual
const baseTime = 0   // 0 for example. 60 for actual
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	// expected answer = 15
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

	var sequence string
	var active string
	workers := make([]*worker, numWorkers)

	// initial setup
	for i := 0; i < numWorkers; i++ {
		workers[i] = &worker{ID: i}
		nextStep := getNextStep(rules, remaining, active)
		if nextStep != "" {
			remaining = strings.Replace(remaining, nextStep, "", 1)
			workers[i].setJob(nextStep)
			active += nextStep
		}
	}

	var totalTime int
	printHeaders(workers)
	printState(totalTime, workers, sequence)

	for len(remaining) > 0 || len(active) > 0 {
		totalTime++
		for _, w := range workers {
			working := w.isBusy()
			if working {
				done := w.tickSec()
				if done != "" {
					active = strings.Replace(active, done, "", 1)
					sequence += done
					nextStep := getNextStep(rules, remaining, active)
					if nextStep != "" {
						remaining = strings.Replace(remaining, nextStep, "", 1)
						w.setJob(nextStep)
						active += nextStep
					}
				}
			} else {
				nextStep := getNextStep(rules, remaining, active)
				if nextStep != "" {
					remaining = strings.Replace(remaining, nextStep, "", 1)
					w.setJob(nextStep)
					active += nextStep
				}
			}
		}
		printState(totalTime, workers, sequence)
	}

	fmt.Printf("steps: %s\n", sequence)
	fmt.Printf("totalTime: %v", totalTime)
}

func printHeaders(workers []*worker) {
	spacer := "   "
	fmt.Print("Second" + spacer)
	for _, w := range workers {
		fmt.Printf("Worker %v"+spacer, w.ID)
	}
	fmt.Println("Done")
}

func printState(totalTime int, workers []*worker, sequence string) {
	spacer3 := "   "
	spacer7 := "       "
	sec := strconv.Itoa(totalTime)
	spaceDigits := 5 - len(sec)
	for i := 0; i < spaceDigits; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%v  "+spacer3, sec)

	for _, w := range workers {
		var printChar string
		if w.activeStep == "" {
			printChar = "."
		} else {
			printChar = w.activeStep
		}
		fmt.Printf(spacer3+"%s"+spacer7, printChar)
	}

	fmt.Println(sequence)
}

type worker struct {
	ID            int
	activeStep    string
	timeRemaining int
}

func (w *worker) setJob(step string) {
	time := getStepLength(step)
	w.activeStep = step
	w.timeRemaining = time
}

func (w worker) isBusy() bool {
	return w.activeStep != ""
}

func (w *worker) tickSec() string {
	w.timeRemaining--
	if w.timeRemaining == 0 {
		step := w.activeStep
		w.activeStep = ""
		return step
	}
	return ""
}

func getStepLength(step string) int {
	ind := strings.Index(alphabet, step)
	return ind + 1 + baseTime
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

func getNextStep(rules []rule, remaining string, active string) string {
	var possibilities []string

	for _, r := range remaining {
		usable := true
		for _, rule := range rules {
			if rule.after == string(r) && (strings.Contains(remaining, rule.before) || strings.Contains(active, rule.before)) {
				usable = false
			}
		}
		if usable {
			possibilities = append(possibilities, string(r))
		}
	}

	sort.Strings(possibilities)
	if len(possibilities) > 0 {
		return possibilities[0]
	}
	return ""
}
