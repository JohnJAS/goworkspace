package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	var maxVal = math.MinInt64
	var index = -1
	for i, v := range nums {
		if v > maxVal {
			maxVal = v
			index = i
		}
	}
	return index
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(arr))
}
