package main

import (
    "fmt"
    "encoding/json"
)


var U map[string]string = map[string]string {
    "json":"json",
    "aaaa":"aaaa",
}

func main() {
    fmt.Println("start")
    b, err := json.Marshal(U)
    fmt.Println(string(b),err)
}