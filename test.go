package main

import (
	"fmt"
	"time"
)

func main() {

	go gotest(2)
	go gotest(3)
	go gotest(4)
	go gotest(1)
	time.Sleep(time.Second *2)
}

func gotest(a int){
	if a == 1{
		panic(1)
	}
	fmt.Println(a)
}