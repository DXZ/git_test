package main

import (
	"fmt"
	"strconv"
)

func shortestBridge(A [][]int) int {
	temp := make(map[string]bool)
	var singlelang [][]int
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] != 0 {
				// strconv.Itoa
				key := strconv.Itoa(i) + "_" + strconv.Itoa(j)
				// if i < len(A)-1 {
				// 	_, ok := temp[strconv.Itoa(i+1)+"_"+strconv.Itoa(j)]
				// 	if ok {
				// 		temp[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = true
				// 	}
				// }
			}
		}
	}

}

func main() {
	fmt.Println(shortestBridge([][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}))
}
