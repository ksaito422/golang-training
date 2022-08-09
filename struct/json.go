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

func Json() {
	f, err := os.Open("book.json")
	if err != nil {
		log.Fatal("file open error: ", err)
	}

	d := json.NewDecoder(f)
	var b Book
	d.Decode(&b)
	fmt.Println(b)
}
