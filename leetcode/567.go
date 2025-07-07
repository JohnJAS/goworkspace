package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {

	var result string
	mt := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}

	ms := make(map[byte]int)

	verify := func() bool {
		for k := range mt {
			if ms[k] < mt[k] {
				return false
			}
		}
		return true
	}

	left, right := 0, 0
	minLen := math.MaxInt
	for right < len(s) {
		ms[s[right]]++
		//检查是否覆盖
		for verify() && left <= right {
			curLen := right - left + 1
			if curLen < minLen {
				minLen = curLen
				result = s[left : right+1]
			}
			if _, ok := mt[s[left]]; ok {
				ms[s[left]]--
			}
			left++
		}

		right++
	}

	return result
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
