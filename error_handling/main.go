package main

import (
	"fmt"
)

func main() {
	err := FindBook("not found")
	fmt.Println(err)

	_, err = ReadContents("http://localhost")
	fmt.Println(err)
}
