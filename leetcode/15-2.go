package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; {
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				//move left to make sum increase
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum > 0 {
				//move right to make sum decrease
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else {
				//curTarget == sum
				//move both
				result = append(result, []int{nums[i], nums[l], nums[r]})
				//去重去重去重
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
		//第3个test case 要考虑的
		i++
		for i < len(nums)-2 && nums[i] == nums[i-1] {
			i++
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{0, 1, -1, -1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
