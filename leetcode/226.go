package main

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
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
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

	invertTree(root)
}
