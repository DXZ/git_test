package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}

func findK(nums1 []int, nums2 []int, target int) int {
	if len(nums1) > len(nums2) {
		return findK(nums2, nums1, target)
	}

	if len(nums1) == 0 {
		return nums2[target]
	}

	// if len(nums1)
}

// find (m+n)/2 min number
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	const (
		INT_MAX = int(^uint(0) >> 1)
		INT_MIN = ^INT_MAX
	)
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	n := 2 * len1
	L1, L2, R1, R2, c1, c2, lo := 0, 0, 0, 0, 0, 0, 0

	for lo <= n {
		c1 = (lo + n) / 2
		c2 = len1 + len2 - c1

		if c1 == 0 {
			L1 = INT_MIN
		} else {
			L1 = nums1[(c1-1)/2]
		}
		if c1 == 2*len1 {
			R1 = INT_MAX
		} else {
			R1 = nums1[c1/2]
		}

		if c2 == 0 {
			L2 = INT_MIN
		} else {
			L2 = nums2[(c2-1)/2]
		}
		if c2 == 2*len2 {
			R2 = INT_MAX
		} else {
			R2 = nums2[c2/2]
		}

		if L1 > R2 {
			n = c1 - 1
		} else if L2 > R1 {
			lo = c1 + 1
		} else {
			break
		}

	}
	return (math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2.0
}

func fasterAnswer(nums1 []int, nums2 []int) float64 {
	var total = len(nums1) + len(nums2)
	var mid = total / 2
	var median, previous, i, j, idx int

	for {
		var nums1Reached = i == len(nums1)
		var nums2Reached = j == len(nums2)

		if nums1Reached {
			median = nums2[j]
			j++
		} else if nums2Reached {
			median = nums1[i]
			i++
		} else {
			if nums1[i] > nums2[j] {
				median = nums2[j]
				j++
			} else {
				median = nums1[i]
				i++
			}
		}

		if idx == mid {
			break
		}
		idx++
		previous = median
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(median+previous) / 2.0
	} else {
		return float64(median)
	}
}

// func finfWidth(nums1 []int, nums2 []int, target int) int {
// 	if len(nums2) == 0 {
// 		return nums1[target]
// 	}

// 	if len(nums1) == 0 {
// 		return nums2[target]
// 	}

// 	if target == 0 {
// 		if nums1[0] < nums2[0] {
// 			return nums1[0]
// 		}
// 		return nums2[0]
// 	}

// 	new_target :=
// }

func main() {
	nums1 := []int{1, 2, 4, 10}
	nums2 := []int{3, 5, 7}
	fmt.Println("aaaa", findMedianSortedArrays(nums1, nums2))
}
