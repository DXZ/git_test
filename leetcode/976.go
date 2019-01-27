package main


import (
    "fmt"
    "sort"
)

func largestPerimeter(A []int) int {
    sort.Ints(A)
    result := 0
    // fmt.Println(A)
    for i:=len(A)-1;i>=2;i-- {
        if A[i] < A[i-1] + A[i-2] {
            result = A[i-1] + A[i-2] + A[i]
            break
        }
    }
    return result
}

func main() {
    fmt.Println(largestPerimeter([]int{9,6,1,2,3,4}))
}