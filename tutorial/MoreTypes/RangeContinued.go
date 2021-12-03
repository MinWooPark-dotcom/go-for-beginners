package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// fmt.Println(pow) // [0 0 0 0 0 0 0 0 0 0]
	// testUint := 1 << uint(3)
	// fmt.Println(testUint)

	fmt.Println(pow) // [0 0 0 0 0 0 0 0 0 0]
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		// uint()는 int형으로 데이터 타입 변환하고 1비트를 unit(i)만큼 비트 이동한 값이 2**i와 같은 결과
	}
	fmt.Println(pow) // [1 2 4 8 16 32 64 128 256 512]
	// '_'이거로 인덱스 또는 값 생략 가능
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
