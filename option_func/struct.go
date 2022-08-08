package main

import (
	"time"
)

type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon2(opt Option) *Udon {
	// ゼロ値に対するデフォルト値処理は関数内部で行う
	// 朝食時間は海老天1本無料
	if opt.ebiten == 0 && time.Now().Hour() < 10 {
		opt.ebiten = 1
	}
	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}
