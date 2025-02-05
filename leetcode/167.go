package main

import "fmt"

func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			//注意下标是从1开始计数的
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum2(numbers, target))
}
