package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next

	for slow != fast {
		if slow != nil {
			slow = slow.Next
		} else {
			return false
		}
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
	}
	return true
}
