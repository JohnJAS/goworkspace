package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	t := []rune(s)
	n := len(s)
	i, j := 0, n-1

	for i < j {
		for i < n && !strings.Contains("AEIOUaeiou", string(s[i])) {
			i++
		}
		for j > 0 && !strings.Contains("AEIOUaeiou", string(s[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}

	}
	return string(t)
}

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}
