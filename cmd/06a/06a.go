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
			closestInd := determineClosest(coordinates, x, y)
			grid[y][x] = closestInd
		}
	}

	count := getCount(grid, coordinates)

	//print grid
	alphabet := "abcdefgh"
	for _, row := range grid {
		for _, value := range row {
			var char string
			if value < 0 {
				char = "."
			} else {
				char = string(alphabet[value])
			}
			fmt.Printf("%s", char)
		}
		fmt.Println()
	}

	fmt.Printf("max count = %v\n", count)
}

func determineClosest(coordinates []pos, currentX int, currentY int) int {
	closestInd := -1
	closestDistance := -1
	duplicate := false
	for ind, p := range coordinates {
		distance := int(math.Abs(float64(p.X-currentX))) + int(math.Abs(float64(p.Y-currentY)))
		if distance == closestDistance {
			duplicate = true
		}
		if closestDistance == -1 || distance < closestDistance {
			closestDistance = distance
			closestInd = ind
			duplicate = false
		}
	}

	if duplicate {
		return -1
	}
	return closestInd
}

func getCount(grid [][]int, coordinates []pos) (max int) {
	excluded := getExcludedCoordinateInds(grid)
	count := make([]int, len(coordinates))
	for _, row := range grid {
		for _, value := range row {
			if value >= 0 && !isExcluded(excluded, value) {
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

func isExcluded(excluded []int, item int) bool {
	for _, e := range excluded {
		if e == item {
			return true
		}
	}
	return false
}

func getExcludedCoordinateInds(grid [][]int) (excluded []int) {
	m := make(map[int]bool)
	max := len(grid) - 1
	for i := 0; i <= max; i++ {
		m[grid[0][i]] = true
		m[grid[max][i]] = true
		m[grid[i][0]] = true
		m[grid[i][max]] = true
	}

	for key := range m {
		excluded = append(excluded, key)
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
