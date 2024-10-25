/*
https://leetcode.cn/problems/two-sum/submissions/575569485/?envType=study-plan-v2&envId=top-100-liked
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, value := range nums {

		if vv, exist := m[target-value]; exist {
			return []int{vv, index}
		}

		m[value] = index
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
