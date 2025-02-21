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
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curLevelSize := len(queue)
		var level []int
		for curLevelSize > 0 {
			node := queue[0]
			queue = queue[1:] // dequeue the node
			level = append(level, node.Val)
			// enqueue left and right children
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			curLevelSize--
		}
		list = append(list, level)
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
