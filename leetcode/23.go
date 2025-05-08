package main

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Status 定义优先队列的节点
type Status struct {
	node *ListNode
	val  int
}

// PriorityQueue 定义优先队列，并实现堆的接口
type PriorityQueue []*Status

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// 只需要实现添加到队尾
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Status))
}

// 只需要弹出最后一个元素，并返回塔
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// 把所有非空链表头结点放入最小堆
	for _, node := range lists {
		if node != nil {
			heap.Push(pq, &Status{node: node, val: node.Val})
		}
	}

	dummy := &ListNode{}
	tail := dummy

	for pq.Len() > 0 {
		s := heap.Pop(pq).(*Status)
		tail.Next = s.node
		tail = tail.Next

		//判断当前的队列是否为空
		if s.node.Next != nil {
			heap.Push(pq, &Status{node: s.node.Next, val: s.node.Next.Val})
		}
	}

	return dummy.Next
}
