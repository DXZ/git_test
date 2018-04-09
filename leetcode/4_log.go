package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	length := (len1 + len2) / 2
	index1 := 0
	index2 := 0
	for i := 0; i < length; i++ {
		if nums1[index1] > nums2[index2] {
			index1 += 1

		} else {
			index2 += 1

		}
	}
	return 0.0
}

func main() {
	fmt.Println("aaaa")
}
