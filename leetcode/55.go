package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums) - 1

	gas := 0

	for _, num := range nums {
		gas = max(gas, num)
		if gas >= n {
			continue
		}
		if gas == 0 {
			return false
		}
		gas--
		n--
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
