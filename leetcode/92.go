package main

import "fmt"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}

	//遍历找到左侧节点左边一位的节点
	var preLeft = dummy
	for i := 0; i < left-1; i++ {
		preLeft = preLeft.Next
	}

	//继续遍历找到右侧节点的地址
	var end = preLeft
	for i := 0; i < right-left+1; i++ {
		end = end.Next
	}

	//保存头尾节点地址
	headNew := preLeft.Next
	endRight := end.Next
	//截取出待反转的中间链表
	preLeft.Next = nil
	end.Next = nil

	//反转中间链表
	headNew, end = reverseList2(headNew)

	//链接链表
	preLeft.Next = headNew
	end.Next = endRight

	return dummy.Next
}

func reverseList2(head *ListNode) (newHead, newTail *ListNode) {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		//for next round
		pre = cur
		cur = tmp
	}

	return pre, head
}

func main() {
	fmt.Println(reverseBetween(nil))
}
