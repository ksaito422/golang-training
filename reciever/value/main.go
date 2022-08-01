package main

import (
	"fmt"
)

// 値レシーバ: 元の値とは別のコピーした値が関数に渡されるので元の値は変更不可

type Square struct {
	width  float64
	height float64
}

func (s Square) Area() float64 {
	return s.width * s.height
}

func main() {
	square := Square{3.0, 4.0}
	fmt.Println(square.Area())
}
