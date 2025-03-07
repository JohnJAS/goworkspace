package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	//排除空图
	if node == nil {
		return nil
	}
	//访问过的节点map
	var visitedNodeMap = make(map[*Node]*Node)
	//初始化访问过的节点
	visitedNodeMap[node] = &Node{Val: node.Val}
	//初始化要遍历的neighbor节点
	queue := []*Node{node}

	for len(queue) > 0 {
		//取队列中的第一个原始节点
		curNode := queue[0]
		queue = queue[1:]
		for _, neighbor := range curNode.Neighbors {
			if _, ok := visitedNodeMap[neighbor]; !ok {
				//添加访问过的节点
				visitedNodeMap[neighbor] = &Node{Val: neighbor.Val}
				//添加要遍历的neighbor节点
				queue = append(queue, neighbor)
			}
			//添加clone后的neighbor节点到克隆节点中
			visitedNodeMap[curNode].Neighbors = append(visitedNodeMap[curNode].Neighbors, visitedNodeMap[neighbor])
		}
	}

	return visitedNodeMap[node]
}
