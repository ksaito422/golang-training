package main

import (
	"fmt"
)

type CarType int
type CarOption uint64

const (
	Sedan CarType = iota + 1
	Hatchback
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriverAssist
)

func main() {
	var t CarType
	t = SUV

	fmt.Println("car type = ", t)

	var o CarOption
	o = HeatedSeat
	fmt.Println("car option = ", o)
}
