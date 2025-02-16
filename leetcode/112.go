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

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}
		return false
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	root := &TreeNode{Val: 5}
	root.addLeft(4)
	root.Left.addLeft(11)
	root.Left.Left.addLeft(7)
	root.Left.Left.addRight(2)
	root.addRight(8)
	root.Right.addLeft(13)
	root.Right.addRight(4)
	root.Right.Right.addRight(1)

	fmt.Println(hasPathSum(root, 22))
}
