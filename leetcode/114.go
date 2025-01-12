package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 扁平化左右子树
	flatten(root.Left)
	flatten(root.Right)

	// 备份右子树
	right := root.Right

	// 将左子树接到右子树位置
	root.Right = root.Left
	root.Left = nil

	// 找到当前右子树的末尾节点
	current := root
	for current.Right != nil {
		current = current.Right
	}

	// 将原来的右子树接到末尾
	current.Right = right
}
