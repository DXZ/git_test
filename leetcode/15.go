package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	Map = make(map[int]bool)
	for _, v := range nums {
		Map[v] = true
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; i < len(nums); j++ {
			if _, ok := Map[nums[i]+nums[j]]; ok {

			}
		}
	}
}

func main() {
	fmt.Println(threeSum([]int{1, 2, -2, -1, 0}))
}
