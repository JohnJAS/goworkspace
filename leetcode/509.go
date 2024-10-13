package main

import "fmt"

// 动态规划
func fib(n int) int {
	if n < 2 {
		return n
	}
	//f(2) = f(1) + f(0)
	y, x := 1, 0
	var result int
	//f(3) = f(2) + f(1) ， 其中f(2)和f(1)都已经计算过了，无需再算
	//同理， f(n) = f(n-1) + f(n-2) , 后面两个再上一次计算中都算过了，可以直接拿来用
	for i := 2; i <= n; i++ {
		result = y + x
		x = y
		y = result
	}
	return result
}

//递归
//func fib(n int) int {
//	if n == 1 {
//		return 1
//	}
//	if n == 0 {
//		return 0
//	}
//	return fib(n-1) + fib(n-2)
//}

func main() {
	fmt.Println(fib(4))
}
