package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i] = digits[i] + 1
		}
		n := (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		digits[i] = n
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}
