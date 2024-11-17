package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 1, 1
	for fast < n {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 1, 1, 2, 3, 4, 5}))
}
