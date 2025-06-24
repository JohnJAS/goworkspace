package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	var pre *ListNode
	cur := dummy
	tail := dummy
	//get head
	for i := 0; i < left; i++ {
		pre = cur
		cur = cur.Next
	}
	//get tail
	for i := 0; i < right; i++ {
		tail = tail.Next
	}
	next := tail.Next

	cur, tail = reverse(cur, tail)

	//重新接上链表
	pre.Next = cur
	tail.Next = next

	return dummy.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	cur := head

	for pre != tail {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next

	}
	return tail, head
}

func main() {}
