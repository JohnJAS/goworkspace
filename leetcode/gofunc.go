package main

import "sort"

func main() {
	var nums []int = []int{1, 2, 3, 3, 4, 5}

	sort.Slice(nums, func(i, j int) bool {
		return i < j
	})

	sort.Ints(nums)
}
