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

func isValidBST(root *TreeNode) bool {
	var curNode = root
	var stack []*TreeNode
	var preVal int //也可以初始化为如下
	//preVal := math.MinInt64
	var count int

	for curNode != nil || len(stack) != 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++
		//跳过第一个
		if count == 1 {
			preVal = curNode.Val
			curNode = curNode.Right
			continue
		}
		if curNode.Val <= preVal {
			return false
		}
		preVal = curNode.Val
		curNode = curNode.Right
	}
	return true
}

func main() {
	root := &TreeNode{Val: 5}
	root.addLeft(1)
	root.addRight(3)
	fmt.Println(isValidBST(root))

}
