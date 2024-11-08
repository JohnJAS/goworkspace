package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for !isNum(s[l]) && l < r {
			l++
		}
		for !isNum(s[r]) && l < r {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}

	}
	return true
}

func isNum(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
