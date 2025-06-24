package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	slow, fast := 0, 1
	for ; fast <= len(nums)-1; fast++ {
		if nums[fast] == nums[slow] {
			continue
		}
		slow++
		nums[slow] = nums[fast]

	}
	return slow + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 1, 1, 2, 3, 4, 5}))
}
