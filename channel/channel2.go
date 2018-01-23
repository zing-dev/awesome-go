package main

import "fmt"

func generateStrings(strings chan string) {
	strings <- "Monday"
	strings <- "Tuesday"
	strings <- "Wednesday"
	strings <- "Thursday"
	strings <- "Friday"
	strings <- "Saturday"
	strings <- "Sunday"
	close(strings)
}

func main() {
	strings := make(chan string) // 无缓冲channel
	go generateStrings(strings)

	for {
		//在读取的时候使用多重返回值来判断一个channel是否已经被关闭。
		if s, ok := <-strings; ok {
			fmt.Println(s)
		} else {
			fmt.Println("channel colsed.")
			break
		}
	}
}