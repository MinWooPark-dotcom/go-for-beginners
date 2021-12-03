/*
Map은 키(Key)에 대응하는 값(Value)을 신속히 찾는 해시테이블(Hash table)을 구현한 자료구조이다. Go 언어는 Map 타입을 내장하고 있는데, "map[Key타입]Value타입" 과 같이 선언할 수 있다.
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 변수명 map[Key타입]Value타입
// 이때 선언된 변수 m은 (map은 reference 타입이므로) nil 값을 갖으며, 이를 Nil Map이라 부른다.
// Nil map에는 어떤 데이타를 쓸 수 없는데, map을 초기화하기 위해 make()함수를 사용할 수 있다 (map 리터럴을 사용할 수도 있다, 리터럴 초기화는 "map[Key타입]Value타입 { key:value }" 와 같이 Map 타입 뒤 { } 괄호 안에 "키: 값" 들을 열거하면 된다.)
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) // 주어진 타입 초기화, 사용 준비된 맵 반환
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
