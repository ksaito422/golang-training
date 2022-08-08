package main

type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon3(p Portion) *fluentOpt {
	// デフォルトはコンストラクタ関数で設定
	// 必須オプションはここに付与可能
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

func (o *fluentOpt) Abburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}
