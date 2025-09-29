package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i < 10; i = i + 2 {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 2; i < 10; i = i + 2 {
			fmt.Println(i)
		}
	}()

	// 等待1秒， 否则主协程先于子协程结束
	time.Sleep(time.Second)

}
