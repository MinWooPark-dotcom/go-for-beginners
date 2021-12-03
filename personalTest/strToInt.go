package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 1. int to string - 숫자(정수)를 문자열로 변환
	a := strconv.Itoa(100)
	b := a + "abc"
	fmt.Println("b: ", b)                      // b: 100
	fmt.Println("type b: ", reflect.TypeOf(b)) // type a: string

}
