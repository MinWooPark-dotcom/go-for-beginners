package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m) // map[]

	m["Answer"] = 42
	fmt.Println(m) // map[Answer:42]
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	// fmt.Println(m) map[]

	/* v 또는 ok가 선언된 경우
	v := 0
	ok := false
	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	*/

	// v 또는 ok가 선언되지 않은 경우
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
