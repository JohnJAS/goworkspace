package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(val int) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{Val: val}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, list1}
	pre := dummy

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			pre.Next = list2
			tempList2 := list2
			list2 = list2.Next
			tempList2.Next = list1
			pre = tempList2
		} else {
			pre = list1
			list1 = list1.Next
		}
	}

	if list2 != nil {
		pre.Next = list2
	}

	return dummy.Next
}

func main() {
	list1 := &ListNode{1, nil}
	list1.add(3)
	list1.add(4)
	list1.add(7)
	list1.add(9)

	list2 := &ListNode{1, nil}
	list2.add(2)
	list2.add(9)
	list2.add(11)
	list2.add(12)

	test := mergeTwoLists(list1, list2)
	for test != nil {
		fmt.Println(test.Val)
		test = test.Next
	}
}
