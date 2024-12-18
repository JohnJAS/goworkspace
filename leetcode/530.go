package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt64
	pre := -1
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		helper(node.Right)
	}

	helper(root)
	return ans
}
