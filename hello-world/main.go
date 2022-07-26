package main

import (
	"fmt"
)

var num5 int = 123

// num6 := 123 ここでは宣言できない

func main() {
	var num1 int = 123
	var num2 = 123
	num3 := 123
	// num4 int := 123 型は不要

	fmt.Println("Hello World")
	fmt.Println(num1, num2, num3, num5)
}
