package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	var r int

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				r = a + b
			case "-":
				r = a - b
			case "*":
				r = a * b
			case "/":
				r = a / b
			}
		} else {
			r, _ = strconv.Atoi(token)
		}
		stack = append(stack, r)
	}
	return r
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
