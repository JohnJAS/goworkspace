package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	count := 0
	var maxj int
	for i := 0; i <= len(nums)-1; i++ {
		if count == 0 {
			maxj = nums[i]
		}
		if maxj == nums[i] {
			count++
		} else {
			count--
		}
	}

	return maxj
}

func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[(len(nums)-1)/2]
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 1, 3, 1, 4, 1, 1, 5, 5}))
	fmt.Println(majorityElement2([]int{1, 2, 1, 3, 1, 4, 1, 1, 5, 5}))
}
