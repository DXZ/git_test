package main

import (
    "fmt"
)

func main() {
    fmt.Println(powerfulIntegers(9,80,100))
}

func powerfulIntegers2(x int, y int, bound int) []int {
    m := make(map[int]int)
    curX := 1
    for curX + 1 <= bound {
        curY := 1
        for curX + curY <= bound {
            m[curX + curY] = 1
            if y == 1 {
                break
            } else {
                curY = y * curY
            }
        }
        if x == 1 {
            break
        } else {
            curX = x * curX
        }
    }
    var ret []int
    for k, _ := range m {
        ret = append(ret, k)
    }
    return ret
}


func powerfulIntegers(x int, y int, bound int) []int {
    result := make([]int,0,0)
    if bound < 1 {
        return result
    }

    Map := make(map[int]bool)
    x_tmp := 1
    for i:=0;i<bound;i++ {
        y_tmp := 1
        for j:=0;j<bound;j++{
            sum := y_tmp+x_tmp
            if sum <= bound {
                if _,ok:=Map[sum];!ok {
                    result = append(result,sum)
                }
                Map[sum] = true    
            }

            if y_tmp == y_tmp * y || x_tmp+y_tmp*y>bound {
                break
            }
            y_tmp = y_tmp * y
        }

        if x_tmp == x_tmp * x || x_tmp*x + 1 > bound {
            break
        }
        x_tmp = x_tmp * x
    }

    return result
}