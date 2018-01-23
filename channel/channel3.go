package main

import "fmt"

//单向channel语法
//▶  使用意义
//√ golang中假如一个channel只允许读，那么channel肯定只会是空的，因为没机会往里面写数据。
//√ golang中假如一个channel只允许写，那么channel最后只会是满的，因为没机会从里面读数据。
//√ 单向channel概念，其实只是对channel的一种使用限制，即在将一个channel变量传递到一个函数时，
// 可以通过将其指定为单向channel变量，从而限制该函数中可以对此channel的操作，达到权限控制作用。

//声明语法
//var ch1 chan elementType   // ch1是一个正常的channel
//var ch2 chan<- elementType // ch2是单向channel，只用于写数据
//var ch3 <-chan elementType // ch3是单向channel，只用于读数据
func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value", value)
	}
}
//类型转换
//ch1 := make(chan elementType)
//ch2 := <-chan elementType(ch1) // ch2是一个单向的读取channel
//ch3 := chan<- elementType(ch1) // ch3是一个单向的写入channel
func main() {

	var ch chan int
	ch = make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		ch <- 6
		ch <- 7
		close(ch)
	}()
	Parse(ch)
}
