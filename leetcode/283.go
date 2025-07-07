package main

import "fmt"

func moveZeroes(nums []int) {
	fast, slow := 0, 0

	for fast < len(nums) {
		if nums[fast] == 0 {
			if nums[slow] != 0 {
				slow = fast
			}
		} else {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			for slow <= fast && nums[slow] != 0 {
				slow++
			}
		}

		fast++
	}

}

func main() {
	a := []int{1, 0, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}
