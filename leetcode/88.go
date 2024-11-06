package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	tail := m + n - 1

	for ; tail >= 0; tail-- {
		if p1 == -1 {
			nums1[tail] = nums2[p2]
			p2--
		} else if p2 == -1 {
			nums1[tail] = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
		} else {
			nums1[tail] = nums2[p2]
			p2--
		}
	}

}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)

}
