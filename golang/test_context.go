package main

import (
    "fmt"
    "context"
    "time"
)

func main() {
    parentctx := context.Background()
    ctx111, cancel := context.WithCancel(parentctx)
    ctx := context.WithValue(ctx111,"gest","test")

    go func(ctx context.Context) {
        ctx1, cancel1 := context.WithCancel(ctx)
        defer func() {
            fmt.Println("cccles")
            cancel1() 
        }()
        for {
            select {
            case <-ctx1.Done():
                fmt.Println("----ctx1--------")
                return
            default:
                fmt.Println(ctx1)
            //     cancel1() 
            }
        }
        
    }(ctx)

    fmt.Println("----sleep1222--------")
    time.Sleep(1*time.Second)

    cancel()
    fmt.Println("----sleep11--------")
    time.Sleep(2*time.Second)

    fmt.Println("----end--------")
}