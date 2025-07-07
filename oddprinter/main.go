package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	oddChan := make(chan struct{})
	evenChan := make(chan struct{})

	// 打印奇数的 goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 99; i += 2 {
			select {
			case <-oddChan: // 等待奇数信号
				fmt.Println(i)
				evenChan <- struct{}{} // 通知偶数开始
			}

		}
	}()

	// 打印偶数的 goroutine
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			select {
			case <-evenChan: // 等待偶数信号
				fmt.Println(i)
				if i != 100 { // 最后一次不再发送信号
					oddChan <- struct{}{}
				}
			}

		}
	}()

	// 首先启动奇数打印
	oddChan <- struct{}{}

	wg.Wait()
}
