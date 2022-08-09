package main

import (
	"fmt"
)

func main() {
	var b Book
	fmt.Println("b = ", b)

	b2 := Book{
		Title: "title",
	}
	fmt.Println("b2 = ", b2)

	b3 := &Book{
		Title: "カンフーマック",
	}
	fmt.Println("b3 = ", b3)
	fmt.Println("b3 = ", *b3)

	Json()
	person := NewPerson("田中", "太朗")
	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
}
