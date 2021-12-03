package main

import "fmt"

func main() {
	a := make([]int, 5) // make 함수는 0으로 이루어진 배열을 할당, 그리고 해당 배열을 참조하는 슬라이스를 반환
	printSlice("a", a)  // len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5) // 용량을 지정하려면 세 번쨰 인자에 값 전달
	printSlice("b", b)     // len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) // len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) // len=3 cap=3 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]
*/
