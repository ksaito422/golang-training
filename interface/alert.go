package main

import (
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
)

// 全ての構造体のインターフェース
type Warning interface {
	Show(message string)
}

// コンソールにメッセージを通知する構造体
type ConsoleWarning struct{}

// デスクトップにメッセージを通知する構造体
type DesktopWarning struct{}

func (c ConsoleWarning) Show(message string) {
	fmt.Fprintf(os.Stderr, "[%s]: %s\n", os.Args[0], message)
}

func (d DesktopWarning) Show(message string) {
	beeep.Alert(os.Args[0], message, "")
}

func Alert() {
	// Show()メソッドを持つインスタンスはなんでも代入可能
	var warn Warning

	warn = &ConsoleWarning{}
	warn.Show("Hello World to console")

	warn = &DesktopWarning{}
	warn.Show("Hello World to desktop")
}
