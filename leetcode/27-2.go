package main

import "fmt"

func removeElement(nums []int, val int) int {
	left := 0
	for i := range nums {
		if nums[i] != val {
			nums[left] = nums[i]
			left++
		}
	}
	return left
}

func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] != val {
			left++
		} else {
			nums[left] = nums[right]
			right--
		}
	}

	return left
}

func main() {
	var nums = []int{2, 3, 1, 2, 1}
	fmt.Println(removeElement2(nums, 1))
	fmt.Println(nums)
}
