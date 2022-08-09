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

	itemCD, sizeCD, colorCD := SKUCode.ItemCD("T01230101"), SKUCode.SizeCD("T01230101"), SKUCode.ColorCD("T01230101")
	fmt.Println(itemCD, sizeCD, colorCD)

	fmt.Println(Season.Price(Peak, 200))

	Secret()
}
