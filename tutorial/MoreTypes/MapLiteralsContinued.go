package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},  // value부분에 Vertex를 생략
	"Google":    {37.42202, -122.08408}, // 위와 같음
}

func main() {
	fmt.Println(m)
}

/*
Map literals(20/27)이랑 차이나는 부분은 key-value에서 value부분에 Vertex를 생략했음
최상위 유형이 유형 이름인 경우(여기서는 Vertex) 생략 가능, 그럼 Vertex부분이 Vertex가 아닐 수 있나?
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
*/
