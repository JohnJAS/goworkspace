package main

import (
	"fmt"
	"sort"
)

func merge2(intervals [][]int) [][]int {
	var result [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 || result[len(result)-1][1] < intervals[i][0] {
			result = append(result, intervals[i])
			continue
		}
		if result[len(result)-1][1] > intervals[i][1] {
			continue
		}
		result[len(result)-1][1] = intervals[i][1]
	}
	return result
}

func main() {
	fmt.Println(merge2([][]int{{1, 2}, {3, 4}, {3, 4}, {0, 3}, {0, 1}, {0, 2}, {5, 6}, {5, 6}}))
}
