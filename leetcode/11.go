package main

import "fmt"

func maxArea(height []int) int {
	var maxA int

	left, right := 0, len(height)-1

	for left < right {
		area := 0
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}
		if maxA < area {
			maxA = area
		}
	}

	return maxA
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
