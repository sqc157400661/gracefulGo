package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct{
	sex int
	age int
}
func main() {
	t := Person{}
	fmt.Println(unsafe.Alignof(t))
	fmt.Println(unsafe.Alignof(t.age))
	fmt.Println(unsafe.Alignof(t.sex))

	fmt.Println(reflect.TypeOf(t).Align())
	fmt.Println(reflect.TypeOf(t).FieldAlign())
}