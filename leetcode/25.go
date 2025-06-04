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
	//函数框架与初始化
	dummy := &ListNode{Next: head} // 创建哨兵节点
	pre := dummy                   // pre 指向上一段的尾部，初始为 dummy
	cur := head                    // cur 指向当前段的头部

	//外层循环：逐组处理
	for cur != nil {
		tail := pre              // tail 初始化为 pre，为了下面迭代找到tail，确定[cur,...,tail]的区间，如果从cur往后推，tail会变成tail.Next
		for i := 0; i < k; i++ { // 向后走 k 步找本组的尾
			tail = tail.Next
			if tail == nil { // 不足 k 个，直接返回结果
				return dummy.Next
			}
		}

		//存下一组起点 + 反转本组
		next := tail.Next              // next 记录下一组的起点
		cur, tail = reverse(cur, tail) // 反转 [cur, tail]，返回新的头尾

		//把反转后的子链接回主链(链表题的“灵魂两行”)
		pre.Next = cur   // pre → 新头,把上一组（或 dummy）连到这组的新头；
		tail.Next = next // 新尾 → next（下一组的头）,把这组的新尾连回下一组的头，从而保持全链表连通。

		//更新指针，开始下一轮
		pre = tail      // pre 移到当前组的新尾,作为下一组的pre
		cur = tail.Next // cur 移到下一组的头
	}
	return dummy.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next  // pre 初始指向尾节点之后的位置（可能是 nil）
	cur := head       // cur 从 head 开始遍历到 tail
	for pre != tail { // 当 pre 走到 tail 才结束
		next := cur.Next
		cur.Next = pre // 反转指针
		pre = cur      // pre 与 cur 双指针前进
		cur = next
	}
	return tail, head // 反转后新的头是旧 tail，新尾是旧 head，不可以用cur,因为cur已经迭代下去了
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
