package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	sign := make(chan int, 2)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()

	go func() {
		for {
			e, ok := <-ch
			fmt.Printf("%d (%v)\n", e, ok)
			if !ok {
				break
			}
			time.Sleep(2 * time.Second)
		}
		fmt.Println("Done.")
		sign <- 1
	}()
	fmt.Println(<-sign)
	fmt.Println(<-sign)
	close(sign)
}
