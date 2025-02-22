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

func averageOfLevels(root *TreeNode) []float64 {
	result := make([]float64, 0)
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0
		l := len(queue)
		for i := 0; i < l; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
		average := float64(sum) / float64(l)
		result = append(result, average)
	}

	return result
}

func main() {
	root := &TreeNode{Val: 3}
	root.addLeft(9)
	root.addRight(20)
	root.Right.addRight(7)
	root.Right.addLeft(15)

	test := averageOfLevels(root)
	fmt.Println(test)
}
