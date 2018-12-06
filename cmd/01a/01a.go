package main

import "fmt"

func main() {
	pos := 0
	seq := []int{+1, -1, 0}

	for _, i := range seq {
		pos += i
	}

	fmt.Printf("pos = %v", pos)
}
