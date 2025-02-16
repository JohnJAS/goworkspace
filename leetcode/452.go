package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var r [][]int
	for i := 0; i < len(points); i++ {
		if len(r) == 0 {
			r = append(r, points[i])
		} else {
			if points[i][0] > r[len(r)-1][1] {
				r = append(r, points[i])
			} else if points[i][1] < r[len(r)-1][1] {
				r[len(r)-1][0] = points[i][0]
				r[len(r)-1][1] = points[i][1]
			} else {
				r[len(r)-1][0] = points[i][0]
			}
		}
	}
	return len(r)
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
}
