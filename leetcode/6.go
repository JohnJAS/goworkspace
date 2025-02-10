package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	strs := make([]string, numRows)
	i, flag := 0, -1
	for _, v := range s {
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		strs[i] = strs[i] + string(v)
		i = i + flag
	}
	return strings.Join(strs, "")
}

func main() {
	fmt.Println(convert("AB", 1))
}
