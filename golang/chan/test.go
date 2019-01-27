package main

import (
	"fmt"
	"time"
)


type user struct {
	name string
	sex  string
	age int
	h  [2]int
}


func main() {
	// go spinner(100 * time.Millisecond)
	// const n = 45
	// fibN := fib(n) // slow
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	var a1 [3]string
	a2 := [3]string{"a","b","c"}
	a1 = a2
	a2[0]="aaa"
	a3:=&a2
	fmt.Println(a1,a2,a3)



	var array1 [2][2]int
	var array2 [2][2]int
	array2[0][0] = 10 
	array2[0][1] = 20 
	array2[1][0] = 30 
	array2[1][1] = 40

	array1 = array2
	array1[0][0]  =1000
	fmt.Println(array2,array1)


	var slice1 []int
	slice2 := []int{}

	fmt.Println(slice2,slice1,len(slice1),len(slice2),cap(slice1),cap(slice2))


	var u user
	var u2 user

	u.name = "u"
	u.h[1] = 12
	u2 = u
	u2.age = 1
	u2.h[0] = 10000
	u2.name = "u1"
	fmt.Println("user:",u,"--",u2)

	t1:= time.Now()
	time.Sleep(1*time.Second)
	t2:= time.Now()
	fmt.Println(t2.Sub(t1).Nanoseconds()/1000000)
}

// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range `-\|/` {
// 			fmt.Printf("\r%c", r)
// 			time.Sleep(delay)
// 		}
// 	}
// }

// func fib(x int) int {
// 	if x < 2 {
// 		return x
// 	}
// 	return fib(x-1) + fib(x-2)
// }
