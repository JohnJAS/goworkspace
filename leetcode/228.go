package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var result []string

	var slow, fast int
	for slow < len(nums) {
		tempStr := ""
		for fast < len(nums) {
			if slow == fast {
				tempStr = strconv.Itoa(nums[slow])
				fast++
				continue
			}
			if nums[fast] == nums[fast-1]+1 {
				fast++
				continue
			} else {
				break
			}
		}
		if fast-1 != slow {
			tempStr = tempStr + "->" + strconv.Itoa(nums[fast-1])
		}
		slow = fast
		result = append(result, tempStr)
	}

	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
