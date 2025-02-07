package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxprofit := 0
	for _, v := range prices {
		minPrice = min(minPrice, v)
		prohit := v - minPrice
		maxprofit = max(maxprofit, prohit)
	}
	return maxprofit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
