package main

import (
    "time"
    "sync"
    // "fmt"
)

type AA struct {
    mu *sync.Mutex
    aa map[string]map[string]int
}

func (a *AA)set(b int) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.aa["aaa"]["aaa"] = b
}

func (a *AA)get() map[string]int {
    a.mu.Lock()
    defer a.mu.Unlock()
    return a.aa["aaa"]    
}

var sm = &sync.Map{} 

func main() {
    sm.Store("aaa",map[string]string{"aaa":"aaa"})
    // a := &AA{aa:map[string]map[string]int{"aaa":map[string]int{"aaa":1}},mu:new(sync.Mutex)} 
    go func() {
        if err1 := recover(); err1 != nil {
            return
        }
        for {
            // a.set(123)
            v,_ := sm.Load("aaa")
            v1 := v.(map[string]string)
            // fmt.Println(v,v1,ok)
            v1["aaa"] = "bbb"
        }
    }()
    go func() {
        if err1 := recover(); err1 != nil {
            return
        }
        for {
            // v := a.get()
            // _ = v["aaa"]

            // v,ok := sm.Load("x")
            // sm.Store("x","bbbb")
            // fmt.Println(v,ok)

            v,_ := sm.Load("aaa")
            v1 := v.(map[string]string)
            // _ = v1[""]
            _ = v1["aaa"]
            // fmt.Sprintln("111",v1)
            // fmt.Println(v,v1,ok)
        }
    }()

    time.Sleep(10 * time.Second)
}