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
	for i := 0; i < pLen; i++ {
		sCount[p[i]-'a']++ //固化p的pattern数组
		pCount[s[i]-'a']++ //初始化s的带比较数组
	}
	if sCount == pCount {
		result = append(result, 0)
	}

	for i := 0; i < sLen-pLen; i++ {
		sCount[p[i]-'a']--      //滑动窗口左位置
		sCount[s[i+pLen]-'a']++ //滑动窗口右位置
		if sCount == pCount {
			result = append(result, i+1)
		}
	}
	return result
}

func main() {
	s := "cbaebabacd"
	p := "abc"

	fmt.Println(findAnagrams(s, p))
}
