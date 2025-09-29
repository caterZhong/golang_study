package main

import (
	"fmt"
	"time"
)

// 发送内容到channel
func send2Chan(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// 发送内容到channel
func receiveFromChan(ch <-chan int) {
	for v := range ch {
		fmt.Println("从channel 接收到", v)
	}
}

func main() {

	ch := make(chan int, 3)
	go send2Chan(ch)
	// time.Sleep(time.Second)
	go receiveFromChan(ch)

	// 使用select进行多路复用
	timeout := time.After(100 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
