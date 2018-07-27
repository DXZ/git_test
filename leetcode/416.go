package main

import (
	"fmt"
	"sort"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func test(nums []int) bool {
	sum := 0
	for _, e := range nums {
		sum += e
	}
	if sum%2 == 1 {
		return false
	} else {
		sum = sum / 2
		n := len(nums)
		dp := make([][]int, n, n)

		for i := 0; i < n; i++ {
			dp[i] = make([]int, sum+1, sum+1)
		}

		for i := nums[0]; i <= sum; i++ {
			dp[0][i] = nums[0]
		}
		// fmt.Println("dp333=", dp)
		for i := 1; i < n; i++ {
			for j := nums[i]; j <= sum; j++ {
				dp[i][j] = Max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
			fmt.Println(dp)
		}

		fmt.Println(dp)
		if dp[n-1][sum] == sum {
			return true
		}
		return false
	}

}

func canPartitiongood(nums []int) bool {
	// get sum of the array
	sum := 0
	for _, e := range nums {
		sum += e
	}
	if sum%2 == 1 {
		return false
	}
	expected := sum / 2
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	// sort.Sort(sort.IntSlice(nums))
	fmt.Println(nums)
	return canPartitionWithExpectation(nums, 0, 0, expected)
}

func canPartitionWithExpectation(nums []int, curr_index int, sum int, expected int) bool {
	if len(nums) <= curr_index {
		return false
	}
	new_sum := sum + nums[curr_index]
	if new_sum == expected {
		return true
	}
	if new_sum > expected {
		return false
	}
	return canPartitionWithExpectation(nums, curr_index+1, new_sum, expected) || canPartitionWithExpectation(nums, curr_index+1, sum, expected)
}

func canPartition(nums []int) bool {
	sum := 0
	// all_nums := make(map[int]int)
	max := 0

	for _, num := range nums {
		sum += num
		// if _, ok := all_nums[num]; ok {
		// 	all_nums[num] += 1
		// } else {
		// 	all_nums[num] = 1
		// }
		if max < num {
			max = num
		}
	}

	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	log := make(map[int]bool)

	if target < max {
		return false
	}
	fmt.Println(target)
	return get(nums, log, target)
}

func get(nums []int, log map[int]bool, target int) bool {
	fmt.Println(len(nums), target, nums, log)

	if _, ok := log[target]; ok {
		return false
	}
	if target < 0 {
		return false
	}
	if target == 0 {
		return true
	}
	if len(nums) < 1 {
		return false
	} else if len(nums) == 1 {
		if nums[0] == target {
			return true
		}
		return false
	} else {

		// if _, ok := log[nums[0]]; ok {
		// 	return get(nums[1:], log, target)
		// }

		if get(nums[1:], log, target-nums[0]) {
			return true
		}

		log[target-nums[0]] = true
		return get(nums[1:], log, target)
	}

}

func main() {
	// nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 97, 95}
	// nums := []int{8, 1, 2, 3, 4, 5, 6, 7}
	nums := []int{2, 1, 3, 4}
	// nums := []int{1, 1, 5, 1, 1, 1}
	fmt.Println(canPartitiongood(nums))
	// fmt.Println(canPartition(nums))

	fmt.Println(test(nums))
}
