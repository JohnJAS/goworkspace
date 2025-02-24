package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0]) //模
	l, r := 0, row*col-1
	for l <= r {
		mid := (l + r) / 2
		midValue := matrix[mid/col][mid%col]
		if midValue == target {
			return true
		} else if midValue > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(matrix, 3))
}
