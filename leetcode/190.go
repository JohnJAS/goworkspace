package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1) // 取 num 的最低位，放到 result 的最高位
		num >>= 1                          // num 右移一位
	}
	return result
}

func main() {
	fmt.Println(reverseBits(5))
}
