package main

import (
	"time"
	"fmt"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		fmt.Println("------------------")
		for {
			select {
			case v := <-c:
				println(v)
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			//default:
			//	o <- true
			//	fmt.Println("++++++++++++++++++++")
			}
		}
	}()
	println(<-o)
}
