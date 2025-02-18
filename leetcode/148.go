package main

// 合并两个已排序的链表
func merge(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

// 分割链表为两部分
func split(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	slow := head
	fast := head
	var prev *ListNode

	// 快慢指针法分割链表
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil
	return head, slow
}

// 归并排序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, right := split(head)
	left = sortList(left)
	right = sortList(right)

	return merge(left, right)
}
