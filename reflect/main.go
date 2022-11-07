package main

import (
	"fmt"
	"reflect"
)

type User struct {
	NickName string
	Age      int32
}

func main() {
	stringV := "test"
	intV := 12345

	fmt.Println("=== string ===")
	typePrint(stringV)
	valuePrint(stringV)

	fmt.Println("=== int ===")
	typePrint(intV)
	valuePrint(intV)

	fmt.Println("=== *string ===")
	typePrint(&stringV)
	valuePrint(&stringV)

	createStruct()
	createSlice()
	createMap()
}

// 型の情報を表す
func typePrint(v any) {
	// Kind()は型の種類を表す
	tv := reflect.TypeOf(v)
	fmt.Println("Kind:", tv.Kind(), "Name:", tv.Name())
}

// 値の情報を表す
func valuePrint(v any) {
	rv := reflect.ValueOf(v)
	fmt.Println("Kind:", rv.Kind(), "Type:", rv.Type(), "Interface:", rv.Interface())
}

// 構造体の生成
func createStruct() {
	tStruct := reflect.TypeOf(User{})
	vStruct := reflect.New(tStruct)
	user := vStruct.Interface().(*User)

	user.NickName = "user1"
	user.Age = 10
	fmt.Println("=== createStruct ===")
	fmt.Println(user)
}

// スライスの生成
func createSlice() {
	tSlice := reflect.TypeOf([]User{})
	vSlice := reflect.MakeSlice(tSlice, 0, 2)
	users := vSlice.Interface().([]User)

	users = append(users, User{
		NickName: "user2",
		Age:      20,
	})
	fmt.Println("=== createSlice ===")
	fmt.Println(users)
}

// マップの生成
func createMap() {
	tMap := reflect.TypeOf(map[string]User{})
	vMap := reflect.MakeMap(tMap)
	userMap := vMap.Interface().(map[string]User)

	userMap["hoge"] = User{
		NickName: "user3",
		Age:      30,
	}
	fmt.Println("=== createMap ===")
	fmt.Println(userMap)
}
