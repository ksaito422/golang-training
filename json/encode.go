package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type user struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"user_name"`
	Languages []string `json:"languages"`
}

func Encode() {
	var b bytes.Buffer
	u := user{
		UserID:   "001",
		UserName: "gopher",
	}
	_ = json.NewEncoder(&b).Encode(u)
	fmt.Printf("%v\n", b.String())
}

func Marshal() {
	u := user{
		UserID:    "001",
		UserName:  "gopher",
		Languages: []string{},
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
