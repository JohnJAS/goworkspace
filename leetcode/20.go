package main

import (
	"fmt"
)

func isValid(s string) bool {
	m := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	var stack []rune
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			//push
			stack = append(stack, c)
		}
		if c == '}' || c == ']' || c == ')' {
			//pop
			if len(stack) == 0 || m[c] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println('{')
	test := []int{1, 2, 3}
	fmt.Println(test[:2])
	fmt.Println(isValid("{{{{123}}}}"))
	fmt.Println(isValid("()"))
}
