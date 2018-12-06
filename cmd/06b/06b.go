package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func main() {
	// expected answer = 17
	input := []string{
		"1, 1",
		"1, 6",
		"8, 3",
		"3, 4",
		"5, 5",
		"8, 9",
	}

	coordinates := parseInput(input)
	grid := initGrid(coordinates)

	for y, row := range grid {
		for x := range row {
			s := sumBelow10K(coordinates, x, y)
			if s {
				grid[y][x] = 1
			} else {
				grid[y][x] = 0
			}
		}
	}

	count := getCount(grid)

	//print grid
	for _, row := range grid {
		for _, value := range row {
			fmt.Printf("%v", value)
		}
		fmt.Println()
	}

	fmt.Printf("max count = %v\n", count)
}

func sumBelow10K(coordinates []pos, currentX int, currentY int) bool {
	sum := 0
	for _, p := range coordinates {
		distance := int(math.Abs(float64(p.X-currentX))) + int(math.Abs(float64(p.Y-currentY)))
		sum += distance
	}

	return sum < 32
}

func getCount(grid [][]int) (count int) {
	for _, row := range grid {
		for _, value := range row {
			if value == 1 {
				count++
			}
		}
	}
	return
}

func initGrid(c []pos) [][]int {
	// find max
	max := 0
	for _, p := range c {
		if p.X > max {
			max = p.X
		}
		if p.Y > max {
			max = p.Y
		}
	}
	coefficient := 2
	output := make([][]int, max*coefficient)
	for ind := range output {
		output[ind] = make([]int, max*coefficient)
	}
	return output
}

type pos struct {
	X int
	Y int
}

func parseInput(str []string) (output []pos) {
	regex, _ := regexp.Compile("(\\d+), (\\d+)")
	for _, s := range str {
		matches := regex.FindAllStringSubmatch(s, 5)
		output = append(output, pos{X: atoi(matches[0][1]), Y: atoi(matches[0][2])})
	}
	return
}

func atoi(str string) (i int) {
	i, _ = strconv.Atoi(str)
	return
}
