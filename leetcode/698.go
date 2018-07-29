package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%k != 0 {
		return false
	}
	expected := sum / k

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	// fmt.Println(sum, expected, nums)

	return canPartitionKSubsetsItem(nums, k, 0, 0, expected)
}

func canPartitionKSubsetsItem(nums []int, expected_count int, sum int, i int, expected int) bool {
	fmt.Println(nums, expected_count, sum, i, expected)
	if expected_count == 0 && len(nums) == 0 {
		return true
	}

	if len(nums) < 1 {
		return false
	}

	if i >= len(nums) {
		return false
	}

	new_nums := make([]int, 0, 0)
	// new_nums
	new_nums = append(new_nums, nums[:i]...)
	new_nums = append(new_nums, nums[i+1:]...)
	new_sum := sum + nums[i]
	if new_sum == expected {
		return canPartitionKSubsetsItem(new_nums, expected_count-1, 0, 0, expected)
	}
	if new_sum < expected {
		return canPartitionKSubsetsItem(new_nums, expected_count, new_sum, i, expected) || canPartitionKSubsetsItem(nums, expected_count, sum, i+1, expected)
	}

	return canPartitionKSubsetsItem(nums, expected_count, sum, i+1, expected)
}

func main() {
	nums := []int{3522, 181, 521, 515, 304, 123, 2512, 312, 922, 407, 146, 1932, 4037, 2646, 3871, 269}
	k := 5
	fmt.Println("1111", canPartitionKSubsets(nums, k))
}
