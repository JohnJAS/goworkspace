package main

import (
	"fmt"
	"sync"
)

type H2O struct {
	hydrogenSem chan struct{} // 用于氢原子的信号量
	oxygenSem   chan struct{} // 用于氧原子的信号量
	wg          sync.WaitGroup
}

func NewH2O() *H2O {
	return &H2O{
		hydrogenSem: make(chan struct{}, 2),
		oxygenSem:   make(chan struct{}, 1),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.hydrogenSem <- struct{}{} // 获取氢原子位置

	// 等待氧原子就位
	<-h2o.oxygenSem

	releaseHydrogen()

	h2o.wg.Done()
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.oxygenSem <- struct{}{} // 获取氧原子位置

	// 等待两个氢原子就位
	<-h2o.hydrogenSem
	<-h2o.hydrogenSem

	releaseOxygen()

	// 重置信号量，准备下一组
	h2o.oxygenSem <- struct{}{}
	h2o.wg.Done()
}

func main() {
	water := "HHHHHHOOO"
	h2o := NewH2O()

	var result []byte
	var mu sync.Mutex

	for _, c := range water {
		h2o.wg.Add(1)
		switch c {
		//一直是H的话，就会导致启动很多H进程，争夺O
		case 'H':
			go h2o.hydrogen(func() {
				mu.Lock()
				result = append(result, 'H')
				mu.Unlock()
			})
		case 'O':
			go h2o.oxygen(func() {
				mu.Lock()
				result = append(result, 'O')
				mu.Unlock()
			})
		}
	}

	h2o.wg.Wait()
	fmt.Println(string(result))
}
