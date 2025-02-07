package main

import "fmt"

func maxProfit2(prices []int) int {
	//只计算折线图上坡
	prohit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			prohit += prices[i] - prices[i-1]
		}
	}

	return prohit
}

func main() {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit2([]int{1, 2, 3, 4, 5}))
}
