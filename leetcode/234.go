package main

import "fmt"

func isPalindrome2(head *ListNode) bool {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	p := 0
	q := len(list) - 1

	for p <= q {
		if list[p] != list[q] {
			return false
		}
		p++
		q--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome2(linkedList))
}
