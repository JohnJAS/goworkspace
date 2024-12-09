package main

import "fmt"

func setZeroes(matrix [][]int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))

	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if row[i] == 1 || col[j] == 1 {
				r[j] = 0
			}
		}
	}
	fmt.Println(matrix)
}

func main() {
	martix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(martix)
}
