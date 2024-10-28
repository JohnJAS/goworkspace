package main

import "fmt"

/*
异或的特殊逻辑
0^n=n
n^n=0
b^a^b^c = a^b^b^c = a^c（这里面b和b异或为0，可以理解为相互抵消了）
*/
func singleNumber(nums []int) int {
	r := 0
	// 异或 a^0=a, a^a=0
	for _, v := range nums {
		r = r ^ v
	}
	return r
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 3}))
}
