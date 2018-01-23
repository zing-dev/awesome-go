package main

import (
	"fmt"
)

func init() {
	fmt.Println("before main function")
}

func main() {
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	)

	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

	const (
		h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
	)
	const (
		a       = iota //a=0
		b       = "B"
		c       = iota             //c=2
		d, e, f = iota, iota, iota //d=3,e=3,f=3
		g       = iota             //g = 4
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)

	fmt.Println("Hello World")
	sum := Add(2, 4)
	fmt.Println(sum)
	fmt.Println("----------------------------------------------------------")

	sum = 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println(sum)

	for sum < 100 {
		sum *= sum
	}
	fmt.Println(sum)

	s, p := SumAndProduct(10, 20)
	fmt.Println(s, p)

}

func Add(a, b int) int {
	return a + b
}
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}
