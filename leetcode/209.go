package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	start, end, sum := 0, 0, 0
	minLength := math.MaxInt
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			minLength = min(minLength, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if minLength == math.MaxInt {
		return 0
	}
	return minLength
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}
