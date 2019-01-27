package main

import (
    "reflect"
    "fmt"
)

func main() {
    var a = "12324"
    fmt.Println(reflect.ValueOf(a).Kind())
    u:= User{"张三",20}
    t:=reflect.TypeOf(u)
    fmt.Println(t)

    v := reflect.ValueOf(u)
    fmt.Println(v)

    fmt.Println(t.Kind())
    

    mPrint:=v.MethodByName("Print")
    args:=[]reflect.Value{reflect.ValueOf("前缀12321")}
    fmt.Println(mPrint.Call(args))

    m := Struct2Map(u)
    fmt.Println(m,reflect.ValueOf(m).Kind())
}



func Struct2Map(obj interface{}) map[string]interface{} {
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)
    var data = make(map[string]interface{})

    for i:=0;i<t.NumField();i++{
        data[t.Field(i).Name] = v.Field(i).Interface()
    }
    return data
}


func (u User) Print(prfix string){
    fmt.Printf("%s:Name is %s,Age is %d",prfix,u.Name,u.Age)
}

type User struct {
    Name string
    Age int
}
