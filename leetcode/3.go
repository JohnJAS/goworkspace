package main

import "fmt"

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	//顺序遍历一次字符串
	//记录每个字符到map
	//**这其实是个滑动窗口的题目，所以要记录左右位置，否则在移动时，无法知道左边界位置**
	var maxLength int
	m := map[byte]int{}
	l := len(s)
	startIndex := 0
	for i := 0; i < l; i++ {
		//判断字符是否存在
		if _, ok := m[s[i]]; ok {
			//********我这里错了一次，注意左边界是删除包括重复字符左边的所有字符，不是单单删除重复字符的***********//
			for startIndex <= m[s[i]] {
				delete(m, s[startIndex])
				startIndex++
			}
		}
		//记录每个字符的index
		m[s[i]] = i
		//计算最大长度,注意i要+1才是实际序号
		maxLength = max(maxLength, i+1-startIndex)
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
