package main

import "fmt"

// i=0,j>0, dp[i][j] = dp[i][j-1] + grid[i][j]
// i>0,j=0, dp[i][j] = dp[i-1][j] + grid[i][j]
// i>0,j>0, dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	if rows == 0 || cols == 0 {
		return 0
	}

	// 正确初始化 dp 数组
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols) // 初始化每一行
	}

	// 初始化起点
	dp[0][0] = grid[0][0]

	// 初始化第一行（只能从左往右）
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 初始化第一列（只能从上往下）
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 计算最小路径和
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[rows-1][cols-1]
}

func minPathWithTrace(grid [][]int) (int, [][]int) {
	rows := len(grid)
	cols := len(grid[0])
	if rows == 0 || cols == 0 {
		return 0, nil
	}

	// 初始化 dp 数组
	dp := make([][]int, rows)
	trace := make([][]string, rows) // 记录路径方向
	for i := range dp {
		dp[i] = make([]int, cols)
		trace[i] = make([]string, cols)
	}

	// 初始化起点
	dp[0][0] = grid[0][0]

	// 计算第一行（只能从左往右）
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
		trace[0][j] = "L" // 记录从左边来
	}

	// 计算第一列（只能从上往下）
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
		trace[i][0] = "U" // 记录从上方来
	}

	// 计算 dp 数组，同时记录路径
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if dp[i-1][j] < dp[i][j-1] { // 从上面来
				dp[i][j] = dp[i-1][j] + grid[i][j]
				trace[i][j] = "U"
			} else { // 从左边来
				dp[i][j] = dp[i][j-1] + grid[i][j]
				trace[i][j] = "L"
			}
		}
	}

	// **回溯路径**
	path := [][]int{}
	x, y := rows-1, cols-1
	for x != 0 || y != 0 {
		path = append([][]int{{x, y}}, path...) // 逆序插入
		if trace[x][y] == "U" {
			x--
		} else {
			y--
		}
	}
	path = append([][]int{{0, 0}}, path...) // 添加起点

	return dp[rows-1][cols-1], path
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	sum, path := minPathWithTrace([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	fmt.Println("sum: ", sum)
	fmt.Printf("path: %v\n", path)
}
