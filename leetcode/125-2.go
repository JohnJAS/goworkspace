package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left <= right {
		for left <= right && !isChar(s[left]) {
			left++
		}
		for left <= right && !isChar(s[right]) {
			right--
		}
		//注意这里的条件判断
		if left <= right && s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isChar(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') // 如果只考虑字母可以去掉后半部分
}

func main() {
	fmt.Println(isPalindrome("0P"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome(" "))

}
