package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			return mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
