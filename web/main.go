package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	// URLごとの処理登録 contoller?
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
		log.Info().Msg("receive hello world request")
	})

	fmt.Println("Start listening at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprint(os.Stderr, "")
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
}
