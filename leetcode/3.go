package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	maxLength := 0
	startIndex := 0
	//<character:index>
	m := make(map[byte]int)

	for i := 0; i < n; i++ {
		//检查是否有重复
		if _, ok := m[s[i]]; ok {
			//删除到重复字符前的所有记录的字符
			for startIndex <= m[s[i]] {
				delete(m, s[startIndex])
				startIndex++
			}
		}
		//记录当前字符
		m[s[i]] = i
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
