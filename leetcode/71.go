package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	slice := strings.Split(path, "/")
	//fmt.Println(slice)

	var stack []string
	for _, s := range slice {
		if s == "" || s == "." {
			continue
		}
		if s == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, s)
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/home//"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/.../a/../b/c/../d/./"))
}
