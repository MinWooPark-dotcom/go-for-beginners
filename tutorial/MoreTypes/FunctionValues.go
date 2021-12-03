package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 5*5 + 12*12 = 169의 제곱근 => 13

	fmt.Println(compute(hypot)) // 3*3 + 4*4 = 25의 제곱근 => 5
	// fmt.Println(math.Pow) 0x45e3a0로 나오는데? 뭐지
	fmt.Println(compute(math.Pow)) // math.Pow의 인자가 없는데?
}

/*
13
5
0x45e3a0
81
*/
