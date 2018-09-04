package main

import (
	"fmt"
)

func maxArea(height []int) int {
	max := 0
	tmp := 0
	for i, j := 0, len(height)-1; i < j; {
		// fmt.Println(i, j, j-i, height[j])
		if height[i] < height[j] {
			tmp = (j - i) * height[i]
			i++
		} else {
			tmp = (j - i) * height[j]
			j--
		}

		if tmp > max {
			max = tmp
		}
	}

	return max

}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(maxArea(nums))
}
