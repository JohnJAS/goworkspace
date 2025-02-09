/*
https://leetcode.cn/problems/two-sum/submissions/575569485/?envType=study-plan-v2&envId=top-100-liked
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{m[target-v], index}
		}
		m[v] = index
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
