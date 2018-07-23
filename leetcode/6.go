package main

import (
	"fmt"
)

func convert2(s string, numRows int) string {
	// Skip := 2*numRows - 2
	var skip int
	var result []rune
	sss := []rune(s)
	for i := 0; i < numRows; i++ {
		skip = 2 * (numRows - i - 1)
		if numRows-1 == i {
			skip = 2 * (numRows - 1)
		}
		// fmt.Println("i=", i, skip, sss[i])
		result = append(result, sss[i])

		for j := i + skip; j < len(sss); j += skip {
			result = append(result, sss[j])
			// fmt.Println("j=", j, sss[j])
		}
	}

	return string(result)
}

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	var skip int
	var result []rune
	sss := []rune(s)
	for i := 0; i < numRows; i++ {
		skip = 2 * (numRows - i - 1)
		if numRows-1 == i {
			skip = 2 * (numRows - 1)
		}
		fmt.Println("i=", i, skip, string(sss[i]))
		result = append(result, sss[i])

		for j := i + skip; j < len(sss); {
			result = append(result, sss[j])
			fmt.Println("j=", j, i, skip, string(sss[j]))
			if numRows-1 == i || i == 0 {
				skip = 2 * (numRows - 1)
			} else {
				skip = 2*(numRows-1) - skip
			}

			j += skip
		}
	}

	return string(result)
}

func main() {
	fmt.Println(convert("AB", 1))
}
