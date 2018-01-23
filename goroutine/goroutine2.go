package main

import (
	"fmt"
	"runtime"
	"time"
)

func loop() {
	fmt.Println("loop starts!")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func main() {

	//并发与并行
	//√ 两个队列，一个Coffee机器，那是并发。
	//√ 两个队列，两个Coffee机器，那是并行。

	//并发执行
	//√ 如果在单核cpu情况下，golang所有的goroutine只能在一个线程里跑 。
	//√ 如果当前goroutine不发生阻塞，它是不会让出cpu时间给其他goroutine，除非调用runtime.Gosched()主动让出时间片。
	//√ 如果当前goroutine发生阻塞，它会主动让出cpu时间给其他goroutine执行。
	//√ golang的runtime包是goroutine的调度器，其中使用runtime.GOMAXPROCS(n)可以控制使用cpu核数。

	runtime.GOMAXPROCS(10) // 强制使用1个cpu

	go loop()
	go loop()
	go loop()

	time.Sleep(3 * time.Second)
}
