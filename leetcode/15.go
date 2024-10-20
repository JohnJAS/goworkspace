package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	//sort, O(nlog(n))
	sort.Ints(nums)

	for p1 := 0; p1 <= len(nums)-2; p1++ {
		if p1 > 0 && nums[p1] == nums[p1-1] {
			continue
		}
		p3 := len(nums) - 1
		for p2 := p1 + 1; p2 <= len(nums)-1; p2++ {
			if p2 > p1+1 && nums[p2] == nums[p2-1] {
				continue
			}
			for nums[p1]+nums[p2]+nums[p3] > 0 && p2 < p3 {
				p3--
			}
			if p2 == p3 {
				break
			}
			if nums[p1]+nums[p2]+nums[p3] == 0 {
				result = append(result, []int{nums[p1], nums[p2], nums[p3]})
			}

		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
