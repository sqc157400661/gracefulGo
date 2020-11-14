package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func main() {
	i := 1
	fmt.Println("i type name", reflect.TypeOf(i).Name())
	fmt.Println("i type kind", reflect.TypeOf(i).Kind())

	u := User{Name: "poloxue"}
	fmt.Println("u type name", reflect.TypeOf(u).Name())
	fmt.Println("u type kind", reflect.TypeOf(u).Kind())

	if reflect.TypeOf(u).Kind() == reflect.Struct {
		fmt.Println("u kind is struct")
	}
 
}
