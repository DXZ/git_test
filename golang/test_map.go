package main

import (
    "time"
    "fmt"
    "sync"
)

var m = make(map[string]string) //wrong
var sm = &sync.Map{} // right

func main() {
    sm.Store("x","aaa")
    m["x"] = "aaa"
    sm.Store("x", m)
    go func() {
        if err1 := recover(); err1 != nil {
            return
        }
        for {
            m["x"] = "xxxx"
            // sm.Store("x","bbbb")
        }
    }()
    go func() {
        if err1 := recover(); err1 != nil {
            return
        }
        for {
            _ = m["x"]
            // v,ok := sm.Load("x")
            // sm.Store("x","bbbb")
            // fmt.Println(v,ok)
        }
    }()
    fmt.Println("----")
    time.Sleep(1 * time.Second)
}