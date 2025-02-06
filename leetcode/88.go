package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	//tail of result
	tail := m + n - 1
	//tail of nums1
	p1 := m - 1
	//tail of nums2
	p2 := n - 1

	for tail >= 0 {
		//这个代码块在前
		if p1 < 0 && p2 >= 0 {
			nums1[tail] = nums2[p2]
			p2--
		}
		if p2 < 0 && p1 >= 0 {
			nums1[tail] = nums1[p1]
			p1--
		}
		//这个代码块在后，否则p1--和p2--会影响逻辑
		if p1 >= 0 && p2 >= 0 {
			if nums1[p1] >= nums2[p2] {
				nums1[tail] = nums1[p1]
				p1--
			} else {
				nums1[tail] = nums2[p2]
				p2--
			}
		}

		tail--
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
