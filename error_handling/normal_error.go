package main

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

func FindBook(isbn string) error {
	return ErrNotFound
}
