package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	//后一个数组[9]int用来代替map
	var row, col [9][9]int
	var box [3][3][9]int

	for i, r := range board {
		for j, c := range r {
			if board[i][j] == '.' {
				continue
			}
			row[i][c-'1']++
			col[j][c-'1']++
			box[i/3][j/3][c-'1']++
			if row[i][c-'1'] > 1 || col[j][c-'1'] > 1 || box[i/3][j/3][c-'1'] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}
