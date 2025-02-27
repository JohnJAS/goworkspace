package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	var result []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return result
	}

	var sCount, pCount [26]int
	// 初始化 pCount 和 sCount
	for i := 0; i < pLen; i++ {
		sCount[s[i]-'a']++ // 初始化 sCount，增加窗口內的字符數
		pCount[p[i]-'a']++ // 固定 p 的字符計數
	}

	// 檢查第一次窗口是否匹配
	if sCount == pCount {
		result = append(result, 0)
	}

	// 滑動窗口
	for i := 0; i < sLen-pLen; i++ {
		sCount[s[i]-'a']--      //滑动窗口左位置
		sCount[s[i+pLen]-'a']++ //滑动窗口右位置
		if sCount == pCount {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	s := "baa"
	p := "aa"

	fmt.Println(findAnagrams(s, p))
}
