package main

import (
    "fmt"
)

func addToArrayForm2(A []int, K int) []int {
    // a := 0
    var a uint64 = 0
    for _,i:=range A {
        a = a*10 + uint64(i)
    }

    
    r := a + uint64(K)
    // fmt.Println(a,r)
    result := make([]int,0,0)
    for ;r>=10; {
        tmp := int(r%10)
        result = append([]int{tmp},result[:]...)
        r = r/10
    }

    result = append([]int{int(r)},result[:]...)
    // fmt.Println(result)

    return result
}


func addToArrayForm(A []int, K int) []int {
    // a := 0
    result := make([]int,0,0)

    k := make([]int,0,0)
    for ;;{
        if K < 10 {
            k = append([]int{K},k[:]...)
            break
        }
        tmp := K%10
        k = append([]int{tmp},k[:]...)
        K = K/10
    }
    fmt.Println(k)

    flag := 0

    i := len(A)-1
    j := len(k)-1
    // i:=0
    // j:=0
    for ;; {
        if i >=0 && j>=0{
            tmp := A[i]+k[j]+flag
            flag = tmp/10
            tmp = tmp%10
            result = append([]int{tmp},result[:]...)
            i--
            j--
        }else if i>=0 {
            tmp := A[i]+flag
            flag = tmp/10
            tmp = tmp%10
            result = append([]int{tmp},result[:]...)
            i--
        }else if j>=0 {
            tmp := k[j]+flag
            flag = tmp/10
            tmp = tmp%10
            result = append([]int{tmp},result[:]...)
            j--
        }else {
            break
        }
    }

    if flag > 0 {
        result = append([]int{flag},result[:]...)
    }

    return result
}


func main() {
    fmt.Println(addToArrayForm([]int{9,9,9},1))
}