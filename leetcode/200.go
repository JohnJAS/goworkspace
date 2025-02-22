package main

import "fmt"

func dfs(grid [][]byte, r, c int) {
	var row = len(grid)
	var col = len(grid[0])
	if r >= row || c >= col || r < 0 || c < 0 || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}

func numIslands(grid [][]byte) int {
	var count int
	var row = len(grid)
	var col = len(grid[0])
	if row == 0 || col == 0 {
		return count
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

func main() {
	var grid = [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '1'}}
	fmt.Println(numIslands(grid))
}
