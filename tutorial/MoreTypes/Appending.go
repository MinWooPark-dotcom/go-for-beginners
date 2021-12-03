package main

import "fmt"

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)    // 2
	n := m + len(data) // 2 + 3 = 5
	// cap: 2
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2) // (5+1)*2 = 12, newSlice는 len = 12, cap = 12
		fmt.Println(newSlice)             // [0 0 0 0 0 0 0 0 0 0 0 0]
		copy(newSlice, slice)             // copy해도 newSlice의 len, cap은 변하지 않음
		fmt.Println(len(newSlice))        // 12
		slice = newSlice
		fmt.Println(slice) // [0 1 0 0 0 0 0 0 0 0 0 0]
	}
	slice = slice[0:n]
	fmt.Println(slice)      // [0 1 0 0 0]
	fmt.Println(len(slice)) // 5
	fmt.Println(cap(slice)) // 12
	copy(slice[m:n], data)
	fmt.Println(slice)      // [0 1 2 3 4]
	fmt.Println(len(slice)) // 5
	fmt.Println(cap(slice)) // 12
	return slice
}

func printSlice(s []byte) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []byte{0, 1}

	fmt.Println(s) // [0 1]
	AppendByte(s, 2, 3, 4)
}
