package main

import (
	"fmt"
)

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}

	sum := 0
	all_nums := make(map[int]int)
	for _, num := range nums {
		sum += num
		if _, ok := all_nums[num]; ok {
			all_nums[num] += 1
		} else {
			all_nums[num] = 1
		}
	}

	if sum%k != 0 {
		return false
	}
	target := sum / k
	fmt.Println(sum, all_nums, target)
	return true
}

func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	fmt.Println("1111", canPartitionKSubsets(nums, k))
}
