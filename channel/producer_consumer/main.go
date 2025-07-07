package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producing:", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // 生产者负责关闭channel
	fmt.Println("Producer finished")
}

func consumer(id int, ch <-chan int, done chan<- bool) {
	for num := range ch {
		fmt.Printf("Consumer %d received: %d\n", id, num)
		time.Sleep(1 * time.Second)
	}
	done <- true
}

func main() {
	// 创建channel
	dataCh := make(chan int)
	doneCh := make(chan bool)

	// 启动生产者
	go producer(dataCh)

	// 启动多个消费者
	for i := 1; i <= 2; i++ {
		go consumer(i, dataCh, doneCh)
	}

	// 等待所有消费者完成
	for i := 1; i <= 2; i++ {
		<-doneCh
	}
	fmt.Println("All consumers finished")
}
