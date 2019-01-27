package main
 
import (
    "fmt"
    "time"
)
 
func main() {
    nums := []int{3522, 181, 521, 515, 304, 123, 2512, 312, 922, 407, 146, 1932, 4037, 2646, 3871, 269}
    // new_nums := make([]int, 0, 0)

    nums2 := []int{3522, 181, 521, 515, 304, 123, 2512, 312, 922, 407, 146, 1932, 4037, 2646, 3871, 269}

    fmt.Println(len(nums),cap(nums),append(nums[:12], []int{1, 2, 3}...), nums,len(nums),cap(nums))

    fmt.Println(len(nums2),cap(nums2),append(nums2[:14], []int{1, 2, 3}...), nums2,len(nums2),cap(nums2))
    t := time.Now()
    fmt.Println(t.Hour())
    switch  {
    case t.Hour() > 12:
        fmt.Println("Morning was passed.")
        fallthrough
    case t.Hour() > 23:
        fmt.Println("Afternoon was passed.")
        fallthrough
    default:
        fmt.Println("Now too early.")
 
    }
}
