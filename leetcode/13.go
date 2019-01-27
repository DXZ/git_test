package main

import (
	"fmt"
)

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// V X L C D M
var Map = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if len(s) < 1 {
		return 0
	}
	last_v := Map[string(s[0])]
	result := 0
	for i := 1; i < len(s); i++ {
		v := Map[string(s[i])]
		// if v ==
		if last_v == 0 {
			last_v = v
			continue
		}
		if v > last_v {
			result += (v - last_v)
			last_v = 0
		} else {
			result += last_v
			last_v = v
		}
	}
	if last_v > 0 {
		result += last_v
	}

	return result
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
