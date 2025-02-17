package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(val int) *ListNode {
	if l == nil {
		return &ListNode{Val: val}
	}
	head := l
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{Val: val}
	return l // 返回原头节点，确保链表完整性
}

func (l *ListNode) print() {
	cur := l
	for cur != nil { // Use l != nil instead of l.Next != nil
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next // Move to the next node
	}
	fmt.Println() // To end the print with a new line
}

func partition(head *ListNode, x int) *ListNode {
	smallHead := &ListNode{}
	largeHead := &ListNode{}

	smallCur := smallHead
	largeCur := largeHead
	for head != nil {
		if head.Val >= x {
			largeCur.Next = head
			largeCur = largeCur.Next
		} else {
			smallCur.Next = head
			smallCur = smallCur.Next
		}
		head = head.Next
	}

	largeCur.Next = nil
	smallCur.Next = largeHead.Next

	return smallHead.Next
}

func main() {
	list1 := &ListNode{1, nil}
	list1.add(4).add(3).add(2).add(5).add(2)

	list1.print()
	fmt.Println(partition(list1, 3))

	list1.print()

}
