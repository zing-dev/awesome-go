package main

import (
	"fmt"
)

type Ban struct {
	ips map[string]bool
}

func TestBan() {
	ban := Ban{}
	var count int
	var index = make(chan int, 100)
	var ipChan = make(chan string, 1)
	var en = make(chan map[string]bool, 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			index <- j
			go func() {
				v := <-index
				ip := fmt.Sprintf("192.111.1.%v", v)
				ipChan <- ip
				fmt.Println("count > ", ip)

				if _, ok := ban.ips[ip]; !ok {
					count++
				}
				ban.ips[<-ipChan] = true
			}()
		}
	}
	fmt.Println("count > ", count)
}

func main() {
	TestBan()
}
