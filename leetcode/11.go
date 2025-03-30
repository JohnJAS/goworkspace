package main

import "fmt"

func maxArea(height []int) int {
	//双指针，双指针，左右逼近，原理，面积是由低的那一边决定的，所以移动低的那一边，才有可能找到更大的面积
	maxA := 0
	leftP, rightP := 0, len(height)-1
	for leftP < rightP {
		if height[leftP] < height[rightP] {
			maxA = max(maxA, height[leftP]*(rightP-leftP))
			//面积是由低的那一边决定的，所以移动低的那一边，才有可能找到更大的面积
			leftP++
		} else {
			maxA = max(maxA, height[rightP]*(rightP-leftP))
			rightP--
		}
	}
	return maxA
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
