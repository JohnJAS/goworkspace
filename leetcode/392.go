package main

import "fmt"

func isSubsequence(s string, t string) bool {
	//p1 for s, p2 for t
	//cal is s is sub string of t
	p1, p2 := 0, 0
	if len(s) == 0 {
		return true
	}
	for ; p2 < len(t); p2++ {
		if s[p1] == t[p2] {
			p1++
		}
		if p1 == len(s) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}
