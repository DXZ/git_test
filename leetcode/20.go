package main


import (
    "fmt"
)

type StackNode struct {
    Data interface{}
    next *StackNode
}

type Stack struct {
    top   *StackNode
    count int
}


func (s *Stack)Push(data interface{}) {
    node := new(StackNode)
    node.Data = data
    if s.top == nil {
        node.next = nil
        s.top = node
    }else{
        node.next = s.top
        s.top = node        
    }
    s.count ++
}


var Map = map[string]int{
    "(":-1,
    ")":1,
    "[":-2,
    "]":2,
    "{":-3,
    "}":3,
}

func (s *Stack)Pop() (data interface{}){
    if s.top == nil {
        return nil
    }
    data = s.top.Data
    s.top = s.top.next
    return
}

func isValid(s string) bool {
    words := ([]rune)(s)
    stack := new(Stack)
    stack.count = 0
    // stack.top = new(StackNode)
    // stack.top.Data = 

    for _,w := range words {
        word := string(w)
        // fmt.Println(word)
        if v,ok:=Map[word];ok{
            if v < 0 {
                stack.Push(word)
            }else{
                data := stack.Pop()
                if data == nil {
                    return false
                }
                if Map[data.(string)] + v != 0 {
                    return false
                }
            }
        }else{
            return false
        }
    }
    if stack.top == nil {
        return true
    }
    return false
}


func isValid2(s string) bool {
    stack := []rune{}
    for _, b := range s{
        if len(stack) != 0 {
            out:= false
            switch b {
                case ']':
                if stack[len(stack) - 1] != '['{
                        return false
                } 
                out = true
                case ')':
                if stack[len(stack) - 1] != '('{
                        return false
                }   
                out = true
                case '}':
                if stack[len(stack) - 1] != '{'{
                        return false
                }
                out = true
                default:
                    stack = append(stack, b)
            }
            if out {
                stack = stack[0:len(stack)-1]
            }
        } else {
            stack = append(stack, b)
        }
    }
    return len(stack) == 0
}

func main() {
    fmt.Println(isValid("[]"))

    a := []int{1,2,3,4,5,6}
    fmt.Println(a[0:len(a)-1])
}