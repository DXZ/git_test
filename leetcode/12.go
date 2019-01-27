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

func intToRoman(num int) string {
	result := ""
	if num/1000 > 0 {

	}
	return result
}

func main() {
	fmt.Println(intToRoman(12))
}
