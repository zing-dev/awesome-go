package main

import (
	"fmt"
	"time"
)

//select操作
//√ golang中的select关键字用于处理异步IO，可以与channel配合使用。
//√ golang中的select的用法与switch语言非常类似，不同的是select每个case语句里必须是一个IO操作。
//√ select会一直等待等到某个case语句完成才结束。

func main() {
	timeout := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second) // sleep 3 seconds
		timeout <- true
	}()

	// 实现了对ch读取操作的超时设置。
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout!")
	}
}
