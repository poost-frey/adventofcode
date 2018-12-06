package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "dabAcCaCBAcCcaDA"
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortestOutput string

	for _, letter := range alphabet {
		temp := strings.Replace(input, string(letter), "", -1)
		temp = strings.Replace(temp, strings.ToLower(string(letter)), "", -1)
		output := react(temp)
		if shortestOutput == "" || len(output) < len(shortestOutput) {
			shortestOutput = output
		}
	}

	fmt.Printf("final string length: %v", len(shortestOutput))
}

func react(input string) string {
	output := input
	reacted := true
	for reacted {
		firstChar, secondChar := findMatch(output)

		if firstChar == "" {
			reacted = false
		} else {
			output = strings.Replace(output, firstChar+secondChar, "", -1)
		}
	}
	return output
}

func findMatch(str string) (string, string) {
	for ind, r := range str {
		if ind < len(str)-1 {
			char := string(r)
			upper := strings.ToUpper(char)
			lower := strings.ToLower(char)
			nextChar := string(str[ind+1])
			if char != nextChar && (upper == nextChar || lower == nextChar) {
				return char, nextChar
			}
		}
	}

	return "", ""
}
