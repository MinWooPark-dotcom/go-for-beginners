package main

import "fmt"

// Closure는 함수 바깥에 있는 변수를 참조하는 함수값(function value)를 일컫는데, 이때의 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 된다.

// adder함수는 받는 인자가 없는데 어떻게 22, 23에서 인자를 받지?
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		// 내가 궁금한 건 어떻게 x를 받지? 바깥 함수가 받는 파라미터 없으면 자동으로 내부 함수가 파라미터를 받나?
		sum += x // 여기서 이 익명함수가 그 함수 바깥에 있는 변수 sum 를 참조하고 있는 것은 이해함.
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() //! adder가 아니라 adder() 함수 실행을 할당 즉, 리턴 값을 할당하니까 10라인의 익명함수를 할당해서 인자를 받을 수 있는 것임
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
