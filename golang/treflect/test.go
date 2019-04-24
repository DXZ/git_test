package main

import (
    "fmt"
    "reflect"
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
    Grade       string `json:"grade" binding:"required"`
}

type UserProfile2 struct {
    Subject     int `json:"subject" binding:"required"`
    Refresh     int `json:"refresh" binding:"required"`
    Active      int `json:"active" binding:"required"`
    Buy         int `json:"buy" binding:"required"`
    exten       map[string]interface{}
}


func transport(u1 interface{}) {
    fmt.Println("---transport--",reflect.TypeOf(u1),reflect.ValueOf(u1),reflect.TypeOf(u1).Kind())

    // a := reflect.TypeOf(u1).Elem()
    // fmt.Println(reflect.ValueOf(a).FieldByName("Subject"))

}

func GetProfiletype(v int) uint64{
    if v == -1 {
        return ^uint64(0)
    }
    return uint64(1 << uint(v+1))
}


func main() {
    fmt.Println(GetProfiletype(-1)&8>0,GetProfiletype(2)&4>0)
    fmt.Println("-----------------")
    var u UserProfile = UserProfile{1,2,3,4,5,6,7,"8"}

    fmt.Println("--main111---",reflect.TypeOf(u),reflect.ValueOf(u))
    u1 := &UserProfile{1,2,3,4,5,6,7,"88888"}

    transport(u1)
    u2 := new(UserProfile2)
    u2.Subject = u1.Subject
    u2.Refresh = u1.Refresh
    u2.Active = u1.Active
    // u2.exten[]
    u2.exten = make(map[string]interface{})
    fmt.Println(u1,u2)

    colums := []string{"Grade","English"}

    for _,i := range colums {
        v := reflect.ValueOf(*u1).FieldByName(i)
        u2.exten[i] = v
        fmt.Println(i,v)

        value,bool := reflect.TypeOf(*u1).FieldByName(i)
        fmt.Println("------2222---",value.Tag.Get("json"),bool)
    }
    
    // u2.exten[] = 





    // p := &Person{}

    // var t = reflect.TypeOf(u1)
    // v := reflect.ValueOf(u1).Elem()

    

    // v.FieldByName("Name").Set(reflect.ValueOf(name))
    // v.FieldByName("Age").Set(reflect.ValueOf(age))

    fmt.Println(u2)

    fmt.Println("end")

}