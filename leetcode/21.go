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
	return l
}

func (l *ListNode) print() {
	cur := l
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummy.Next
}

func main() {
	list1 := &ListNode{1, nil}
	list1.add(3)
	list1.add(4)
	list1.add(7)
	list1.add(9)
	list1.print()

	list2 := &ListNode{1, nil}
	list2.add(2)
	list2.add(9)
	list2.add(11)
	list2.add(12)
	list2.print()

	test := mergeTwoLists(list1, list2)
	test.print()
}
