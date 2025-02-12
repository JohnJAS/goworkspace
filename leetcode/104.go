package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var l1 TreeNode = TreeNode{1, nil, nil}
var l2 TreeNode = TreeNode{2, nil, nil}
var root TreeNode = TreeNode{3, &l1, &l2}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	fmt.Println(maxDepth(&root))
}
