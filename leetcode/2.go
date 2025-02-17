package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//进位
	flag := 0
	var cur *ListNode
	//dummy head
	head := &ListNode{0, nil}
	pre := head
	for l1 != nil || l2 != nil {
		sum := flag
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		flag = sum / 10
		sum = sum % 10

		//创建当前节点的node
		cur = &ListNode{
			sum,
			nil,
		}
		pre.Next = cur
		//for next round
		pre = cur
	}
	if flag == 1 {
		cur.Next = &ListNode{1, nil}
	}

	return head.Next
}
