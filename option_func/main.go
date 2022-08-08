package main

import (
	"fmt"
)

type Portion int

const (
	Regular Portion = iota // 普通
	Small                  // 小盛り
	Large                  // 大盛り
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

func main() {
	var tempuraUdon = NewUdon(Large, false, 2)
	fmt.Println("天ぷらうどん = ", tempuraUdon)

	udonOption := Option{Large, false, 0}
	var kitsuneUdon = NewUdon2(udonOption)
	fmt.Println("きつねうどん = ", kitsuneUdon)

	oomorikitsune := NewUdon3(Large).Abburaage().Ebiten(3).Order()
	fmt.Println("大盛りキツネ = ", oomorikitsune)

	tokuseiUdon := NewUdon4(OptAburaage(), OptEbiten(3))
	fmt.Println("特製うどん = ", tokuseiUdon)
}
