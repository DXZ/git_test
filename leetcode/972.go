package main

import (
    "fmt"
    "strconv"
    "strings"
    "math"
)

func transport(s string) float64 {
    start := strings.Index(s,"(")
    if start < 0 {
        result,_ := strconv.ParseFloat(s,64)
        return result
    }
    tmp := s[start+1:len(s)-1]
    result_str := s[:start]
    for i:=0;i<15;i++ {
        result_str += tmp
    }
    fmt.Println(result_str)

    result,_ := strconv.ParseFloat(result_str,64)
    return result
}

func isRationalEqual(S string, T string) bool {
    return math.Abs(transport(S) - transport(T)) < 0.00000000000001 
}

func main() {
   fmt.Println(isRationalEqual2("0.99(9)","1.00"))
}

func isRationalEqual2(S string, T string) bool {
    ns, ds := parse(S)
    nt, dt := parse(T)
    fmt.Println(ns, ds,nt, dt)
    return ns*dt == ds*nt
}

func parse(s string) (n, d int64) {
    var n1, d1, n2, d2 int64
    for _, c := range s {
        switch {
        case c == '.':
            d1 = 1
        case c == '(':
            d2 = 1
        case c == ')':
        case d2 > 0:
            n2 = n2*10 + (int64(c) - '0')
            d2 = d2 * 10
        case d1 > 0:
            n1 = n1*10 + (int64(c) - '0')
            d1 = d1 * 10
        default:
            n1 = n1*10 + (int64(c) - '0')
        }
        fmt.Println(s,"-----",n1, d1, n2, d2)
    }
    
    if d1 < 1 {
        d1 = 1
    }
    if d2 < 10 {
        n, d = n1, d1
    } else {
        n, d = n1*(d2-1)+n2, d1*(d2-1)
    }
    fmt.Println(s,d1,d2,n1,n2,n,d)
    return n, d
}