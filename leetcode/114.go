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

// 先序遍历
func print(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf(" %d ", node.Val)
	print(node.Left)
	print(node.Right)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = temp

}

func main() {
	root := &TreeNode{Val: 1}
	root.addLeft(2)
	root.Left.addLeft(3)
	root.Left.addRight(4)
	root.addRight(5)
	root.Right.addRight(6)
	print(root)
	fmt.Println()

	flatten(root)

	print(root)
}
