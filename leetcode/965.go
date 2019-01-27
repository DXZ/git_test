package main

import (
    "fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    return Serach(root,root.Val)
}

func Serach(root *TreeNode,val int) bool{
    if root == nil {
        return true
    }

    if root.Val != val {
        return false
    }

    if Serach(root.Left,val) {
        return false
    }

    if Serach(root.Right,val) {
        return false
    }
}

func main() {
    
}