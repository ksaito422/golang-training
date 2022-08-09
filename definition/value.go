package main

type SKUCode string

const (
	skuCD = "T01230101"
)

func (c SKUCode) ItemCD() string {
	return skuCD[0:5]
}

func (c SKUCode) SizeCD() string {
	return skuCD[5:7]
}

func (c SKUCode) ColorCD() string {
	return skuCD[7:9]
}
