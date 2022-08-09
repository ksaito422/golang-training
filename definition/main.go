package main

import (
	"fmt"
)

func main() {
	Reflect()

	fmt.Println(HTTPStatus.String(200))
	fmt.Println(HTTPStatus.String(401))
	fmt.Println(HTTPStatus.String(402))
	fmt.Println(HTTPStatus.String(403))
	fmt.Println(NationalRoute.String(200))
	fmt.Println(NationalRoute.String(401))
	fmt.Println(NationalRoute.String(402))
	fmt.Println(NationalRoute.String(403))
}
