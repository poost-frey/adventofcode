// Package main is for the core binary to play the game
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

	for _, c := range claims {
		fmt.Printf("%v\n", c)
	}
}

type claim struct {
	ID     int
	X      int
	Y      int
	Width  int
	Height int
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
