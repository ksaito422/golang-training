package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("関数定義 ", calc(1, 2))
	age, err := calcAge(1, 4, 5)
	fmt.Println("返り値に名前をつける ", age, err)

	m := func(x, y int) int {
		return x * y
	}

	fmt.Println(10, 20, m(10, 20))
}

// 関数定義
func calc(x, y int) int {
	return x * y
}

// 返り値に名前をつける
func calcAge(y int, m time.Month, d int) (age int, err error) {
	b := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	n := time.Now()

	if b.After(n) {
		err = errors.New("誕生日が未来です")
		return
	}

	for {
		b = time.Date(y+age+1, m, d, 0, 0, 0, 0, time.Local)
		if b.After(n) {
			return
		}
		age++
	}
}

func doCalc(x, y int, f func(int, int) int) {
	fmt.Println()
}
