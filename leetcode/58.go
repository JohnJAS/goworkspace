package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	count := 0

	for index >= 0 {
		if s[index] == ' ' {
			index--
			continue
		}
		break
	}

	for index >= 0 && s[index] != ' ' {
		count++
		index--
	}

	return count
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}
