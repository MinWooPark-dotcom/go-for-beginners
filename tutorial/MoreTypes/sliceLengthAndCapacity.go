package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)       // len=0 cap=6 []
	fmt.Println(len(s)) // 0, 왜냐면 0부터 0전까지 슬라이싱하면 0이지
	fmt.Println(cap(s)) // 6, 그대롱지 왜냐면 헤드 위치가 그대로니까

	// Extend its length.
	s = s[:4]
	printSlice(s)       // len=4 cap=6 [2 3 5 7], 왜냐하면 헤드 위치가 2니까 2부터 4전까지 슬라이싱
	fmt.Println(len(s)) // 4
	fmt.Println(cap(s)) // 6, 헤드 위치 그대로

	// Drop its first two values.
	s = s[2:]
	printSlice(s)       // [5 7 11 13]이 아니라 위의 [2 3 5 7] 기준으로 슬라이싱을 하니까 [5 7]
	fmt.Println(len(s)) // 2
	fmt.Println(cap(s)) // 4, 헤드 위치 2로 가서 변경됨
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
