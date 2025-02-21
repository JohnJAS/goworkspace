package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (n *Node) addLeft(val int) {
	n.Left = &Node{Val: val}
}

func (n *Node) addRight(val int) {
	n.Right = &Node{Val: val}
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		tmpQueue := []*Node{}
		for i := 0; i < levelSize; i++ {
			if i == levelSize-1 {
				queue[i].Next = nil
			} else {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				tmpQueue = append(tmpQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmpQueue = append(tmpQueue, queue[i].Right)
			}
		}
		queue = tmpQueue
	}
	return root
}

func main() {
	root := &Node{Val: 1}
	root.addLeft(2)
	root.Left.addLeft(4)
	root.Left.addRight(5)
	root.addRight(3)
	root.Right.addRight(7)

	connect(root)

}
