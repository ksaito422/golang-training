package main

import (
	"fmt"
)

func main() {
	Alert()
	normalize, _ := NormalizeString("Hello")
	fmt.Println(normalize)

	Any()
}
