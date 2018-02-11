package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = 10
	var pInt, pInt2 *int
	pInt = &c
	var arr = []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println("a address is ", &a)
	fmt.Println("b address is ", &b)
	fmt.Println("c address is ", &c)
	fmt.Println(arr)
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(*&arr[1]) //22
	var str = "golang is better than clang"
	fmt.Println(str)
	fmt.Println(str[0])         //103
	fmt.Println(string(str[0])) //g
	fmt.Println("pInt is ", pInt)
	fmt.Println("*pInt is ", *pInt) //10
	fmt.Println(pInt2)              //nil
	//fmt.Println(*pInt2)             //error
}
