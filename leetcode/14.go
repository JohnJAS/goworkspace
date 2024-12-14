package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	count := len(strs)
	var retStr string

	for i, v := range strs[0] {
		for j := 0; j < count; j++ {
			//don't forget the shortest string, maybe it is a blank string
			if len(strs[j]) == i {
				return retStr
			}
			if strs[0][i] != strs[j][i] {
				return retStr
			}
		}
		retStr += string(v)
	}

	return retStr
}

func main() {
	str := []string{"flower", "flow", ""}
	fmt.Println(longestCommonPrefix(str))
}
