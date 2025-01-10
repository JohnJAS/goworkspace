package main

import "fmt"

func buildLPSArray(pattern string) []int {
	m := len(pattern)
	arr := make([]int, m)
	//第一个字符的公共前后缀必为0，所以从第2个开始
	i := 1
	//最长公共前后缀长度
	length := 0
	for i < m {
		//抽象出来思考，下一个字母要么相等要么不同
		//那么如果相等，则最长公共字串的长度+1,并记录到LPS[i]位置，i后移
		if pattern[i] == pattern[length] {
			length++
			arr[i] = length
			i++
		} else if length > 0 {
			//那么如果不相等，其前一个子串的最长公共前后缀的记录length长度
			length = arr[length-1]
		} else {
			arr[i] = 0
			i++
		}
	}
	return arr
}

// KMPSearch 使用KMP算法查找模式字符串在文本中的位置
func strStr(text string, pattern string) int {
	next := buildLPSArray(pattern)

	m := len(text)
	//字符串指针
	i := 0
	//匹配串指针
	length := 0
	for i < m {
		if text[i] == pattern[length] {
			i++
			length++
		} else if length > 0 {
			length = next[length-1]
		} else {
			i++
		}
		if length == len(pattern) {
			//i 减去 匹配串的长度就是 结果的初始位置
			return i - length
		}
	}
	return -1
}

func main() {
	fmt.Println(buildLPSArray("aabaad"))
	fmt.Println(strStr("aabaabaad", "aabaad"))
}
