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
	return verify(root.Left, root.Right)
}

func verify(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil {
		return false
	}
	return verify(l.Left, r.Right) && verify(l.Right, r.Left) && l.Val == r.Val
}

func main() {
	root := &TreeNode{Val: 5}
	root.addLeft(3)
	root.Left.addLeft(1)
	root.addRight(3)
	root.Right.addRight(1)

	fmt.Println(isSymmetric(root))
}
