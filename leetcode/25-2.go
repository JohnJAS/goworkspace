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

func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{Next: head}
	pre := dummy
	cur := pre.Next

	for cur != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		next := tail.Next

		cur, tail = reverse(cur, tail)

		pre.Next = cur
		tail.Next = next

		pre = tail
		cur = next

	}
	return dummy.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {

	pre := tail.Next
	cur := head

	for pre != tail {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next

	}
	return tail, head
}

func main() {
	list1 := &ListNode{1, nil}
	list1.add(2)
	list1.add(3)
	list1.add(4)
	list1.add(5)
	list1.print()

	head := reverseKGroup(list1, 2)
	head.print()

}
