package main

import "fmt"

func strictPalindrome(s string, l, r int) bool {
	for ; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return strictPalindrome(s, l, r-1) || strictPalindrome(s, l+1, r)
		}
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abca"))
}
