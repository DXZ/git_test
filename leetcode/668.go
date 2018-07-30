package main

import (
	"fmt"
	"sort"
)

func findKthNumber(m int, n int, k int) int {
	x := m
	if m > n {
		x = n
	}

	if k*2 > x*x {
		k = x*x - k
	}
	end = x / 2

	for {
		mid := (end) / 2
		if mid*(mid+1)/2 > k {
			end = mid
		}
	}
	return 0
}

func main() {
	fmt.Println("1111", findKthNumber(3, 3, 5))
}
