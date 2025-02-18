package main

import (
	"fmt"
)

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

/*
 * 完全二叉树 ： 从上到下，从左到右排列一棵树
 * 所以，左子树高度最多比右子树高度多1，或者相等
 * 不可能出现右子树比左子树高的情况
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	if leftHeight == rightHeight {
		// 左子树高度 = 右子树高度，则左子树必定是满二叉树
		//左子树的节点数 + 右子树的节点数 + 根节点
		// 1<<n,就是2^n,
		return (1 << leftHeight) - 1 + countNodes(root.Right) + 1
	} else {
		// 左子树高度 > 右子树高度，说明右子树必定是满二叉树
		return (1 << rightHeight) - 1 + countNodes(root.Left) + 1
	}
}

func getHeight(root *TreeNode) (h int) {
	if root == nil {
		return h
	}
	for root != nil {
		h++
		root = root.Left
	}
	return
}

// 递归求解 O(nlogn)
//func countNodes(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return countNodes(root.Left) + countNodes(root.Right) + 1
//}

func main() {
	root := &TreeNode{Val: 5}
	fmt.Println(countNodes(root))

}
