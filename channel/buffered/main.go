package main

import (
	"fmt"
)

func main() {
	bufferedCh := make(chan int, 2)
	bufferedCh <- 1
	bufferedCh <- 2
	fmt.Println("Buffered channel:", <-bufferedCh, <-bufferedCh)
}
