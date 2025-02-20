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

// 先序遍历
func print(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf(" %d ", node.Val)
	print(node.Left)
	print(node.Right)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	rootIndex := -1
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
			break
		}
	}
	rootNode := &TreeNode{
		Val: rootVal,
	}
	leftTreeLength := rootIndex

	rootNode.Left = buildTree(inorder[0:leftTreeLength], postorder[0:leftTreeLength])
	rootNode.Right = buildTree(inorder[leftTreeLength+1:], postorder[leftTreeLength:len(postorder)-1])

	return rootNode

}

func main() {
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}

	//inorder中root的index
	rootIndex := 1
	leftTreeLength := rootIndex
	rightTreeLength := len(inorder) - rootIndex - 1
	inorderLeftTree := inorder[0:leftTreeLength]
	inorderRightTree := inorder[leftTreeLength+1:]
	fmt.Printf("leftTreeLength: %d\n", leftTreeLength)
	fmt.Printf("rightTreeLength: %d\n", rightTreeLength)
	fmt.Printf("inorderLeftTree: %v\n", inorderLeftTree)
	fmt.Printf("inorderRightTree: %v\n", inorderRightTree)

	postorderLeftTree := postorder[0:leftTreeLength]
	postorderRightTree := postorder[leftTreeLength : len(postorder)-1]

	fmt.Printf("postorderLeftTree: %v\n", postorderLeftTree)
	fmt.Printf("postorderRightTree: %v\n", postorderRightTree)

	root := buildTree(inorder, postorder)

	print(root)

}
