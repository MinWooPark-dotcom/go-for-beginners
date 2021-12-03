package main

import (
	"fmt"
	"strings"
)

func main() {
	fileName := "abc"
	fmt.Println(fileName)
	removeColon := strings.ReplaceAll(fileName, "", "")
	fmt.Println(removeColon)
}
