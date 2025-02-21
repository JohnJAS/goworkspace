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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var list [][]int
	var level []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tempQ []*TreeNode
		node := queue[0]
		level = append(level, node.Val)
		queue = queue[1:]
		if node.Left != nil {
			tempQ = append(tempQ, node.Left)
		}
		if node.Right != nil {
			tempQ = append(tempQ, node.Right)
		}
		if len(queue) == 0 {
			queue = tempQ
			list = append(list, level)
			level = nil
		}
	}
	return list
}

func main() {
	root := &TreeNode{Val: 3}
	root.addLeft(9)
	root.addRight(20)
	root.Right.addLeft(15)
	root.Right.addRight(7)

	test := levelOrder(root)
	fmt.Println(test)
}
