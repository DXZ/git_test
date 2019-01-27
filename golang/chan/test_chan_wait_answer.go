package main

import (
	"errors"
	"fmt"
	// "ginfoxer/tools/gpool"
	"time"
)

var (
	ENDERR = errors.New("END")
	Good   = errors.New("Good")
)

func main() {
	fmt.Println("---main start---")
	// a:=chan
	// var a chan error
	a := make(chan error, 10)
	length := 20
	countchan := make(chan int, length)
	// fgpool := gpool.New(10)
	for i := 0; i < length; i++ {
		// fgpool.Add(1)
		go func(bbb int) {
			defer func() {
				countchan <- 1
				if err := recover(); err != nil {
					a <- err.(error)
					// close(a)
				}
			}()
			fmt.Println("go-----", bbb)
			if bbb == 15 {
				panic(ENDERR)
			}
			time.Sleep(3 * time.Second)
			// fgpool.Done()
		}(i)
	}
	// fgpool.Wait()
	// a <- Good

	count := 0
	for {
		if count == length {
			break
		}
		select {
		case i := <-a:
			fmt.Println("select", i)
			panic(i)
			break
		case i := <-countchan:
			if i == 1 {
				count += 1
			}
			continue
		}
	}

	close(a)

	// fmt.Println(<-a)
	fmt.Println("nbbbbaaaaaaaa")

	fmt.Println("aaaaaaaa")
	// for i := range a {
	// 	fmt.Println("error:", i)
	// }

	fmt.Println("end")
}
