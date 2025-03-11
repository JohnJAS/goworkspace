package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, used, &res)
	return res
}

func backtrack(nums, path []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}
	for i := range nums {
		if !used[i] {
			used[i] = true
			path = append(path, nums[i])
			backtrack(nums, path, used, res)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
