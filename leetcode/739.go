package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	//单调栈，从底向上温度递减
	stack := []int{}
	//初始化为全0数组
	ans := make([]int, length)

	for i := 0; i < length; i++ {
		//如果栈不为空，且当前温度大于栈顶温度，则出栈
		//出栈的值preindex,计算得出i-predex,即为ans[preindex]应该记录的值
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		//否则则入栈当前温度的index
		stack = append(stack, i)
	}

	return ans
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
