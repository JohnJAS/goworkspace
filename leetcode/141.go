package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var p5 = ListNode{5, nil}
var p4 = ListNode{4, &p5}
var p3 = ListNode{3, &p4}
var p2 = ListNode{2, &p3}
var p1 = ListNode{1, &p2}
var HEAD = ListNode{1, &p1}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(hasCycle(&HEAD))
}
