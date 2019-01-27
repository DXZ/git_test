package main

import (
    "time"
    "fmt"
)

func main() {
    start:= time.Now()

    time.Sleep(1*time.Second)
    end := time.Now()
    latency := end.Sub(start)
    fmt.Println("%s",int64(latency/time.Millisecond))
}