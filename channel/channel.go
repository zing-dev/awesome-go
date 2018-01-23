package main

import "fmt"


//不要通过共享内存来通信，而应该通过通信来共享内存。
//golang提供一种基于消息机制而非共享内存的通信模型。消息机制认为每个并发单元都是自包含的独立个体，并且拥有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。
//channel是golang在语言级提供的goroutine间的通信方式，可以使用channel在两个或多个goroutine之间传递消息。
//channel是进程内的通信方式，如果需要跨进程通信，建议使用分布式的方法来解决，比如使用Socket或HTTP等通信协议。
//channel是类型相关的，即一个channel只能传递一种类型的值，需要在声明channel时指定。可以认为channel是一种类型安全的管道。

func generateString(strings chan string) {
	strings <- "Monday"
	strings <- "Tuesday"
	strings <- "Wednesday"
	strings <- "Thursday"
	strings <- "Friday"
	strings <- "Saturday"
	strings <- "Sunday"
	close(strings)
}

//channel读写语法
//√ 向无缓冲的channel写入数据会导致该goroutine阻塞，直到其他goroutine从这个channel中读取数据。
//√ 向带缓冲的且缓冲已满的channel写入数据会导致该goroutine阻塞，直到其他goroutine从这个channel中读取数据。
//√ 向带缓冲的且缓冲未满的channel写入数据不会导致该goroutine阻塞。
//√ 从无缓冲的channel读出数据，如果channel中无数据，会导致该goroutine阻塞，直到其他goroutine向这个channel中写入数据。
//√ 从带缓冲的channel读出数据，如果channel中无数据，会导致该goroutine阻塞，直到其他goroutine向这个channel中写入数据。
//√ 从带缓冲的channel读出数据，如果channel中有数据，该goroutine不会阻塞。
//√ 总结：无缓冲的channel读写通常都会发生阻塞，带缓冲的channel在channel满时写数据阻塞，在channel空时读数据阻塞。
func main() {
	strings := make(chan string) // 无缓冲channel
	go generateString(strings)

	for s := range strings {
		fmt.Println(s)
	}
}
