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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
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

	rootNode.Left = buildTree(preorder[1:leftTreeLength+1], inorder[0:leftTreeLength])
	rootNode.Right = buildTree(preorder[leftTreeLength+1:], inorder[leftTreeLength+1:])

	return rootNode
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	rootIndex := 1
	leftTreeLength := rootIndex
	rightTreeLength := len(inorder) - rootIndex - 1
	inorderLeftTree := inorder[0:leftTreeLength]
	inorderRightTree := inorder[leftTreeLength+1:]
	fmt.Printf("leftTreeLength: %d\n", leftTreeLength)
	fmt.Printf("rightTreeLength: %d\n", rightTreeLength)
	fmt.Printf("inorderLeftTree: %v\n", inorderLeftTree)
	fmt.Printf("inorderRightTree: %v\n", inorderRightTree)

	preorderLeftTree := preorder[1 : leftTreeLength+1]
	preorderRightTree := preorder[leftTreeLength+1:]

	fmt.Printf("preorderLeftTree: %v\n", preorderLeftTree)
	fmt.Printf("preorderRightTree: %v\n", preorderRightTree)

	root := buildTree(preorder, inorder)
	print(root)
}
