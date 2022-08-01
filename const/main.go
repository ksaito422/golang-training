package main

import (
	"fmt"
)

// コンパイル時に計算される
// バイナリファイルに答えが入った状態でないといけない
// 関数の返り値、構造体のインスタンス、マップ、スライスなどの複合型は扱えない

const (
	a = 1
	b = 1 + 2
	// c = 9223372036854775807 + 1
	c = 9223372036854775807
	d = "hello world"
	e = iota + 10
)

func main() {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
}
