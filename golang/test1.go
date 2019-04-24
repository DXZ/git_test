package main

import (
    "strings"
    "strconv"
    "fmt"
)

const INT_MAX = int(^uint(0) >> 1)

func AppVersionCount(version string) int{
    if version == "-1" {
        return INT_MAX
    }

    vl := strings.Split(version, ".")
    if len(vl) <= 0 {
        return -1
    }
    sum := 0
    for index,item := range vl {
        itemnum, err := strconv.Atoi(item)
        if err != nil {
            return -1
        }
        sum += itemnum << uint(8*(4-index))
    }

    return sum
}
func main() {
    x := 2
    y := 0
    fmt.Println(x<<1)
    fmt.Println(y>>1,INT_MAX)

    fmt.Println(AppVersionCount("1.100.2"))
}
