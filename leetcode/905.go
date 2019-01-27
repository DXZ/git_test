package main

import (
	"fmt"
)

func sortArrayByParity(A []int) []int {
	result := make([]int, len(A), len(A))
	j := 0
	k := len(A) - 1
	for i := 0; i < len(A); i++ {

		if A[i]%2 == 0 {
			result[j] = A[i]
			j++
		} else {
			result[k] = A[i]
			k--
		}

	}
	return result
}

// func sortArrayByParity2(A []int) []int {
// 	k := 0
// 	length := len(A) - 1
// 	for i := 0; i < len(A)/2 && k < length/2; {
// 		if A[i]%2 == 0 {
// 			i++
// 		} else {
// 			A[i], A[length-k] = A[length-k], A[i]
// 			k++
// 		}
// 	}
// 	return A
// }

func main() {
	// fmt.Println(sortArrayByParity2([]int{4, 1, 3, 2}))
	// fmt.Println(sortArrayByParity2([]int{3, 1, 2, 4}))
}
