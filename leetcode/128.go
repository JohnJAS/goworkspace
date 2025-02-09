package main

import (
	"fmt"
	"sort"
)

// 排序
func longestConsecutive(nums []int) int {
	sort.Ints(nums)

	count := 0
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			count++
			continue
		}
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			count++
		} else {
			maxCount = max(maxCount, count)
			count = 1
		}
	}
	return max(maxCount, count)
}

// 哈希
func longestConsecutive2(nums []int) int {
	m := make(map[int]bool)

	for _, v := range nums {
		m[v] = true
	}

	count := 0
	maxCount := 0

	for key, _ := range m {
		if !m[key-1] {
			count = 1
			currentNum := key
			for m[currentNum+1] {
				count++
				currentNum++
			}
			maxCount = max(maxCount, count)
		}
	}

	return maxCount
}

func main() {
	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
	fmt.Println(longestConsecutive([]int{0}))
	fmt.Println(longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}))

	fmt.Println(longestConsecutive2([]int{1, 2, 0, 1}))
	fmt.Println(longestConsecutive2([]int{0}))
	fmt.Println(longestConsecutive2([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}))
}
