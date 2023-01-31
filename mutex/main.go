package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var data map[string]string

func main() {
	data = map[string]string{"hoge": "fuga"}
	mu = sync.RWMutex{}

	go read()
	go read()
	go write()
	go read()

	time.Sleep(5 * time.Second)
}

func read() {
	fmt.Println("read start")
	mu.RLock()
	defer mu.RUnlock()

	time.Sleep(1 * time.Second)

	fmt.Println("read complete", data["hoge"])
}

func write() {
	fmt.Println("write start")
	mu.Lock()
	defer mu.Unlock()

	time.Sleep(2 * time.Second)

	data["hoge"] = "piyo"

	fmt.Println("write complete")
}
