package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	dummy := &Node{Next: head}

	//创建副本
	for head != nil {
		head.Next = &Node{Val: head.Val, Next: head.Next}
		head = head.Next.Next
	}

	//设置random指针
	head = dummy.Next
	for head != nil {
		//注意randam可能为空
		if head.Random != nil {
			head.Next.Random = head.Random.Next
		}
		head = head.Next.Next
	}

	//创建链表
	head = dummy.Next
	headNew := head.Next
	for head != nil {
		nodeNew := head.Next
		head.Next = nodeNew.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
		head = head.Next
	}
	//注意回复head的值
	head = dummy.Next
	return headNew
}
