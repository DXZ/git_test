package main

import (
	"fmt"
	// "math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	int_list := []int{}
	for x > 0 {
		int_list = append(int_list, x%10)
		x /= 10

	}
	fmt.Println(int_list)

	for i := 0; i <= len(int_list)/2; i++ {
		// fmt.Println(int_list[i] != int_list[len(int_list)-i-1], len(int_list)/2, i)
		if int_list[i] != int_list[len(int_list)-i-1] {
			fmt.Println("1111")
			return false
		}
	}

	return true
}

func isPalindrome2(x int) bool {
	if x >= 0 && x <= 9 {
		return true
	}

	if x < 0 || x%10 == 0 {
		return false
	}

	result := 0
	x1 := x
	for ; x1 > 0; x1 /= 10 {
		result *= 10
		result += x1 % 10
	}
	fmt.Println(x1, result)
	return result+x1 == x
}

func main() {
	fmt.Println(isPalindrome2(1010))
	// fmt.Println(math.Log10(100))
}
