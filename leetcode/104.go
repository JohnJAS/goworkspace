package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) addLeft(val int) {
	t.Left = &TreeNode{Val: val}
}

func (t *TreeNode) addRight(val int) {
	t.Right = &TreeNode{Val: val}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	root := &TreeNode{Val: 5}
	root.addLeft(4)
	root.Left.addLeft(11)
	root.Left.Left.addLeft(7)
	root.Left.Left.addRight(2)

	fmt.Println(maxDepth(root))
}
