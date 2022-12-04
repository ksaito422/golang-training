package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

type User struct {
	ID   int
	Name string
	Age  *int
}

type A struct {
	User User
}

type B struct {
	User    User
	Country *string
}

func main() {
	paramA := A{
		User: User{
			ID:   1,
			Name: "saito",
			Age:  Ptr(31),
		},
	}
	paramB := B{
		User: User{
			ID:   1,
			Name: "saito",
			Age:  Ptr(31),
		},
		Country: Ptr("Japan"),
	}
	if err := validateParam(paramA); err != nil {
		fmt.Printf("Error paramA:%v\n", err)
	}
	fmt.Printf("ID: %v, Name: %v, Age: %v\n", paramA.User.ID, paramA.User.Name, *paramA.User.Age)

	if err := validateParam(paramB); err != nil {
		fmt.Printf("Error paramB: %v\n", err)
	}
	fmt.Printf("ID: %v, Name: %v, Age: %v, Address: %v\n", paramB.User.ID, paramB.User.Name, *paramB.User.Age, *paramB.Country)
}

func validateParam[T A | B](param T) error {
	if err := validateParamInternal(param); err != nil {
		return err
	}
	return nil
}

func validateParamInternal(param any) error {
	switch v := param.(type) {
	case A:
		if err := validateUserAge(v.User); err != nil {
			return err
		}

	case B:
		if err := validateUserAge(v.User); err != nil {
			return err
		}
		if v.Country == nil {
			return xerrors.Errorf("Countryは必須です: %v", v)
		}
	}

	return nil
}

func validateUserAge(param User) error {
	if param.Age == nil {
		return xerrors.Errorf("User.Ageは必須です")
	}
	return nil
}

func Ptr[T any](x T) *T {
	return &x
}
