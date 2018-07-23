package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	var flag int
	if x > 0 {
		if x >= 2<<30 {
			return 0
		}
		flag = 1
	} else if x < 0 {
		if x <= -2<<30 {
			return 0
		}
		flag = -1
		x = -x
	} else {
		return x
	}

	s := strconv.Itoa(x)

	result := 0

	for i := 0; i < len(s); i++ {
		result_tmep := (x % 10)
		fmt.Println("1=", result_tmep)
		for j := 0; j < (len(s) - 1 - i); j++ {
			result_tmep *= 10
		}
		fmt.Println("2", result_tmep)
		result += result_tmep
		x = x / 10
	}

	if result >= 2<<30 {
		return 0
	}
	return flag * result
}

func reversegood(x int) int {
	result := 0
	for x != 0 {
		tmp := result*10 + x%10
		if tmp > math.MaxInt32 || tmp < math.MinInt32 {
			return 0
		}

		result = tmp
		x /= 10
	}

	return result
}

func main() {
	fmt.Println(2 << 30)
	fmt.Println(reverse(-1002))
}
