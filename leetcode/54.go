package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows := len(matrix)
	cols := len(matrix[0])
	var result []int

	top, bottom, left, right := 0, rows-1, 0, cols-1

	for top <= bottom && left <= right {
		// move right
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		// move down
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// move left
		if top <= bottom { //注意判断上下界是否重合
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}

		// move up
		if left <= right { //=注意判断左右界是否重合
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	matrix := [][]int{{1, 2, 3}}
	fmt.Println(spiralOrder(matrix))
}
