package main

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}
	if root.Left != nil {
		//...是一个解包操作，表示将 slice2 中的所有元素逐一追加到 slice1
		result = append(inorderTraversal(root.Left), result...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}

	return result
}
