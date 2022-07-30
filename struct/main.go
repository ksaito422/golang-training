package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Book struct {
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	ReleasedAt time.Time `json:"release_at"`
	ISBN       string    `json:"isbn"`
}

// func main() {
// 	var b Book
// 	fmt.Println("b = ", b)
//
// 	b2 := Book{
// 		Title: "title",
// 	}
// 	fmt.Println("b2 = ", b2)
//
// 	b3 := &Book{
// 		Title: "カンフーマック",
// 	}
// 	fmt.Println("b3 = ", b3)
// 	fmt.Println("b3 = ", *b3)
// }

func main() {
	f, err := os.Open("book.json")
	if err != nil {
		log.Fatal("file open error: ", err)
	}

	d := json.NewDecoder(f)
	var b Book
	d.Decode(&b)
	fmt.Println(b)
}
