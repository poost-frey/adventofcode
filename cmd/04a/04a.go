// Package main is for the core binary to play the game
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	input := []string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up",
	}

	records := buildRecords(input)

	for _, rec := range records {
		fmt.Printf("%v\n", rec)
	}
}

type record struct {
	ID         int
	Month      int
	Day        int
	SleepStart int
	WakeStart  int
}

type sleep struct {
	Month  int
	Day    int
	Minute int
}

type wake struct {
	Month  int
	Day    int
	Minute int
}

func buildRecords(input []string) []record {
	sleeps := getSleeps(input)
	wakes := getWakes(input)
	return findGuards(input, sleeps, wakes)
}

func findGuards(input []string, sleeps []sleep, wakes []wake) (records []record) {
	regex, _ := regexp.Compile("-(\\d{2})-(\\d{2}) \\d{2}:\\d{2}\\] Guard #(\\d+)")

	for _, str := range input {
		matches := regex.FindAllStringSubmatch(str, 10)

		if len(matches) > 0 {
			month := atoi(matches[0][1])
			day := atoi(matches[0][2])

			sleepMin := findSleepMin(sleeps, month, day)
			wakeMin := findWakeMin(wakes, month, day)

			records = append(records, record{
				ID:         atoi(matches[0][3]),
				Month:      month,
				Day:        day,
				SleepStart: sleepMin,
				WakeStart:  wakeMin,
			})
		}
	}
	return
}

func getSleeps(input []string) (sleeps []sleep) {
	regex, _ := regexp.Compile("-(\\d{2})-(\\d{2}) \\d{2}:(\\d{2})\\] f")

	for _, str := range input {
		matches := regex.FindAllStringSubmatch(str, 10)

		if len(matches) > 0 {
			sleeps = append(sleeps, sleep{
				Month:  atoi(matches[0][1]),
				Day:    atoi(matches[0][2]),
				Minute: atoi(matches[0][3]),
			})
		}
	}
	return
}

func getWakes(input []string) (wakes []wake) {
	regex, _ := regexp.Compile("-(\\d{2})-(\\d{2}) \\d{2}:(\\d{2})\\] w")

	for _, str := range input {
		matches := regex.FindAllStringSubmatch(str, 10)

		if len(matches) > 0 {
			wakes = append(wakes, wake{
				Month:  atoi(matches[0][1]),
				Day:    atoi(matches[0][2]),
				Minute: atoi(matches[0][3]),
			})
		}
	}
	return
}

func findSleepMin(sleeps []sleep, month int, day int) int {
	for _, sleep := range sleeps {
		if sleep.Month == month && sleep.Day == day {
			return sleep.Minute
		}
	}
	return -1
}

func findWakeMin(wakes []wake, month int, day int) int {
	for _, wake := range wakes {
		if wake.Month == month && wake.Day == day {
			return wake.Minute
		}
	}
	return -1
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
