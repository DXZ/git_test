package main

import (
	"fmt"
)

// find (m+n)/2 min number
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	m := (len1 + len2) / 2

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
	nums1 := []int{1, 2, 4}
	nums2 := []int{3, 5, 7}
	fmt.Println("aaaa", findMedianSortedArrays(nums1, nums2))
}
