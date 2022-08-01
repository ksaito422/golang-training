package main

import (
	"fmt"
)

// ポインタレシーバ: ポインタとして引数を渡すため関数内でオブジェクトの値を変更可能
// 構造体型のデータの値を変えたい時にポインタレシーバを使う
// レシーバに渡す構造体がメモリ上で大きい領域を占めている場合にポインタレシーバを使う(メモリに優しい)

type Square struct {
	width  float64
	height float64
}

func (s *Square) Reshape(w float64, h float64) {
	s.width = w
	s.height = h
}

func main() {
	square := Square{3.0, 4.0}
	fmt.Println(square)

	square.Reshape(5.0, 6.0)
	fmt.Println(square)
}
