package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在時刻のtime.Timeインスタンス取得
	now := time.Now()

	// 指定日時のtime.Timeインスタンス取得
	tz, _ := time.LoadLocation("America/Los_angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))

	Time()
}
