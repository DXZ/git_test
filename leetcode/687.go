package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivalue(root *TreeNode, max int) (maxEnd int, deep int) {
	deep = 0
	maxEnd = max
	if root == nil {
		return
	} else if root.Left == nil && root.Right == nil {
		return
	} else if root.Left == nil {
		maxEnd, deep = longestUnivalue(root.Right, max)
		if root.Right.Val == root.Val {
			deep += 1
		} else {
			deep = 0
		}
	} else if root.Right == nil {
		maxEnd, deep = longestUnivalue(root.Left, max)
		if root.Left.Val == root.Val {
			deep += 1
		} else {
			deep = 0
		}
	} else {
		maxEndl, deepl := longestUnivalue(root.Left, max)
		maxEndr, deepr := longestUnivalue(root.Right, max)

		if maxEndl > maxEnd {
			maxEnd = maxEndl
		}
		if maxEndr > maxEnd {
			maxEnd = maxEndr
		}
		// if deepl > deepr {

		// }
		if root.Left.Val == root.Val {
			deepl += 1
		} else {
			deepl = 0
		}

		if root.Right.Val == root.Val {
			deepr += 1
		} else {
			deepr = 0
		}

		if root.Right.Val == root.Val && root.Left.Val == root.Right.Val {
			if maxEnd < deepl+deepr {
				maxEnd = deepl + deepr
			}
		}

		if deepr > deepl {
			deep = deepr
		} else {
			deep = deepl
		}
	}

	if deep > maxEnd {
		maxEnd = deep
	}

	if max > maxEnd {
		maxEnd = max
	}

	return
}

func longestUnivaluePath(root *TreeNode) int {
	max, _ := longestUnivalue(root, 0)
	return max
}

func longestUnivaluePath2(root *TreeNode) int {
	maxL := 0

	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 0
	} else if root.Left == nil {
		if root.Right.Val == root.Val {
			rightL := longestUnivaluePath(root.Right) + 1
			if maxL < rightL {
				maxL = rightL
			}
		} else {
			rightL := longestUnivaluePath(root.Right)
			if maxL < rightL {
				maxL = rightL
			}
		}
	} else if root.Right == nil {
		leftL := longestUnivaluePath(root.Left)
		if root.Left.Val == root.Val {
			if maxL < leftL+1 {
				maxL = leftL + 1
			}

		} else {
			if maxL < leftL {
				maxL = leftL
			}
		}
	} else {
		leftL := longestUnivaluePath(root.Left)
		rightL := longestUnivaluePath(root.Right)
		fmt.Println("1111111:", leftL, rightL)
		if root.Val == root.Left.Val {
			leftL += 1
		}
		if root.Val == root.Right.Val {
			rightL += 1
		}
		if root.Right.Val == root.Left.Val && root.Right.Val == root.Val {
			leftL = rightL + leftL
			rightL = leftL
		}
		fmt.Println("222222:", leftL, rightL)
		if leftL > maxL {
			maxL = leftL
		}
		if rightL > maxL {
			maxL = rightL
		}
	}

	return maxL
}

func main() {
	fmt.Println("1111")
}
