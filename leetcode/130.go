package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board[0]), len(board)
	//寻找左右和4边相邻的O，制成A
	//遍历左右两边
	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}
	//遍历上下两行
	for i := 0; i < m; i++ {
		dfs(board, 0, i)
		dfs(board, n-1, i)
	}

	//遍历矩阵
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'

	dfs(board, x, y+1)
	dfs(board, x, y-1)
	dfs(board, x-1, y)
	dfs(board, x+1, y)
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board)
}
