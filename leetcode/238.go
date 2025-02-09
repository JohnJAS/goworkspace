package main

import "fmt"

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))

	//求前缀积，当前元素左侧所有数的乘积
	//用answer暂存前缀积表
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	R := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= R
		R *= nums[i]
	}

	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
}
