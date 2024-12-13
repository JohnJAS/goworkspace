package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return createTree(nums, 0, len(nums)-1)
}

func createTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{
		nums[mid],
		createTree(nums, left, mid-1),
		createTree(nums, mid+1, right),
	}
	return root
}
