package main

import (
	"fmt"
)

func main() {
	nums := []int{3522, 181, 521, 515, 304, 123, 2512, 312, 922, 407, 146, 1932, 4037, 2646, 3871, 269}
	new_nums := make([]int, 0, 0)

	fmt.Println(append(nums[:1], nums[2:]...), nums)
}
