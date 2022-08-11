package main

import (
	"fmt"
)

func Any() {
	var slices = []any{
		"関ヶ原",
		1600,
	}

	var ieyasu = map[string]any{
		"名前":  "徳川家康",
		"生まれ": 1543,
	}

	fmt.Println(slices)
	fmt.Println(ieyasu)
}
