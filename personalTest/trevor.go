package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	ss := s
	p := make([]int, 20)
	printSlice(s)  // len=6 cap=6 [2 3 5 7 11 13]
	printSlice(ss) // len=6 cap=6 [2 3 5 7 11 13]
	printSlice(p)  // len=20 cap=20 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	// Slice the slice to give it zero length.
	s = s[:0]
	copy(p, s[:cap(s)]) // s[0:6]을 p에 카피
	printSlice(s)       // len=0 cap=6 []
	printSlice(ss)      // len=6 cap=6 [2 3 5 7 11 13]
	printSlice(p)       // len=20 cap=20 [2 3 5 7 11 13 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//! cap의 값이 있는건 메모리에 값이 있는 것임
	//! slice가 갖고 있는 건 배열 주소, len, cap, 시작점

	// Extend its length.
	s = s[:4]
	p = p[:4]
	printSlice(s)  // len=4 cap=6 [2 3 5 7]
	printSlice(ss) // len=6 cap=6 [2 3 5 7 11 13]
	printSlice(p)  // len=4 cap=20 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	p = p[2:]
	s[1] = 55
	printSlice(s)  // len=2 cap=4 [5 55]
	printSlice(ss) // len=6 cap=6 [2 3 5 55 11 13]
	printSlice(p)  // len=2 cap=18 [5 7]

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
