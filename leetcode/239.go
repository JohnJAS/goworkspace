package main

import (
	"container/heap"
	"fmt"
)

/**
1. 实现大顶堆的优先队列（堆）
2. 初始化优先队列，压入k个元素
3. 那么堆顶就是k个元素的最大值
4. 注意边界，窗口右移可能会使最大值出界，所以需要index记录最大值的位置，不在的情况，需要弹出堆顶
*/

type item struct {
	val   int
	index int
}

type highPriorityHeap []*item

func (hp highPriorityHeap) Len() int           { return len(hp) }
func (hp highPriorityHeap) Less(i, j int) bool { return hp[i].val > hp[j].val }
func (hp highPriorityHeap) Swap(i, j int)      { hp[i], hp[j] = hp[j], hp[i] }

func (hp *highPriorityHeap) Push(x interface{}) {
	*hp = append(*hp, x.(*item))
}

func (hp *highPriorityHeap) Pop() interface{} {
	old := *hp
	n := len(old)
	tail := old[n-1]
	*hp = old[0 : n-1]
	return tail
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int

	hp := &highPriorityHeap{}
	heap.Init(hp)
	for i := 0; i < k; i++ {
		heap.Push(hp, &item{val: nums[i], index: i})
	}
	result = append(result, (*hp)[0].val)

	for i := k; i < len(nums); i++ {
		heap.Push(hp, &item{val: nums[i], index: i})
		for (*hp)[0].index <= i-k {
			heap.Pop(hp)
		}
		result = append(result, (*hp)[0].val)
	}

	return result
}

func main() {
	list1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	list2 := []int{1}

	fmt.Println(maxSlidingWindow(list1, 3))
	fmt.Println(maxSlidingWindow(list2, 1))
}
