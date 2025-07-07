package main

import (
	"fmt"
)

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	m := map[byte]struct{}{}

	maxLength := 0
	slow, fast := 0, 0
	for ; fast < len(s); fast++ {
		if _, ok := m[s[fast]]; ok {
			for slow <= fast && s[slow] != s[fast] {
				delete(m, s[slow])
				slow++
			}
			slow = slow + 1
		} else {
			m[s[fast]] = struct{}{}
			if maxLength < fast-slow {
				maxLength = fast - slow
			}
		}
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
