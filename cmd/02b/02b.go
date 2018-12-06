package main

import "fmt"

func main() {
	input := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	firstInd, secondInd, pos := findMatch(input)

	output := slice(input[firstInd], pos)

	fmt.Printf("first code = %s\n", input[firstInd])
	fmt.Printf("second code = %s\n", input[secondInd])
	fmt.Printf("output = %s\n", output)
}

func findMatch(input []string) (int, int, int) {
	for indI, i := range input {
		for indJ, j := range input {
			if indI == indJ {
				continue
			}
			difference := -1
			for pos := 0; pos < len(i); pos++ {
				if i[pos] != j[pos] {
					if difference > 0 {
						difference = -1
						break
					}
					difference = pos
				}
			}

			if difference >= 0 {
				return indI, indJ, difference
			}
		}
	}

	return 0, 0, 0
}

func slice(str string, pos int) (output string) {
	for ind, r := range str {
		if ind != pos {
			output += string(r)
		}
	}
	return
}
