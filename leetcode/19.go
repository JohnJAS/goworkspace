package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}

	fast, slow := head, dummy

	for ; n > 0; n-- {
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	//delete node
	slow.Next = slow.Next.Next

	return dummy.Next
}
