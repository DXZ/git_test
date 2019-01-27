package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	sums := nums[0]
	target := 0
	for i := 0; i < len(nums); i++ {
		if sums < nums[i] {
			sums = nums[i]
		}
		if nums[i] < 0 && target == 0 {
			continue
		} else if nums[i] > 0 {
			target += nums[i]
		}
		if target > sums {
			sums = target
		}
		if nums[i] < 0 && target > 0 {
			if target+nums[i] > 0 {

				target += nums[i]

			} else {
				target = 0
			}
		}
	}
	return sums
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-1, -2}))
}
