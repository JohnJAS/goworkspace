package main

// type TreeNode struct {
//     Val   int
//     Left  *TreeNode
//     Right *TreeNode
// }

func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	depth(root, &result)
	return result
}

func depth(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}

	l := depth(root.Left, result)
	r := depth(root.Right, result)

	// 更新直径
	*result = max(*result, l+r)

	// 返回当前节点的深度
	return max(l, r) + 1
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
