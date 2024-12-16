package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	var cnt [26]int
	for _, ch := range magazine {
		cnt[ch-'a']++
	}

	for _, ch := range ransomNote {
		if cnt[ch-'a'] > 0 {
			cnt[ch-'a']--
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("aa", "ab"))
}
