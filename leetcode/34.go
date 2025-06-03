package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	i := sort.SearchInts(nums, target)
	if i == len(nums) || nums[i] != target {
		return []int{-1, -1}
	}
	j := sort.SearchInts(nums, target+1) - 1
	return []int{i, j}
}

func main() {
	nums1 := []int{5, 7, 7, 8, 8, 10}
	nums2 := []int{5, 7, 7, 8, 8, 10}
	nums3 := []int{1}
	fmt.Println(searchRange(nums1, 8))
	fmt.Println(searchRange(nums2, 6))
	fmt.Println(searchRange(nums3, 1))
}
