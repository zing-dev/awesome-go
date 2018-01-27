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


	f := 3.141 // a float64
	fi := int(f)
	fmt.Println(f, fi) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"
	fmt.Println(0x1111111)
	fmt.Println(0xf1111111)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
	fmt.Printf("%c",22269)
}
