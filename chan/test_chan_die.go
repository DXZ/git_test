package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go father(i)
	}
	time.Sleep(1 * time.Minute)
}

func father(f int) {
	fmt.Println("father----start--", f)
	defer func() {
		fmt.Println("father----end--", f)
	}()
	for i := 0; i < 20; i++ {
		go test(f, i)
	}

	time.Sleep(5 * time.Second)
}

func test(i int, j int) {
	for {
		fmt.Println("--------start-------", i, j)
		time.Sleep(1 * time.Second)
		fmt.Println("--------end-------", i, j)
	}
}
