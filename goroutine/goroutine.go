package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func numbers() {
	for i := 1; i <= 9; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'i'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {
	//当创建一个Go协程时，创建这个Go协程的语句立即返回。与函数不同，程序流程不会等待Go协程结束再继续执行。
	// 程序流程在开启Go协程后立即返回并开始执行下一行代码，忽略Go协程的任何返回值。
	//在主协程存在时才能运行其他协程，主协程终止则程序终止，其他协程也将终止。

	//√ 在一个函数前加上go关键字调用，这次调用就会在一个新的goroutine中并发执行，开启goroutine的线程将继续执行。
	//√ 当被go调用的函数返回时，这个goroutine也自动结束了。如果这个函数有返回值，那么这个返回值会被丢弃。
	//√ golang程序从main()函数开始执行，当main()函数返回时，程序结束且不等待其他goroutine结束。
	fmt.Println("-----------------------------------------------------")
	go hello()
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main function")
	fmt.Println("-----------------------------------------------------")
}
