package main

import "fmt"

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
	}

	for _, c := range t {
		m[c]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
