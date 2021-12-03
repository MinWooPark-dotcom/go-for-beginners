package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

//문자열 s 의 각 word 의 개수의 맵을 반환해야 합니다
func WordCount(s string) map[string]int {
	// 0. map을 만듦
	// 1. s를 띄어쓰기 기준으로 스플릿 할 수 있어야 함.
	// 2. 스플릿한 값을 키로 두고 그 값이 s에서 몇 번 나왔는지 반복문을 통해 체크
	// 3. 단어 개수를 해당 단어 키의 값에 넣음

	var m map[string]int
	m = make(map[string]int) // 주어진 타입 초기화, 사용 준비된 맵 반환

	splitStr := strings.Fields(s)

	for i := 0; i < len(splitStr); i++ {
		_, ok := m[splitStr[i]]
		// fmt.Println(ok)
		if ok == true {
			// 이미 존재하면 해당 키의 값에 +1
			//m[splitStr] = m[splitStr] + 1
			m[splitStr[i]]++
		} else {
			// 만약에 m에 splitStr[i]라는 키가 없으면 새로 만듦
			m[splitStr[i]] = 1
		}
	}
	fmt.Println(m)
	return m
}

func main() {
	wc.Test(WordCount)
}
