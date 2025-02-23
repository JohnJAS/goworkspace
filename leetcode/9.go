package main

import "fmt"

func isPalindrome(x int) bool {
	//x为复数，或者10的倍数
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	if x == 0 {
		return true
	}

	var turnX int
	for x != 0 {
		x, turnX = x/10, turnX*10+x%10
		fmt.Printf("x = %d, turnX = %d\n", x, turnX)
		if turnX == x || turnX/10 == x {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(1223221))
	fmt.Println(isPalindrome(12233221))
	fmt.Println(isPalindrome(100))
}
