package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	// ポインタ型
	p := &Point{X: 10, Y: 5}
	// ポインタ型
	o := new(Point)
	o.X = 10
	o.Y = 5
	// ポインタ型じゃない
	q := Point{X: 10, Y: 5}

	fmt.Printf("pX = %d pY = %d\n", p.X, p.Y)
	fmt.Printf("oX = %d oY = %d\n", o.X, o.Y)
	fmt.Printf("qX = %d qY = %d\n", q.X, q.Y)

	fmt.Println("p ", p)
	fmt.Println("&p ", &p)
	fmt.Println("o ", o)
	fmt.Println("&o ", &o)
	fmt.Println("q ", q)
	fmt.Println("&q ", &q)

	// 関数に渡して元の初期化した構造体がどうなるか
	pp(p)
	qq(q)
	fmt.Printf("pX = %d pY = %d\n", p.X, p.Y)
	fmt.Printf("qX = %d qY = %d\n", q.X, q.Y)
}

func pp(p *Point) {
	p.X = 11
	p.Y = 6
}

func qq(q Point) {
	q.X = 11
	q.Y = 6
}
