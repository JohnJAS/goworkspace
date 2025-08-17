package main

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil {
		for cur.Next != nil {
			if cur.Val == cur.Next.Val {
				cur.Next = cur.Next.Next
			} else {
				break
			}
		}

		cur = cur.Next
	}
	return head
}
