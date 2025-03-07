package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses) // 邻接表，存储课程的依赖关系
		indeg  = make([]int, numCourses)   // 入度表，记录每门课程的前置课程数量
		result []int                       // 存储拓扑排序结果
	)

	//构建课程依赖关系图
	for _, info := range prerequisites {
		//建立有向图，题中的[a,b] 表示b->a 的边，这里就是在统计所有到a的节点
		edges[info[1]] = append(edges[info[1]], info[0])
		//入度计算
		indeg[info[0]]++ //入度自增
	}

	//找到所有“无前置要求”的课程，即入度为0的，加入队列
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}

	//拓扑排序（Kahn算法）
	for len(queue) > 0 {
		u := queue[0]              // 取出队列中的课程
		queue = queue[1:]          // 从队列中移除
		result = append(result, u) // 记录学习的课程
		for _, v := range edges[u] {
			indeg[v]--         // 课程 v 的前置课程数量减少
			if indeg[v] == 0 { // 如果前置课程全部学完，入度为0，表示可以学习，则加入队列
				queue = append(queue, v)
			}
		}
	}

	return len(result) == numCourses
}

func main() {
	test := [][]int{{1, 0}}
	fmt.Println(canFinish(2, test))
	fmt.Println(test[0])
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
