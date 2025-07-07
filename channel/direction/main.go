package main

import (
	"fmt"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)
	
	go pingPong(ping, pong)
	ping <- "ping"
	fmt.Println("Directional channel:", <-pong)
}

func pingPong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg + " pong"
}
