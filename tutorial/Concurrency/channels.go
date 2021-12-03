package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	// Range continued: _ 을 할당하여 인덱스 또는 값을 건너뛸 수 있습니다.
	for _, v := range s {
		//fmt.Println(v)
		sum += v
	}
	fmt.Println(sum) // -5(-9+4+0), 17(7+2+8)
	c <- sum         // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // s[:3], c
	go sum(s[len(s)/2:], c) // s[3:], c
	// 아래가 의문인게 같은 c인데 어떻게 -5, 17 값을 구분하지?
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/*
-5
17
-5 17 12
*/
