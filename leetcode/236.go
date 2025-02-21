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

// 假设x为公共节点，x的ltree和rtree必包含p或q, 或者x为p或q的情况，q或p比在ltree或rtree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//找到左右的p和q
	if left != nil && right != nil {
		return root
	}

	//左边没找到，说明p,q都在右边
	if left == nil {
		return right
	}
	//否则就是在左边，或者为nil
	return left
}

func main() {
	root := &TreeNode{Val: 3}
	root.addLeft(5)
	root.Left.addLeft(6)
	root.Left.addRight(2)
	root.Left.Right.addLeft(7)
	root.Left.Right.addRight(4)
	root.addRight(1)
	root.Right.addLeft(0)
	root.Right.addRight(8)

	test := lowestCommonAncestor(root, root.Right, root.Right.Right)
	fmt.Println(test)
}
