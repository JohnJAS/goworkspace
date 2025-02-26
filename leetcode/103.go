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

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	flag := true
	for len(queue) != 0 {
		l := len(queue)
		var subResult []int
		var tmpQ []*TreeNode
		for i := 0; i < l; i++ {
			if flag {
				subResult = append(subResult, queue[i].Val)
			} else {
				subResult = append(subResult, queue[l-i-1].Val)
			}

			if queue[i].Left != nil {
				tmpQ = append(tmpQ, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmpQ = append(tmpQ, queue[i].Right)
			}

		}
		queue = queue[l:]
		result = append(result, subResult)
		//reverse flag
		flag = !flag
		if len(tmpQ) != 0 {
			queue = tmpQ
		}
	}
	return result
}

func main() {
	root := &TreeNode{Val: 3}
	root.addLeft(9)
	root.addRight(20)
	root.Right.addLeft(15)
	root.Right.addRight(7)

	test := zigzagLevelOrder(root)
	fmt.Println(test)

	fmt.Println(zigzagLevelOrder(nil))

}
