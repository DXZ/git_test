package main

import (
    "fmt"
    "time"
    "sync"
    // "encoding/json"
    "runtime"
)

type AA struct {
    mu *sync.Mutex
    aa map[string]map[string]int
}


func (a *AA)ts(b int) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.aa["aaa"]["aaa"] = b
    fmt.Println(a,b)
}

func (a *AA)get() map[string]int {
    a.mu.Lock()
    defer a.mu.Unlock()
    return a.aa["aaa"]    
}

func test1(a *map[string]int,b int) {
    (*a)["aaa"] = b
    fmt.Println(a,b)
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    a :=&AA{aa:map[string]map[string]int{"aaa":map[string]int{"aaa":1}},mu:new(sync.Mutex)} 
    fmt.Println(a)
    // go test1(a, "1111")
    // go test1(a, "2222")
    // var beginWg sync.WaitGroup
    f := func(a *AA,b int) {
        // if err1 := recover(); err1 != nil {
        //     return
        // }
        // beginWg.Wait()
        // a.ts(b)
        // fmt.Println("------------------")
        c := a.get()
        c["aaa"] = b
    }

    N := 10000000000000
    // for i := 0; i < N; i++ {
    //     beginWg.Add(1)
    // }

    for i := 0; i < N; i++ {
        go f(a, i)
        // beginWg.Done()
    }

    fmt.Println(a,"---------end------")

    // m := map[string]string{"\xff": "a"}

    // go func() {
    //     for i := 0; i < 10000; i++ {
    //         m["x"] = "b"
    //     }
    // }()

    // for i := 0; i < 10000; i++ {
    //     if _, err := json.Marshal(m); err != nil {
    //         panic(err)
    //     }
    // }
    time.Sleep(20*time.Second)
    fmt.Println(a)
    // f := func(a *map[string]int,b int) {
    //     beginWg.Wait()
    //     test1(a, b)
    // }

    // N := 100
    // for i := 0; i < N; i++ {
    //     beginWg.Add(1)
    // }

    // for i := 0; i < N; i++ {
    //     go f(&a, i)
    //     beginWg.Done()
    // }

    // fmt.Println(a)
    // time.Sleep(2*time.Second)
    // fmt.Println(a)
}