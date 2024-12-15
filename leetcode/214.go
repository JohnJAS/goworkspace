package main

import "fmt"

func isHappy(n int) bool {
	slow, fast := bitSquareSum(n), bitSquareSum(bitSquareSum(n))

	for slow != fast {
		slow = bitSquareSum(slow)
		fast = bitSquareSum(bitSquareSum(fast))
	}

	return slow == 1 // 如果相遇时值为1，说明是快乐数
}

func bitSquareSum(n int) (sum int) {
	for n > 0 {
		m := n % 10
		sum = sum + m*m
		n = n / 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
}
