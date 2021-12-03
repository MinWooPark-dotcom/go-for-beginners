package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// fmt.Println(pow) [1 2 4 8 16 32 64 128]
	// for문 정석은 for i := 0; i < 10; i++ {
	//	...
	// }
	//

	// The range form of the for loop iterates over a slice or map.
	// i는 값 할당 안하면 제로 벨류 Variables declared without an explicit initial value are given their zero value.
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
2**4 = 16
2**5 = 32
2**6 = 64
2**7 = 128
*/
