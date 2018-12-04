// Package main is for the core binary to play the game
package main

import "fmt"

func main() {
	input := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	twosCount := 0
	threesCount := 0

	for _, str := range input {
		// create map
		charMap := make(map[rune]int)
		for _, r := range str {
			_, exists := charMap[r]
			if exists {
				charMap[r]++
			} else {
				charMap[r] = 1
			}
		}

		//find twos and threes
		plusTwo := false
		plusThree := false
		for key := range charMap {
			if charMap[key] == 2 {
				plusTwo = true
			} else if charMap[key] == 3 {
				plusThree = true
			}
		}

		if plusTwo {
			twosCount++
		}
		if plusThree {
			threesCount++
		}
	}

	fmt.Printf("twosCount = %v\n", twosCount)
	fmt.Printf("threesCount = %v\n", threesCount)
	fmt.Printf("checksum = %v\n", twosCount*threesCount)
}
