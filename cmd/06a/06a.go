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

	for x, row := range grid {
		for y, _ := range row {
			closestInd := determineClosest(coordinates, x, y)
			grid[x][y] = closestInd
		}
	}

	count := getCount(grid, len(coordinates))

	//print grid
	for _, row := range grid {
		for _, value := range row {
			fmt.Printf("%v", value)
		}
		fmt.Println()
	}

	fmt.Printf("max count = %v\n", count)
}

func determineClosest(coordinates []pos, currentX int, currentY int) int {
	closestInd := -1
	closestDistance := -1
	for ind, p := range coordinates {
		distance := int(math.Abs(float64(p.X-currentX))) + int(math.Abs(float64(p.Y-currentY)))
		if distance == closestDistance {
			return -1
		}
		if closestDistance == -1 || distance < closestDistance {
			closestDistance = distance
			closestInd = ind
		}
	}
	return closestInd
}

func getCount(grid [][]int, numberOfCoordinates int) (max int) {
	count := make([]int, numberOfCoordinates)
	for _, row := range grid {
		for _, value := range row {
			if value >= 0 {
				count[value]++
			}
		}
	}

	for _, c := range count {
		if c > max {
			max = c
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

type bounds struct {
	minX int
	maxX int
	minY int
	maxY int
}

func getBounds(coordinates []pos) bounds {

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
