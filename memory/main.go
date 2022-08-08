package main

import (
	"fmt"
)

// メモリ再確保が時には、重い処理となることがある
// メモリの最大値、最小値がわかっている場合は指定するとパフォーマンス向上につながる
func main() {
	// 正確な長さがわかっている場合
	s1 := make([]int, 1000)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	// 正確な長さがわからないが最大量の見込みがつく場合
	s2 := make([]int, 0, 1000)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}
