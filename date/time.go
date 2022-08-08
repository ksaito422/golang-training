package main

import (
	"fmt"
	"time"
)

func Time() {
	// 5分を作成
	fiveMinute := 5 * time.Minute

	// intとは型違いで直接演算できないので、即値との計算以外は
	// time.Durationへの明示的なキャストが必要
	var seconds int = 10
	tenSeconds := time.Duration(seconds) * time.Second

	// Timeの演算でDuration作成
	past := time.Date(1955, time.November, 12, 6, 38, 0, 0, time.UTC)
	dur := time.Now().Sub(past)

	fmt.Println("five minutes = ", fiveMinute)
	fmt.Println("ten seconds = ", tenSeconds)
	fmt.Println("past = ", past)
	fmt.Println("dur = ", dur)
}
