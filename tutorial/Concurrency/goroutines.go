package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)
}

/*
Go루틴(goroutine)은 Go 런타임이 관리하는 Lightweight 논리적 (혹은 가상적) 쓰레드(주1)이다.
Go에서 "go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행한다.
goroutine은 비동기적으로(asynchronously) 함수루틴을 실행하므로, 여러 코드를 동시에(Concurrently) 실행하는데 사용된다.

실행 결과
Sync *** 0
Sync *** 1
Sync *** 2
goroutines는 실행 순서가 일정하지 않아서 프로그램 실행 시 마다 다른 출력 결과가 나올 수 있음
Async3 *** 0
Async3 *** 1
Async3 *** 2
Async1 *** 0
Async2 *** 0
Async2 *** 1
Async2 *** 2
Async1 *** 1
Async1 *** 2
*/
