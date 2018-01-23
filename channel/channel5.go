package main

import (
	"fmt"
	"time"
)

func main() {
	// 实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3 * time.Second) // 等待3秒钟
		timeout <- true
	}()

	// 然后结合使用select实现超时机制
	ch := make(chan int)
	select {
	case <-ch:
		// 从ch中读取到数据
	case <-timeout:
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
		fmt.Println("timeout!")
	}
}
