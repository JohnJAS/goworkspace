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

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

func main() {
	root := &TreeNode{Val: 5}
	root.addLeft(3)
	root.Left.addLeft(1)
	root.addRight(3)
	root.Right.addRight(1)

	fmt.Println(isSymmetric(root))
}
