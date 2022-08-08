package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func Reflect() {
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))

	fmt.Println(reflect.TypeOf(1))
}
