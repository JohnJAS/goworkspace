package main

import "fmt"

func jump(nums []int) int {
	count := 0
	//当前次跳跃的最远点
	end := 0
	//在遍历过程中能达到的最远点
	maxPosition := 0

	for i, num := range nums {
		maxPosition = max(maxPosition, i+num)
		//最后一位不用遍历到
		if i == end && i != len(nums)-1 {
			end = maxPosition
			count++
		}

	}

	return count
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}
