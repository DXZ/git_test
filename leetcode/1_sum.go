package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		m[num] = index
	}
	for index, num := range nums {
		if v, ok := m[target-num]; ok && v != index {
			return []int{v, index}
		}
	}
	return []int{-1, -1}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for index, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, index}
		} else {
			m[num] = index
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
}

func twoSum(nums []int, target int) []int {
	t := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, ok := t[v]; ok {
			return []int{vv, i}
		}
		t[target-v] = i
	}
	return nil
}
