package main

import "fmt"

func main() {
	defer func() {
		defer func() {
			fmt.Println("3:", recover())
		}()
	}()
	defer func() {
		func() {
			fmt.Println("2:", recover())
		}()
	}()
	func() {
		defer func() {
			fmt.Println("1:", recover())
		}()
	}()
	panic(121)
}

