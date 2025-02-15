package main

import "fmt"

func reverseList(head *ListNode) *ListNode {
	//定义空指针
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		//迭代下一轮
		pre = cur
		cur = next
	}
	//注意cur是nil，返回的是pre
	return pre
}

func main() {
	fmt.Println(reverseList(nil))
}
