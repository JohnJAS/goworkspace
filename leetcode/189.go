package main

func rotate(nums []int, k int) {
	n := len(nums)
	newNums := make([]int, n)
	for i, v := range nums {
		newNums[(i+k)%n] = v
	}
	copy(nums, newNums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(nums, 3)
}
