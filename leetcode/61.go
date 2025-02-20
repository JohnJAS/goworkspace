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

func rotateRight(head *ListNode, k int) *ListNode {
	//不移动，或者链表为空
	if k == 0 || head == nil {
		return head
	}
	l, tail := getLength(head)
	mod := k % l
	//移动次数为链表长度的整数倍
	if mod == 0 {
		return head
	}
	//形成环
	tail.Next = head
	//计算环断开的地方
	pre := &ListNode{}
	for move := l - mod; move > 0; move-- {
		pre = head
		head = head.Next
	}
	pre.Next = nil
	return head
}

func getLength(head *ListNode) (count int, tail *ListNode) {
	tail = head
	for head != nil {
		count++
		tail = head
		head = head.Next
	}
	return
}

func main() {
	list1 := &ListNode{1, nil}
	list1.add(2)
	list1.add(3)
	list1.add(4)
	list1.add(5)
	list1.print()
	list1 = rotateRight(list1, 2)
	list1.print()

	list2 := &ListNode{0, nil}
	list2.add(1)
	list2.add(2)
	list2.print()
	list2 = rotateRight(list2, 4)
	list2.print()
}
