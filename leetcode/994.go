package main

import (
	"fmt"
)

func orangesRotting(grid [][]int) int {
	//计算边界
	R, C := len(grid), len(grid[0])
	//记录腐烂橘子的队列
	var queue [][2]int
	//初始化队列
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	ans := 0
	direction := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(queue) != 0 {
		length := len(queue)
		for length > 0 {
			curR, curC := queue[length-1][0], queue[length-1][1]
			queue = queue[1:]
			for _, v := range direction {
				newR, newC := curR+v[0], curC+v[1]
				if newR >= 0 && newR < R && newC >= 0 && newC < C && grid[newR][newC] == 1 {
					grid[newR][newC] = 2
					queue = append(queue, [2]int{newR, newC})
				}
			}
			length--
		}
		if len(queue) != 0 {
			ans++
		}
	}

	return ans
}

func main() {
	grid1 := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	grid2 := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	grid3 := [][]int{{0, 2}}

	fmt.Println(orangesRotting(grid1))
	fmt.Println(orangesRotting(grid2))
	fmt.Println(orangesRotting(grid3))
}
