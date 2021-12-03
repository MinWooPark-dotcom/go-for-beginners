package main

import "fmt"

func main() {
	slice := []int{0, 10, 20, 30}

	// 새로운 슬라이스를 생성
	copyslice := make([]int, 10)
	fmt.Println(copyslice) // [0 0 0 0 0 0 0 0 0 0]

	// copy 함수로 슬라이스 복사
	copy(copyslice, slice)

	fmt.Println(copyslice)      // [0 10 20 30 0 0 0 0 0 0]
	fmt.Println(len(slice))     // 4
	fmt.Println(cap(slice))     // 4
	fmt.Println(len(copyslice)) // 10
	fmt.Println(cap(copyslice)) // 10
	// 값 변경
	copyslice[0] = 100

	fmt.Println(copyslice) // [100 10 20 30 0 0 0 0 0 0]
	fmt.Println(slice)     // [0 10 20 30]
}
