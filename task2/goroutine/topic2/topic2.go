package main

import (
	"fmt"
	"time"
)

type Scheduler struct {
}

type Task func()

func (a *Scheduler) callAndRecordTime(tasks []Task) {
	for i, task := range tasks {
		go func() {
			started := time.Now()
			task()
			fmt.Printf("task%d耗时：%d微秒\n", i, time.Since(started).Microseconds())
		}()
	}
}

func main() {
	a := func() {
		sum := 0
		for i := 1; i < 1000000; i = i + 2 {
			sum += i
		}

		fmt.Printf("a 的和是 ： %d\n", sum)
	}

	b := func() {
		sum := 0
		for i := 0; i < 1000000; i++ {
			sum += i
		}
		fmt.Printf("b 的和是 ： %d\n", sum)
	}

	tasks := []Task{a, b}
	s := Scheduler{}
	s.callAndRecordTime(tasks)

	time.Sleep(time.Second)
}
