package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		tp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tp
	}

	//return is pre, not cur!
	return pre
}

func main() {
	fmt.Println(reverseList(nil))
}
