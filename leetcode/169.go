package main

import "fmt"

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

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
