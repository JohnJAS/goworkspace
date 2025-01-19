package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	p2w := make(map[byte]string)
	w2p := make(map[string]byte)
	i := 0
	plen := len(pattern)
	if plen != len(words) {
		return false
	}
	for _, word := range words {
		if v, ok := p2w[pattern[i]]; ok && v != word {
			return false
		}
		if v, ok := w2p[word]; ok && v != pattern[i] {
			return false
		}
		p2w[pattern[i]] = word
		w2p[word] = pattern[i]
		i++
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
