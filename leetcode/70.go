package main

import "fmt"

// f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs2(n int) int {

	if n < 3 {
		return n
	}
	prePre, pre := 1, 2
	var result int
	for i := 3; i <= n; i++ {
		result = pre + prePre
		prePre = pre
		pre = result
	}
	return result
}

func main() {
	fmt.Println(climbStairs(5))
}
