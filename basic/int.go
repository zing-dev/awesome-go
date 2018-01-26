package main

import (
	"fmt"
	"math"
)

func main() {
	//uint8 0-255
	var u uint8 = 255
	var i int8 = 127
	//255 0 1
	fmt.Println(u, u+1, u*u)
	fmt.Println(i, i+1, i*i)

	fmt.Println("math.MinInt8 ", math.MinInt8)
	fmt.Println("math.MaxInt8 ", math.MaxInt8)
	fmt.Println(1<<7 - 1)
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}
