package main

import (
	"fmt"
	"sort"
)

func merge2(intervals [][]int) [][]int {
	var result [][]int

	//sort the intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	fmt.Println(intervals)

	for i := 0; i < len(intervals); i++ {
		if i == 0 {
			result = append(result, intervals[i])
		}
		if i > 0 {
			if result[len(result)-1][0] == intervals[i][0] {
				result[len(result)-1][1] = intervals[i][1]
				continue
			}
			if result[len(result)-1][1] < intervals[i][0] {
				result = append(result, intervals[i])
			} else {
				if result[len(result)-1][1] <= intervals[i][1] {
					result[len(result)-1][1] = intervals[i][1]
				} else {
					continue
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(merge2([][]int{{1, 2}, {3, 4}, {3, 4}, {0, 3}, {0, 1}, {0, 2}, {5, 6}, {5, 6}}))
}
