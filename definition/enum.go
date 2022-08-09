package main

type Season int

const (
	Peak Season = iota + 1
	Normal
	Off
)

func (s Season) Price(price float64) Season {
	if s == Peak {
		return s + 200
	}
	return s
}
