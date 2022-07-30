package main

import (
	"fmt"
)

func main() {
	// スライスやマップの書く要素に対してループ
	scketches := []string{"Dead Parrot", "Killer joke", "Spanish Inquisition", "Spam"}
	for i, s := range scketches {
		fmt.Println(i, s)
	}

	// 1変数だけ書けばインデックスのみを受け取れる
	for i := range scketches {
		fmt.Println(i)
	}

	// ブランク識別子でインデックスを読み飛ばして値だけを使う
	for _, s := range scketches {
		fmt.Println(s)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("10回繰り返す")
	}
}
