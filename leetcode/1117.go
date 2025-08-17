package main

import (
	"fmt"
	"sync"
)

type H2O struct {
	hQueue chan chan struct{} // Hydrogen 请求 channel
	oQueue chan chan struct{} // Oxygen 请求 channel
}

func NewH2O() *H2O {
	h2o := &H2O{
		hQueue: make(chan chan struct{}, 100),
		oQueue: make(chan chan struct{}, 100),
	}

	// 组装 goroutine
	go func() {
		for {
			// 等待 2 个 H 和 1 个 O
			h1 := <-h2o.hQueue
			h2 := <-h2o.hQueue
			o := <-h2o.oQueue

			// 同时释放信号
			h1 <- struct{}{}
			h2 <- struct{}{}
			o <- struct{}{}
		}
	}()

	return h2o
}

func (h2o *H2O) Hydrogen(releaseHydrogen func()) {
	signal := make(chan struct{})
	h2o.hQueue <- signal // 请求加入
	<-signal             // 等待组装 goroutine释放
	releaseHydrogen()    // 打印 H
}

func (h2o *H2O) Oxygen(releaseOxygen func()) {
	signal := make(chan struct{})
	h2o.oQueue <- signal // 请求加入
	<-signal             // 等待组装 goroutine释放
	releaseOxygen()      // 打印 O
}

func main() {
	h2o := NewH2O()
	output := make(chan string, 100)
	var wg sync.WaitGroup

	releaseH := func() { output <- "H" }
	releaseO := func() { output <- "O" }

	input := "HOHOOOOOOOOHHHHOOHHHHHHHHHHOHHHHHHHHOHH"

	for _, c := range input {
		wg.Add(1)
		switch c {
		case 'H':
			go func() {
				defer wg.Done()
				h2o.Hydrogen(releaseH)
			}()
		case 'O':
			go func() {
				defer wg.Done()
				h2o.Oxygen(releaseO)
			}()
		}
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	for v := range output {
		fmt.Print(v)
	}
	fmt.Println()
}
