package main

import "fmt"

func main() {
	slice := []byte {'a', 'b', 'c', 'd'}
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))


}
