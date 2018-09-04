package main

type Node struct {
	v        int
	child    []*Node
	depth    int
	MaxDepth int
	parent   *Node
}

func (node *Node) getNow() *Node {
	tmp := node
	for i := 0; i < node.depth; i++ {
		for _, n := range tmp.child {
			// if n.depth != tmp.depth {
			// 	break
			// }
		}
	}
	return tmp
}

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	all := make([]*Node, 0, 0)

	start := new(Node)

	now := start

	depth := 0

	for i = 0; i < len(nums); i++ {
		if start.depth != nil && start.depth > 0 {
			temp := new(Node)
			temp.v = nums[i]

			if temp.v <= start.v {
				// all = append(all, start)
				// start := temp
				// depth = 1
				// start.depth = 1

			} else if temp.v <= now.v {
				if now.parent != nil && tmp.v > now.parent.v {
					now.parent.child = append(now.parent.child, temp)
				}
			} else {

			}
		} else {
			depth = 1
			start.v = i
			start.depth = depth
		}
	}

	return 1
}

func main() {
	nums := []int{3, 2, 4}
	fmt.Println(findNumberOfLIS(nums))
}
