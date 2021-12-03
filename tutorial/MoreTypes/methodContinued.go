package main

import (
	"fmt"
	"math"
)

// You can declare a method on non-struct types, too.
// 비구조체(MyFloat)의 타입(float)과 메서드의 타입(f가 MyFloat타입이니 즉 float64) 둘 다 MyFloat여서 가능?
// int나 다른 데이터 타입은 안되긴 함
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)       // -1.4142135623730951
	fmt.Println(f.Abs()) // 1.4142135623730951

}
