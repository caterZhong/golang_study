package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	count := 100
	for i := 0; i < count; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println("从channel中接收到：", v)
	}
}

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch)

	time.Sleep(3 * time.Second)

}
