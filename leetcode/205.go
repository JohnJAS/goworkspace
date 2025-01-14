package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)

	for i := range s {
		a, b := s[i], t[i]
		if s2t[a] != 0 && s2t[a] != b || t2s[b] != 0 && t2s[b] != a {
			return false
		}

		s2t[a] = b
		t2s[b] = a

	}

	return true
}

func main() {

	fmt.Println(isIsomorphic("foo", "bar"))
}
