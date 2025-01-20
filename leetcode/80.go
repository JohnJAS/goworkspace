package main

import "fmt"

func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	var slow, fast = 2, 2
	for fast < length {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

func main() {
	fmt.Println(removeDuplicates2([]int{1, 1, 1, 2, 2, 3}))
}
