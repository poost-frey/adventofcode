package main

import (
	"fmt"
	"regexp"
	"sort"
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

	history := make(map[int][]int)
	for _, rec := range records {
		// init guard in history
		_, exists := history[rec.ID]
		if !exists {
			history[rec.ID] = make([]int, 60)
		}

		// mark sleeping times in schedule
		for min, asleep := range rec.Asleep {
			if asleep {
				history[rec.ID][min]++
			}
		}
	}

	// find sleepiest guard and minute
	sleepiestGuard := parsedHistory{}
	for id, schedule := range history {
		currentGuard := parsedHistory{ID: id}
		for min, sleepCount := range schedule {
			currentGuard.sleepCount += sleepCount
			if sleepCount > currentGuard.sleepiestMinHits {
				currentGuard.sleepiestMinHits = sleepCount
				currentGuard.sleepiestMin = min
			}
		}

		if currentGuard.sleepiestMinHits > sleepiestGuard.sleepiestMinHits {
			sleepiestGuard = currentGuard
		}
	}

	fmt.Printf(
		"sleepiest guard #%v slept for %v min\n",
		sleepiestGuard.ID,
		sleepiestGuard.sleepCount,
	)
	fmt.Printf(
		"with the sleepiest min being %v that occured %v times\n",
		sleepiestGuard.sleepiestMin,
		sleepiestGuard.sleepiestMinHits,
	)
	fmt.Printf("id x min = %v", sleepiestGuard.ID*sleepiestGuard.sleepiestMin)
}

type parsedHistory struct {
	ID               int
	sleepCount       int
	sleepiestMin     int
	sleepiestMinHits int
}

type record struct {
	ID     int
	Month  int
	Day    int
	Asleep []bool
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
	regex, _ := regexp.Compile("-(\\d{2})-(\\d{2}) (\\d{2}):\\d{2}\\] Guard #(\\d+)")

	for _, str := range input {
		matches := regex.FindAllStringSubmatch(str, 10)

		if len(matches) > 0 {
			month := atoi(matches[0][1])
			day := atoi(matches[0][2])

			if atoi(matches[0][3]) == 23 {
				day++
			}

			sleepMins := findSleepsByDay(sleeps, month, day)
			wakeMins := findWakesByDay(wakes, month, day)

			schedule := make([]bool, 60)

			for ind, s := range sleepMins {
				for i := s; i < wakeMins[ind]; i++ {
					schedule[i] = true
				}
			}

			records = append(records, record{
				ID:     atoi(matches[0][4]),
				Month:  month,
				Day:    day,
				Asleep: schedule,
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

func findSleepsByDay(sleeps []sleep, month int, day int) (output []int) {
	for _, sleep := range sleeps {
		if sleep.Month == month && sleep.Day == day {
			output = append(output, sleep.Minute)
		}
	}
	sort.Ints(output)
	return output
}

func findWakesByDay(wakes []wake, month int, day int) (output []int) {
	for _, wake := range wakes {
		if wake.Month == month && wake.Day == day {
			output = append(output, wake.Minute)
		}
	}
	sort.Ints(output)
	return output
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
