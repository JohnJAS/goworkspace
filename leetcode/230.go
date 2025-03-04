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

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	var curNode = root
	var val int
	for k > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		val = curNode.Val
		curNode = curNode.Right
	}
	return val
}

func main() {
	root := &TreeNode{Val: 3}
	root.addLeft(1)
	root.Left.addRight(2)
	root.addRight(4)

	fmt.Println(kthSmallest(root, 3))
}
