package main

import (
	"fmt"
	"unicode"
)

func detectCapitalUse(word string) bool {
	t := []rune(word)
	sum := 0
	for _, c := range t {
		if unicode.IsUpper(c) {
			sum++
		}
	}

	return sum == 0 || (sum == 1 && t[0] >= 'A' && t[0] <= 'Z') || sum == len(t)
}

func main() {
	fmt.Println(detectCapitalUse("GGoogle"))
}
