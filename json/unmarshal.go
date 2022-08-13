package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func Unmarshal() {
	s := `{"origin":"255.255.255.255","url":"https://httpbin.org/get"}`
	var resp ip
	if err := json.Unmarshal([]byte(s), &resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
