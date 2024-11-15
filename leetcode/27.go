package main

import "fmt"

func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		if nums[right] == val {
			right--
		} else if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else {
			left++
		}
	}

	return left
}

func main() {
	var nums = []int{1}
	fmt.Println(removeElement(nums, 1))
	fmt.Println(nums)
}
