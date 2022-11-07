package main

import "fmt"

type StructA struct {
	fieldA string
}

func (a *StructA) Hoge() string {
	return "hoge"
}

type StructB struct {
	// 構造体Aを構造体Bに埋め込む
	StructA
	fieldB string
}

func main() {
	a := &StructA{fieldA: "aaa"}
	b := &StructB{StructA{fieldA: "aaa"}, "bbb"}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b.Hoge())
}
