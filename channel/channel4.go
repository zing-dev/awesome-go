package main

import "fmt"

func MyRoutineFunc(ch chan int) {
	// 函数处理
	ch <- 1

	fmt.Println("MyRoutineFunc process finished.")
}
//主函数等待所有goroutine完成后返回
//我们已经知道golang程序从main()函数开始执行，当main()函数返回时，程序结束且不等待其他goroutine结束。
//如果main函数使用time.Sleep方式阻塞等待所有goroutine返回，那么这个休眠时间势必无法控制精确。通过使用channel可以很好解决这个问题。
func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go MyRoutineFunc(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}

	fmt.Println("All goroutine finished.")
}
