package main

import (
    "fmt"
    "os"
    "time"
    "syscall"
    "os/signal"
)

func main() {
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        sig := <-sigs
        fmt.Println("----------------",sig)
        time.Sleep(100*time.Second)
        done <- true
    }()

    fmt.Println("awaiting signal",len(sigs))
    <-done
    fmt.Println("exiting")
}