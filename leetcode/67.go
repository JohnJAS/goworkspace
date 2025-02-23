package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	lena := len(a)
	lenb := len(b)
	prefix := ""
	if lena > lenb {
		for i := 0; i < lena-lenb; i++ {
			prefix = prefix + "0"
		}
		b = prefix + b
	} else if lena < lenb {
		for i := 0; i < lenb-lena; i++ {
			prefix = prefix + "0"

		}
		a = prefix + a
	}

	carry := 0
	result := ""
	for i := len(a) - 1; i >= 0; i-- {
		p, _ := strconv.Atoi(string(a[i]))
		q, _ := strconv.Atoi(string(b[i]))

		n := (p + q + carry) % 2
		carry = (p + q + carry) / 2

		result = strconv.Itoa(n) + result
	}
	if carry == 1 {
		result = "1" + result
	}

	return result
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
