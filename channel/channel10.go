package main

import (
	"fmt"
	"time"
)

func test(a chan int) {
	i := <-a
	fmt.Println(i)
	a <- i + 1
}
func main() {
	a := make(chan int, 5) //修改2为1就报错，修改2为3可以正常运行
	c := make(chan int, 5) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5

	fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	c <- 6
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	a <- 10
	go test(a)
	go test(a)
	go test(a)
	go test(a)
	time.Sleep(1 * time.Second)
	fmt.Println(<-a)
}

//修改为1报如下的错误:
//fatal error: all goroutines are asleep - deadlock!
