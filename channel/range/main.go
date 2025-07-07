package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	
	fmt.Print("Range over channel: ")
	for n := range numbers {
		fmt.Print(n, " ")
	}
	fmt.Println()
}
