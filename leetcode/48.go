package main

import "fmt"

func rotate1(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m/2; i++ {
		//奇数时，中心点不用动，行数或列数可以加1
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

func main() {
	m1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate1(m1)
	fmt.Println(m1)

}
