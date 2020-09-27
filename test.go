package main

import (
	"fmt"
	"time"
)

var x,y int
func main() {
	go func(){
		x=1
		y=1
	}()
	go func(){
		if y==1 {
			h := y/x
			fmt.Println(h)
		}
	}()
time.Sleep(time.Second)
}