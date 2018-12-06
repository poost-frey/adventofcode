package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	input := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}

	claims := buildClaims(input)
	cloth := prepareMatrix(1000)

	// map out claims on cloth
	for _, c := range claims {
		for i := 0; i < c.Width; i++ {
			for j := 0; j < c.Height; j++ {
				x := c.X + i
				y := c.Y + j
				if cloth[x][y] == 0 {
					cloth[x][y] = 1
				} else {
					cloth[x][y]++
				}
			}
		}
	}

	// count overlaps
	overlap := 0
	for _, row := range cloth {
		for _, sqin := range row {
			if sqin > 1 {
				overlap++
			}
		}
	}

	fmt.Printf("square inches of overlap: %v\n", overlap)
}

type claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
}

func prepareMatrix(side int) [][]int {
	matrix := make([][]int, side)

	for indX, row := range matrix {
		matrix[indX] = make([]int, side)
		for indY := range row {
			matrix[indX][indY] = 0
		}
	}

	return matrix
}

func buildClaims(input []string) (claims []claim) {
	for _, i := range input {
		regex, err := regexp.Compile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
		if err != nil {
			log.Fatalf("failed to parse input string %s", i)
		}
		matches := regex.FindAllStringSubmatch(i, 10)
		claims = append(claims, claim{
			ID:     atoi(matches[0][1]),
			X:      atoi(matches[0][2]),
			Y:      atoi(matches[0][3]),
			Width:  atoi(matches[0][4]),
			Height: atoi(matches[0][5]),
		})
	}
	return
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, sqin := range row {
			fmt.Printf("%v", sqin)
		}
		fmt.Println()
	}
}
