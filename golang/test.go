package main

import (
    "time"
    "fmt"
    "reflect"
    "encoding/json"
)


var SourceMapping map[string]int = map[string]int {
    "ALL":                  0,
    "iPhoneNewParent":      1,
}

var UserColums map[int][]string = map[int][]string {
    1:[]string{"aaa","bbb"},
    2:[]string{"aaa","bbb"},
}

type UserProfiles []UserProfile

type UserProfile struct {
    Subject     int `json:"subject" binding:"required"`
    Refresh     int `json:"refresh" binding:"required"`
    Active      int `json:"active" binding:"required"`
    Buy         int `json:"buy" binding:"required"`
    Math        int `json:"math" binding:"required"`
    English     int `json:"english" binding:"required"`
    Chinese     int `json:"chinese" binding:"required"`
    Grade       int `json:"grade" binding:"required"`
}

type UserProfile2 struct {
    Subject     int `json:"subject" binding:"required"`
    Refresh     int `json:"refresh" binding:"required"`
    Active      int `json:"active" binding:"required"`
    Buy         int `json:"buy" binding:"required"`
}



func CheckProfile(u1 interface{},u2 interface{}) bool {
    // v1 := reflect.ValueOf(u1)
    // v2 := reflect.ValueOf(u2)
    // // vt := reflect.TypeOf(u1)
    // // fmt.Println(vt)
    // fmt.Println(v1,)
    // for i := 0; i < v1.Elem().NumField(); i++ {
    //     field := v1.Field(i)  //返回结构体的第i个字段
    //     field2 := v2.Field(i) //返回结构体的第i个字段

        
    //     fmt.Println(field,field2)
    // }
    return true
}


func main() {
    fmt.Println(UserColums)
    start:= time.Now()

    aa := 1
    bb := 10
    aa,bb = bb,aa
    fmt.Println(bb,aa)
    time.Sleep(1*time.Second)
    end := time.Now()
    latency := end.Sub(start)
    fmt.Println("%s",int64(latency/time.Millisecond))
    fmt.Println(time.Now().Format("20060102"))

    fmt.Println(SourceMapping["ALL"])
    fmt.Println(SourceMapping["ALL1"])

    fmt.Println("----------------------")
    var u1 = &UserProfile{1,2,3,4,5,6,7,8}
    var u2 = &UserProfile{10,20,30,40,50,60,7,8}
    fmt.Println(CheckProfile(u1,u2))
    // uu := new(UserProfiles)
    
    uu := make([]*UserProfile,0,0)
    uu = append(uu,u1)
    uu = append(uu,u2)

    b, err := json.Marshal(uu)
    fmt.Println(string(b),err)


    flag := 3
    switch flag{
    case 1:fmt.Println("111")
    case 2:fmt.Println("111")
    default:break
    }

    fmt.Println("end")

}