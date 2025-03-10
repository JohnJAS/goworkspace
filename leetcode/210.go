package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	var result []int
	indeg := make([]int, numCourses)
	edges := make([][]int, numCourses)
	var queue []int
	//虽然是个二维数组，但是每一行的元素只有2列，所以没有必要遍历列
	for _, row := range prerequisites {
		indeg[row[0]]++
		edges[row[1]] = append(edges[row[1]], row[0])
	}

	//怎么处理第一个?
	//遍历入度表，先找到找到值为零的
	for i, v := range indeg {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		head := queue[0]
		result = append(result, head)
		queue = queue[1:]
		//查询出栈的边，有没有对应的后续课程，有的话，入度-1，并且查询入度是否为0
		for _, next := range edges[head] {
			indeg[next]--
			if indeg[next] == 0 {
				queue = append(queue, next)
			}
		}

	}
	if len(result) != numCourses {
		return nil
	}

	return result
}

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
