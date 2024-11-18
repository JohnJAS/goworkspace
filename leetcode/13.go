package main

import "fmt"

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	var ans int
	length := len(s)
	for i := 0; i <= length-1; i++ {
		v := symbolValues[s[i]]
		if i < length-1 && v < symbolValues[s[i+1]] {
			ans = ans - v
		} else {
			ans = ans + v
		}
	}
	return ans
}

func main() {
	fmt.Println(romanToInt("IXX"))
}
