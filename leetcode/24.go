package main

import "fmt"

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	pre, cur := dummy, head

	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = cur

		pre = pre.Next.Next
		cur = cur.Next
	}

	return dummy.Next
}

func main() {
	root := &ListNode{Val: 1}
	root.add(2)
	root.add(3)

	swapPairs(root).print()
}
