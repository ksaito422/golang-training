package main

import ()

type Person struct {
	FirstName string
	LastName  string
}

// ファクトリ関数のメリット
// ゼロ値以外の初期値を与えられる
// よく使うユースケースが複数ある場合、ケースごとにファクトリ関数を用意できる
// 入力値をバリデーションできる
// GoDoc上で型の説明のそばに表示されユーザーがインスタンス作成方法で迷わずに済む
func NewPerson(first, last string) *Person {
	return &Person{
		FirstName: first,
		LastName:  last,
	}
}
