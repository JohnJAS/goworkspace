package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	fmt.Println(citations)

	h := 0
	//第1篇1次，第二篇2次，第3篇3次...第h篇h次
	for _, v := range citations {
		if v > h {
			h++
		}
	}

	return h
}

func main() {
	fmt.Println(hIndex([]int{1, 2, 3, 4, 5}))
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{100000, 0, 5, 5, 5, 5}))
}
