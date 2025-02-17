package main

import "fmt"

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}

func main() {
	var nums = []int{2, 3, 1, 2, 1}
	fmt.Println(removeElement(nums, 1))
	fmt.Println(nums)
}
