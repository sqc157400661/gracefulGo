package main

import "fmt"

func alwaysFalse() bool { return false }
func main() {
	switch alwaysFalse() {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	switch alwaysFalse()
	{
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}
}
