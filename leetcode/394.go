package main

import "fmt"

func decodeString(s string) string {
	var countStack []int
	var stringStack []string
	var currentString string
	k := 0

	for _, c := range s {
		if c >= '0' && c <= '9' {
			//处理多位数数字
			k = k*10 + int(c-'0')
		} else if c == '[' {
			//压栈
			countStack = append(countStack, k)
			k = 0
			stringStack = append(stringStack, currentString)
			currentString = ""
		} else if c == ']' {
			// 弹出栈并拼接重复字符串
			repeatTimes := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			prevString := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]

			// 重复添加栈顶的curString
			for i := 0; i < repeatTimes; i++ {
				prevString = prevString + currentString
			}
			currentString = prevString

		} else {
			//如果是字母，直接加到结果字符串上
			currentString = currentString + string(c)
		}
	}
	return currentString
}

func main() {
	s1 := "a2[c]"
	s2 := "3[a2[c]]"
	s3 := "2[abc]3[cd]ef"

	fmt.Println(decodeString(s1))
	fmt.Println(decodeString(s2))
	fmt.Println(decodeString(s3))
}
