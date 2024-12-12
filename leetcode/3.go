package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	//record character maps
	m := map[byte]int{}
	//define the right pointer(notice from -1)
	rk := -1
	//define the length of longest substring
	maxLength := 0
	//go through each character in the string
	for i := 0; i < len(s); i++ {
		//delete the previous character from map
		if i > 0 {
			delete(m, s[i-1])
		}
		//increase right pointer to the rightest place
		for rk+1 < len(s) && m[s[rk+1]] == 0 {
			rk = rk + 1
			m[s[rk]] = 1
		}
		maxLength = max(maxLength, rk-i+1)
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
