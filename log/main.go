package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, _ := os.Create("log.txt")
	log.SetOutput(io.MultiWriter(file, os.Stderr))

	log.Println("ファイルと標準エラー出力に同時に出力します")

	n := 10
	s := "文字列"
	c := 1 + 2i
	log.Printf("%d, %sなどを使って変数出力もできます。%vはどんな型でも表示します", n, s, c)
}
