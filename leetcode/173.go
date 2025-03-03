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

type BSTIterator struct {
	Stack   []*TreeNode
	CurNode *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		Stack:   make([]*TreeNode, 0),
		CurNode: root,
	}
}

func (this *BSTIterator) Next() int {
	//中序遍历
	for node := this.CurNode; node != nil; node = node.Left {
		this.Stack = append(this.Stack, node)
	}
	this.CurNode = this.Stack[len(this.Stack)-1]
	val := this.CurNode.Val
	this.Stack = this.Stack[:len(this.Stack)-1]

	//由于二叉搜索树的特性，比cur大子节点只存在右侧
	this.CurNode = this.CurNode.Right

	return val
}

func (this *BSTIterator) HasNext() bool {
	//处于根节点的时候Len为0，但此时指向右节点
	if len(this.Stack) == 0 && this.CurNode == nil {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	root := &TreeNode{Val: 7}
	root.addLeft(3)
	root.addRight(15)
	root.Right.addLeft(9)
	root.Right.addRight(20)

	test := Constructor(root)
	fmt.Println(test.Next())
	fmt.Println(test.Next())
	fmt.Println(test.HasNext())
	fmt.Println(test.Next())
	fmt.Println(test.HasNext())
	fmt.Println(test.Next())
	fmt.Println(test.HasNext())
	fmt.Println(test.Next())
	fmt.Println(test.HasNext())
}
