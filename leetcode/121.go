package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var lowPrice int
	var maxProfit int

	lowPrice = prices[0]

	for _, i := range prices {
		lowPrice = min(lowPrice, i)
		if i-lowPrice > maxProfit {
			maxProfit = i - lowPrice
		}

	}
	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
