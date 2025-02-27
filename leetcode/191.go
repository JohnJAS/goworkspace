package main

import "fmt"

func hammingWeight(n int) int {
	var count int
	for i := 0; i < 32; i++ {
		if 1<<i&n > 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(1 << 0)
	fmt.Println(1 << 1)
	fmt.Println(1 << 2)
	fmt.Println(hammingWeight(11))
}
