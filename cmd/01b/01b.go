package main

import "fmt"

func main() {
	pos := 0
	seq := []int{}

	pos = cycle(pos, seq)

	fmt.Printf("pos = %v", pos)
}

func cycle(pos int, seq []int) int {
	var set []int
	for {
		for _, i := range seq {
			pos += i
			if contains(set, pos) {
				return pos
			}
			set = append(set, pos)
		}
	}

}

func contains(slice []int, item int) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
