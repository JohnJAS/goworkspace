package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome2(head *ListNode) bool {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	p := 0
	q := len(list) - 1

	for p <= q {
		if list[p] != list[q] {
			return false
		}
		p++
		q--
	}
	return true
}

var linkedList *ListNode

func (l *ListNode) add(val int) {
	l.Next = &ListNode{Val: val}
}

func main() {
	linkedList = &ListNode{Val: 1}
	linkedList.add(2)
	linkedList.add(2)
	linkedList.add(1)

	fmt.Println(isPalindrome2(linkedList))
}
