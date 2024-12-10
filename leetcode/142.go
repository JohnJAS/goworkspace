
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			//has Circle
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}

	}

	return nil
}