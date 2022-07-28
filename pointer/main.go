package main

import (
	"fmt"
)

func main() {
	// 変数のポインター型には*を前置する
	// 既存の変数のポインターを取り出すには&を利用する
	i := 10
	var p *int

	fmt.Println("p = ", p)
	fmt.Println("p = ", &p)

	p = &i

	fmt.Println("i = ", i)
	fmt.Println("i = ", &i)
	fmt.Println("p = ", *p)
	fmt.Println("p = ", p)
}
