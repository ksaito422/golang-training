package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i int = 123
	// 数値同士の変換は、かっこでくくり型を前置する
	var f float64 = float64(i)
	// 64ビットのintと、int64も明示的な変換が必要
	var i64 int64 = int64(i)
	// boolへの変換は比較演算子を使う
	var b bool = i != 0

	fmt.Println("f = ", f)
	fmt.Println("i64 = ", i64)
	fmt.Println("b = ", b)

	// 文字列との変換はstrconvパッケージを利用
	in := 12345
	// strconvの数値入力はint64, uint64, float64なので、それ以外の変数を使うときは型変換が必要
	s := strconv.FormatInt(int64(in), 10)
	fmt.Println("s = ", s)
	fmt.Println("s = ", reflect.TypeOf(s))

	// parse系は変換失敗時にエラーを返す
	// 成功時のerrはnil
	f, err := strconv.ParseFloat("12.3456", 64)
	fmt.Println("f = ", f)
	fmt.Println("f =", reflect.TypeOf(f))
	fmt.Println("err = ", err)
	fmt.Println("err = ", reflect.TypeOf(err))
}
