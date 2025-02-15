package main

import "fmt"

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

func (l *ListNode) print() {
	for l != nil {
		fmt.Printf(" %d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			//cur节点的下一挑跳过所有val为x的值
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			//正常迭代到下一个
			cur = cur.Next
		}
	}

	return dummy.Next
}

func main() {
	list1 := &ListNode{1, nil}
	list1.add(2)
	list1.add(3)
	list1.add(3)
	list1.add(3)
	list1.add(4)
	list1.add(4)
	list1.add(4)
	list1.add(4)
	list1.add(5)
	list1.print()

	list1 = deleteDuplicates(list1)
	list1.print()
}
