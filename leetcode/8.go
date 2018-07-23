package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	s := []rune(str)
	result := 0
	temp := 1
	flag := false
	for i := 0; i < len(s); i++ {
		if s[i] == 32 && !flag {
			continue
		}

		fmt.Println(i, result, s[i], s[i] == 45, int(s[i]) == 45, !flag && (s[i] == 45 || s[i] == 43))
		if !flag && (s[i] == 45 || s[i] == 43) {
			temp = 44 - int(s[i])
		} else if s[i] >= 48 && s[i] <= 57 {
			result = result*10 + (int(s[i]) - 48)
		} else {
			break
		}

		flag = true
		if temp < 0 && result > 2<<30 {
			result = math.MinInt32
			break
		} else if temp > 0 && result > 2<<30-1 {
			result = 2<<30 - 1
			break
		}
	}

	return result * temp
}

func main() {
	fmt.Println(myAtoi("0-1"))
	fmt.Println([]rune(" +-0123456789"))
}
