package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	result := math.MinInt
	if len(height) <= 0 {
		return result
	}
	left, right := 0, len(height)-1
	for left < right {
		area := int(math.Min(float64(height[left]), float64(height[right]))) * (right - left)
		result = int(math.Max(float64(area), float64(result)))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
