package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var count atomic.Int64

func plusOne() {
	for i := 0; i < 1000; i++ {
		count.Add(1)
	}
}

func main() {
	count.Store(0)
	for i := 0; i < 10; i++ {
		go plusOne()
	}

	time.Sleep(time.Second)
	fmt.Println("启动10个协程， 每个协程加1000，结果是：", count.Load())
}
