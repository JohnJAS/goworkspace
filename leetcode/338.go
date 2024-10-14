package main

import "fmt"

// BrianKernighan 算法的原理是：对于任意整数 x，令 x=x&(x−1)，该运算将 x 的二进制表示的从左往右数的最后一个 1 变成 0。因此，对 x 重复该操作，直到 x 变成 0，则操作次数即为 x 的「一比特数」。
func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}

func main() {
	fmt.Println(countBits(100000))
}
