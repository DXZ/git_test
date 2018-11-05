package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	max := prices[len(prices)-1]
	result := -1
	for i := len(prices) - 1; i >= 0; i-- {

		if max < prices[i] {
			max = prices[i]
		}

		if result < max-prices[i] {
			result = max - prices[i]
		}
	}

	if result < 0 {
		return 0
	}
	return result
}

func main() {
	fmt.Println(maxProfit([]int{1, 2}))
}
