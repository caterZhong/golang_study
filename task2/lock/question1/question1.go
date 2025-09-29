package main

import (
	"fmt"
	"sync"
	"time"
)

var count int = 0

var mutex sync.Mutex

func plusOne() {
	mutex.Lock()
	defer mutex.Unlock()
	for i := 0; i < 1000; i++ {
		count = count + 1
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go plusOne()
	}

	time.Sleep(time.Second)
	fmt.Println("启动10个协程， 每个协程加1000，结果是：", count)
}
