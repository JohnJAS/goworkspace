package main

import "fmt"

// 哈希 时间复杂度和空间复杂度都是O(n)
func singleNumber0(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k

		}
	}

	return 0
}

func singleNumber(nums []int) int {
	//注意结果可能是负数
	var result int32
	for i := 0; i < 32; i++ {
		var total int32
		for _, v := range nums {
			total += (int32(v) >> i) & 1
		}
		result = (total%3)<<i | result
	}

	return int(result)
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, -4, 2, 1, 1, 1, 4, 4, 4}))
}
