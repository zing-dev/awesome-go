package main

import (
	"fmt"
	"math"
)

func main() {
	//float32的有效bit位只有23,其它的bit位用于指数和符号
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	fmt.Println(f)
	fmt.Println(f + 1)
	const (
		e        = 2.71828        // (approximately)
		Avogadro = 6.02214129e23  // 阿伏伽德罗常数
		Planck   = 6.62606957e-34 // 普朗克常数
	)
	fmt.Println(e, Avogadro, Planck)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
	fmt.Println(nan == math.NaN())

}
