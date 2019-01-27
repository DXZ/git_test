package main

import (
    "fmt"
)

func main() {
    fmt.Println("111")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    result := make([]int,0,0)
    index := 0
    return flipMatchVoyageSelf(root,voyage,&index,result)
}


func flipMatchVoyageSelf(root *TreeNode,voyage []int,index *int,result []int) []int{
    if root == nil {
        return result
    }

    if root.Val != voyage[*index] {
        return []int{-1}
    }

    *index = *index + 1
    if root.Left != nil && root.Left.Val != voyage[*index] {
        root.Left,root.Right = root.Right,root.Left
        result = append(result,root.Val)
    }

    result = flipMatchVoyageSelf(root.Left,voyage,index,result)
    return flipMatchVoyageSelf(root.Right,voyage,index,result)
}



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

func (s *Stack)Pop() (data interface{}){
    if s.top == nil {
        return nil
    }
    data = s.top.Data
    s.top = s.top.next
    s.count-- 
    return
}


func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    result := make([]int,0,0)
    index := 0
    stack := new(Stack)
    stack.count = 0
    stack.Push(root)

    for stack.count>0 {
        curin := stack.Pop()
        fmt.Println(curin)
        if curin == nil {
            continue
        }

        cur := curin.(*TreeNode)
        if cur == nil {
            continue
        }
        if cur.Val != voyage[index] {
            return []int{-1}
        }
        index += 1
        if cur.Left != nil && cur.Left.Val != voyage[index] {
            cur.Left,cur.Right = cur.Right,cur.Left
            result = append(result,cur.Val)
        }
        stack.Push(cur.Right)
        stack.Push(cur.Left)
    }

    return result
}