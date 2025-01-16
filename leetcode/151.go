package main

import "fmt"

func reverseWords(s string) string {
	//transfer to list
	var list []string
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			continue
		} else {
			t := ""
			for i < len(s) && string(s[i]) != " " {
				t = t + string(s[i])
				i++
			}
			list = append(list, t)
		}
	}

	//reverse list
	var result string
	for i := len(list) - 1; i >= 0; i-- {
		if result != "" {
			result = result + " " + list[i]
		} else {
			result = result + list[i]
		}
	}
	return result
}

func main() {
	fmt.Println(reverseWords("  ab aba   aba ab  ab  "))
}
