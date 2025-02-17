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
