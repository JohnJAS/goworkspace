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

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	sum = sum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}
	return dfs(node.Left, sum) + dfs(node.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(sumNumbers(root))
}
