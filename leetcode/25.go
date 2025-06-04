package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(val int) *ListNode {
	if l == nil {
		return &ListNode{Val: val}
	}
	head := l
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{Val: val}
	return l // 返回原头节点，确保链表完整性
}

func (l *ListNode) print() {
	cur := l
	for cur != nil { // Use l != nil instead of l.Next != nil
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next // Move to the next node
	}
	fmt.Println() // To end the print with a new line
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for cur != nil {
		tail := pre
		// 修复1：使用 k 而不是硬编码的 2
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		next := tail.Next
		// 调用修复后的 reverse 函数
		cur, tail = reverse(cur, tail)

		// 重新连接链表
		pre.Next = cur
		tail.Next = next

		// 更新指针，准备下一组
		pre = tail
		cur = tail.Next
	}

	return dummy.Next
}

// 修复2：循环条件使用 tail.Next 而不是 nil
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	cur := head
	for cur != tail.Next {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}

func main() {
	list1 := &ListNode{1, nil}
	list1.add(2)
	list1.add(3)
	list1.add(4)
	list1.add(5)
	list1.print()

	head := reverseKGroup(list1, 2)
	head.print()

}
