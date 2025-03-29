/*
https://leetcode.cn/problems/two-sum/submissions/575569485/?envType=study-plan-v2&envId=top-100-liked
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//创建map, 存target-cur的值，向后搜索，如果有就说明找到了，因为只有唯一解，直接返回即可
	m := make(map[int]int)
	for index, v := range nums {
		find := target - v
		if _, ok := m[find]; ok {
			return []int{m[find], index}
		} else {
			m[v] = index
		}
	}

	return nil

}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
